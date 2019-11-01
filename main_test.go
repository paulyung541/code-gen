package main

import (
	"math"
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestJson(t *testing.T) {
	var m map[string]interface{}
	data, _ := ioutil.ReadFile("test.json")
	json.Unmarshal(data, &m)

	//number is float64
	a := m["updateUser"].(float64)
	b := m["good"].(bool)

	//object is []byte
	data := m["data"].([]byte)

	t.Log(a, b)
	t.Logf("data type = %T data value = %v", data, data)

}

func TestMath(t *testing.T) {
	var a float64 = 1000.00001
	if a - math.Ceil(a) == 0 {
		t.Error("is 0")
	}
}