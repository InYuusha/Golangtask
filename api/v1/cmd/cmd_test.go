package cmd

import (
	"reflect"
	"testing"

	c "github.com/InYuusha/api/v1/cmd/controller"
	"github.com/InYuusha/api/v1/cmd/store"
	"github.com/go-playground/assert/v2"
)

var kvs *c.KeyValueStore = store.NewKeyValueStore()

func TestSetCmd(t *testing.T) {
	setTestCases := []struct {
		key       string
		value     string
		expiry    int64
		condition string
	}{
		{
			key:       "a",
			value:     "2",
			expiry:    60,
			condition: "",
		},
		{
			key:       "c",
			value:     "2",
			expiry:    0,
			condition: "NX",
		},
		{
			key:       "l1",
			value:     "6",
			expiry:    60,
			condition: "XX",
		},
	}

	for _, tc := range setTestCases {
		result, err := kvs.SetCmd(tc.key, &tc.value, &tc.expiry, tc.condition)
		assert.Equal(t, err, nil)
		assert.Equal(t, result, "OK")
	}

}

func TestGetCmd(t *testing.T) {
	getTestCases := []struct {
		key string
	}{
		{
			key: "a",
		},
		{
			key: "c",
		},
		{
			key: "l1",
		},
	}

	for _, tc := range getTestCases {
		result, err := kvs.GetCmd(tc.key)
		assert.Equal(t, err, nil)
		if reflect.TypeOf(result).Kind() != reflect.String {
			t.Errorf("Expected result to be of type string, but got %v", reflect.TypeOf(result))
		} 
		t.Logf("Test Cases Passed for input %v", tc)

	}
}

func TestQpushCmd(t *testing.T) {
	qpushTestCases := []struct {
		key    string
		values []string
	}{
		{
			key:    "a",
			values: []string{"2", "5", "6"},
		},
		{
			key:    "c",
			values: []string{"9", "4", "6"},
		},
		{
			key:    "l1",
			values: []string{"0", "3", "1"},
		},
	}

	for _, tc := range qpushTestCases {
		result, err := kvs.Qpush(tc.key, tc.values)
		assert.Equal(t, err, nil)
		assert.Equal(t, result, "OK")
		t.Logf("Test Cases Passed for input %v",tc)
	}
}

func TestQpop(t *testing.T) {
	qpushTestCases := []struct {
		key string
	}{
		{
			key: "a",
		},
		{
			key: "c",
		},
		{
			key: "l1",
		},
	}

	for _, tc := range qpushTestCases {
		result, err := kvs.Qpop(tc.key)
		assert.Equal(t, err, nil)
		if reflect.TypeOf(result).Kind() != reflect.String {
			t.Errorf("Expected result to be of type string, but got %v", reflect.TypeOf(result))
		}
		t.Logf("Test Cases Passed for input %v",tc)
	}
}
