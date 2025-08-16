package model

import "time"

type LogObject struct {
	StructTimeStamp time.Time `json:"timestamp"`
	StructError     string    `json:"error"`
}
