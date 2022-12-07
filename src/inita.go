package src

import (
	"flag"
	"time"
)

func InitA() {
	var host = flag.String("host", "0.0.0.0:80", "Host. Without schema.")
	var user = flag.String("u", "root", "User.")
	var pass = flag.String("p", "hide", "Password. Hided by default.")
	var db = flag.String("db", "surr", "Database.")
	var ns = flag.String("ns", "surr", "Namespace.")
	var sch = flag.String("sc", "http", "Schema http|https")

	var cop = flag.Int("comp", 5, "Completion/suggestions. Set 0 for disable.")
	var tout = flag.Int("t", 5, "Timeout.")
	var pret = flag.Bool("pretty", true, "Pretty print.")

	flag.Parse()
	p := *pass
	s := SurrDB{
		User:      *user,
		Pass:      p,
		Schema:    *sch,
		Namespace: *ns,
		Database:  *db,
		Host:      *host,

		Comple:  *cop,
		Timeout: time.Duration(*tout),
		Pretty:  *pret,
	}
	s.InitCLI()
}
