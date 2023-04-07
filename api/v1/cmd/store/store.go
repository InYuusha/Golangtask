package store

import (
	c "github.com/InYuusha/api/v1/cmd/controller"
	"github.com/InYuusha/api/v1/cmd/dto"
)

func NewKeyValueStore() *c.KeyValueStore {
	return &c.KeyValueStore{
		Data: make(map[string]*dto.Value),
	}
}
