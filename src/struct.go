package src

import "time"

type SurrDB struct {
	// Just address
	Host      string
	Namespace string
	Database  string
	User      string
	Pass      string

	Schema  string
	Pretty  bool
	Timeout time.Duration
	Comple  int

	Query string
}

type Payload struct {
	Time   string      `json:"time"`
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}
