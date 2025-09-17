package main

import (
	"encoding/json"
)

type message struct {
	Msg string `json:"message"`
}

func MarshalJSON(b []byte) ([]byte, error) {
	var m message
	err := json.Unmarshal(b, &m)
	if err == nil {
		return []byte(m.Msg), nil
	}
	return nil, err
}
