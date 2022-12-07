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
}
