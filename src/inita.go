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
	var qry = flag.String("q", "none", "Run query.")
	var prof = flag.String("profile", "none", "Use existent profile.")

	var cop = flag.Int("comp", 5, "Completion/suggestions. Set 0 for disable.")
	var tout = flag.Int("t", 5, "Timeout.")
	var pret = flag.Bool("pretty", true, "Pretty print.")

	flag.Usage = HelpCMD
	flag.Parse()

	var p string
	if *pass == "hide" {
		p = GetNoEchos("[password]: ")
		print("\n")
	} else {
		p = *pass
	}
	s := SurrDB{
		User:      *user,
		Pass:      p,
		Schema:    *sch,
		Namespace: *ns,
		Database:  *db,
		Host:      *host,
		Comple:    *cop,
		Timeout:   time.Duration(*tout),
		Pretty:    *pret,
	}
	DBFileInit()
	if *prof != "none" {
		// if passing profile, use it
		s.DBSetProfileByIdx(*prof)
	}

	// Get data from stdin
	sin, e := FromSTDIN()
	// if stdin execute it
	if e {
		s.ContactSurr(sin)
		return
	}

	if *qry == "none" {
		// if no query from command line enter interactive mode
		s.InitCLI()
	} else {
		s.ContactSurr(*qry)
	}
}
