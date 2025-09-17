package main

import (
	"encoding/json"
)

type message struct {
	Msg string `json:"message"`
}

func MarshalJSON(b []byte) (string, error) {
	var m message
	err := json.Unmarshal(b, &m)
	if err == nil {
		return m.Msg, nil
	}
	return "", err
}
