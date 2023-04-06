package controller

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/InYuusha/api/v1/cmd/dto"
)

type KeyValueStore struct {
	sync.RWMutex
	Data map[string]*dto.Value
}

func (kvs *KeyValueStore) SetCmd(key string, val *string, expiry *int64, condition string) (string, error) {
	log.Printf("Key %s value %v expiry %v ", key, *val, *expiry)
	kvs.Lock()
	defer kvs.Unlock()
	v, ok := kvs.Data[key]
	if ok {
		if condition == "NX" {
			return "", fmt.Errorf("key already exists")
		}
		if condition == "XX" {
			v.Val = *val
			v.Expiry = *expiry
			return "OK", nil
		}
	}
	kvs.Data[key] = &dto.Value{Val: *val, Expiry: *expiry, Created: time.Now().UnixNano()}
	return "OK", nil
}

func (kvs *KeyValueStore) GetCmd(key string) (string, error) {
	kvs.RLock()
	defer kvs.RUnlock()
	v, ok := kvs.Data[key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}
	if v.Expiry > 0 && time.Now().UnixNano() > v.Created+v.Expiry*1e9 {
		delete(kvs.Data, key)
		return "", fmt.Errorf("key not found")
	}
	return v.Val, nil
}

func (kvs *KeyValueStore) Qpush(key string, values []string) (string, error) {
	kvs.Lock()
	defer kvs.Unlock()
	v, ok := kvs.Data[key]
	if !ok {
		kvs.Data[key] = &dto.Value{}
		v = kvs.Data[key]
	}
	for _, value := range values {
		v.Val += " " + value
	}
	return "OK", nil
}

func (kvs *KeyValueStore) Qpop(key string) (string, error) {
	kvs.Lock()
	defer kvs.Unlock()
	v, ok := kvs.Data[key]
	if !ok {
		return "", fmt.Errorf("queue is empty")
	}
	values := strings.Split(v.Val, " ")
	last := values[len(values)-1]
	if len(values) == 1 {
		delete(kvs.Data, key)
	} else {
		v.Val = strings.Join(values," ")
	}
	return last, nil
}


func (kvs *KeyValueStore) Bqpop(key string, timeout float64) (string, error) {
	deadline := time.Now().Add(time.Duration(timeout * 1e9))
	for {
		kvs.Lock()
		v, ok := kvs.Data[key]
		if !ok {
			kvs.Unlock()
			if time.Now().After(deadline) {
				return "", fmt.Errorf("queue is empty")
			}
			time.Sleep(100 * time.Millisecond)
			continue
		}
		values := strings.Split(v.Val, " ")
		if len(values) > 0 {
			last := values[len(values)-1]
			if len(values) == 1 {
				delete(kvs.Data, key)
			} else {
				v.Val = strings.Join(values," ")
			}
			kvs.Unlock()
			return last, nil
		}
		kvs.Unlock()
		if time.Now().After(deadline) {
			return "", fmt.Errorf("queue is empty")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

