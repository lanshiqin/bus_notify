package db

import (
	"testing"
)

func TestSetString(t *testing.T) {
	key := 1
	Value := "蓝士钦"
	err := SetString(key, Value)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("set success")
	}
}

func TestGetString(t *testing.T) {
	key := 1
	value := GetString(key)
	t.Log(value)
}
