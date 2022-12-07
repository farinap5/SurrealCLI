package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cheynewallace/tabby"
	"os"
)

func HandErrs(e error) {
	fmt.Println(e.Error())
	os.Exit(1)
}

func (s SurrDB) PrettyPrint(p string) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(p), "", "  ")
	if err != nil {
		HandErrs(err)
	}
	fmt.Printf("%s\n", out.String())
}

func (s SurrDB) Print(p string) {
	fmt.Printf("%s\n", p)
}

func Help() {
	t := tabby.New()
	t.AddHeader("COMMAND", "DESCRIPTION")
	t.AddLine(".help", "Show help menu")
	t.AddLine(".options", "Env variables")
	t.AddLine(".set", "Set variable")
	print("\n")
	t.Print()
	print("\n")
}

func (s SurrDB) Options() {
	t := tabby.New()
	t.AddHeader("VARIABLE", "VALUE")
	t.AddLine("Host", s.Host)
	t.AddLine("User", s.User)
	t.AddLine("Namespace", s.Namespace)
	t.AddLine("Database", s.Database)

	t.AddLine("Schema", s.Schema)
	t.AddLine("Pretty", s.Pretty)
	t.AddLine("Timeout", s.Timeout)
	t.AddLine("Suggestion", s.Comple)
	print("\n")
	t.Print()
	print("\n")
}

func (s *SurrDB) SetVars(v string, n string) {
	if v == "user" || v == "User" {
		s.User = n
		println("[\u001B[1;32mOK\u001B[0;0m]- Use <- " + s.User)
	} else if v == "pass" || v == "Pass" {
		s.Pass = n
		println("[\u001B[1;32mOK\u001B[0;0m]- Pass <- ********")
	} else if v == "host" || v == "Host" {
		s.Host = n
		println("[\u001B[1;32mOK\u001B[0;0m]- Host <- " + s.Host)
	} else if v == "pretty" || v == "Pretty" {
		if s.Pretty {
			s.Pretty = false
			println("[\u001B[1;32mOK\u001B[0;0m]- Pretty print <-", s.Pretty)
		} else {
			s.Pretty = true
			println("[\u001B[1;32mOK\u001B[0;0m]- Pretty print <- ", s.Pretty)
		}
	} else if v == "ns" || v == "NS" || v == "nameserver" {
		s.Namespace = n
		println("[\u001B[1;32mOK\u001B[0;0m]- Namespace <- " + s.Namespace)
	} else if v == "db" || v == "DB" || v == "database" {
		s.Database = n
		println("[\u001B[1;32mOK\u001B[0;0m]- Database <- " + s.Database)
	} else if v == "schema" || v == "Schema" || v == "sch" {
		if n == "http" || n == "https" {
			s.Schema = n
			println("[\u001B[1;32mOK\u001B[0;0m]- Schema <- " + s.Schema)
		} else {
			println("[\u001B[1;31m!\u001B[0;0m]- No options.")
		}
	} else {
		println("[\u001B[1;31m!\u001B[0;0m]- No options.")
	}
}

func (s SurrDB) TestConnection() {
	_, c := s.Requester("INFO FOR DB;")
	if c == 200 {
		println("[\u001B[1;32mOK\u001B[0;0m]- Connection is OK!")
	} else if c == 403 {
		println("[\u001B[1;31m!\u001B[0;0m]- There was a problem with authentication.\nUse \u001B[33m.set user <username>\u001B[0m and \u001B[33m.set pass <password>\u001B[0m to reset credentials.")
	} else {
		println("[\u001B[1;31m!\u001B[0;0m]- Error!")
	}
}
