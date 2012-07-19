package hob

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrJSONDecode = errors.New("error decoding JSON")
)

func ParseJson(jsondata []byte) (st interface{}, err error) {
	err = json.Unmarshal(jsondata, &st)

	if data_type, ok := st.(map[string]interface{})["type"]; ok {
		if data_type == "lww-e-set" {
			return st, nil
		} else {
			return nil, ErrJSONDecode
		}
	}

	return st, nil
}

func Timestamp() (now string) {
	now = time.Now().UTC().Format(time.RFC3339Nano)
	return now
}
