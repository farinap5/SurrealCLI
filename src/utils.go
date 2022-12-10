package src

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/term"
	"io"
	"os"
	"syscall"
)

func PrintErr(s string) {
	println("[\u001B[1;31m!\u001B[0;0m]- " + s)
}

func PrintSuc(s string) {
	println("[\u001B[1;32mOK\u001B[0;0m]- " + s)
}

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

/*func GetPayload(j string) {
	println(j)
	data := []Payload{}
	err := json.Unmarshal([]byte(j), &data)
	if err != nil {
		HandErrs(err)
	}
	x, err := json.Marshal(data[0].Result)
	if err != nil {
		HandErrs(err)
	}
	fmt.Println(string(x))
	tableprinter.PrintJSON(os.Stdout, x)

}

func PrettyTable() {
	printer := tableprinter.New(os.Stdout)
	printer.PrintJSON()
}*/

func (s SurrDB) Print(p string) {
	fmt.Printf("%s\n", p)
}

func (s *SurrDB) SetVars(v string, n string) {
	if v == "user" || v == "User" {
		s.User = n
		println("[\u001B[1;32mOK\u001B[0;0m]- Use <- " + s.User)
		s.Pass = GetNoEchos("[password:" + s.User + "]: ")
		print("\n")
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
		println("[\u001B[1;31m!\u001B[0;0m]- There was a problem with authentication.\nUse \u001B[33m.set user <username>\u001B[0m to reset credentials.")
	} else {
		println("[\u001B[1;31m!\u001B[0;0m]- Error!")
	}
}

func GetNoEchos(s string) string {
	fmt.Print(s)
	bytepw, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		os.Exit(1)
	}
	return string(bytepw)
}

func FromSTDIN() (string, bool) {
	//var sin bool
	f, err := os.Stdin.Stat()
	if err != nil {
		HandErrs(err)
	}

	if f.Mode()&os.ModeNamedPipe == 0 {
		return "", false
	} else {
		read := bufio.NewReader(os.Stdin)
		var output []rune

		for {
			input, _, err := read.ReadRune()
			if err != nil && err == io.EOF {
				break
			}
			output = append(output, input)
		}
		return string(output), true
	}
}
