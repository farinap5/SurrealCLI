package src

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

func (s *SurrDB) InitCLI() {
	Banner()
	s.TestConnection()
	print("\n")
	p := prompt.New(
		s.execute,
		completer,
		prompt.OptionPrefix("["+s.Database+"]> "),
		prompt.OptionCompletionOnDown(),
		prompt.OptionMaxSuggestion(uint16(s.Comple)),
	)
	p.Run()
}

func (s SurrDB) ContactSurr(p string) {
	resp, _ := s.Requester(p)
	if s.Pretty {
		s.PrettyPrint(resp)
	} else {
		s.Print(resp)
	}

}

func (s *SurrDB) execute(p string) {
	ps := strings.Split(p, " ")
	if p == ".help" {
		Help()
	} else if p == ".options" {
		s.Options()
	} else if ps[0] == ".set" && len(ps) == 3 {
		s.SetVars(ps[1], ps[2])
	} else if ps[0] == ".save" {
		s.savecommands(ps, p)
	} else if ps[0] == ".delete" {
		delete(ps)
	} else if ps[0] == ".show" {
		showstorage(ps)
	} else if ps[0] == ".run" {
		s.run(ps, p)
	} else {
		s.Query = p
		s.ContactSurr(p)
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "SELECT", Description: "Select data"},
		{Text: "FROM", Description: "From where to select"},
		{Text: "WHERE", Description: "Define property"},
		{Text: "LIMIT", Description: "Limit amount of rows"},
		{Text: "CREATE", Description: "Create new record"},
		{Text: "RELATE", Description: "Create relations"},
		{Text: "CONTENT", Description: "Set content"},
		{Text: "INFO", Description: "Show information"},
		{Text: "FOR", Description: "Show info for something"},
		{Text: "DB", Description: "Database"},
		{Text: "NS", Description: "Namespace"},
		{Text: "TABLE", Description: "Table"},
		{Text: "GROUP", Description: "Group data"},
		{Text: "BY", Description: "By"},

		{Text: ".help", Description: "Show help menu"},
		{Text: ".options", Description: "Env variables"},
		{Text: ".set", Description: "Set variable"},
		{Text: ".save", Description: "Save profile|query"},
		{Text: ".show", Description: "Show profiles|queries"},
		{Text: ".delete", Description: "Delete a profile|query"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func (s SurrDB) savecommands(ps []string, p string) {
	switch ps[1] {
	case "profile":
		s.DBSaveProfile(ps[2])
	case "query":
		s.DBSaveQuery(ps[2])
	default:
		PrintErr("Not a command.")
	}
}

func (s *SurrDB) run(ps []string, p string) {
	switch ps[1] {
	case "profile":
		s.DBSetProfileByIdx(ps[2])
	case "query":
		q, e := DBGetQueryByIdx(ps[2])
		if e {
			PrintSuc("Running ")
			s.ContactSurr(q)
		} else {
			PrintErr("Error searching query name.")
		}
	default:
		PrintErr("Not a command.")
	}
}

func showstorage(ps []string) {
	switch ps[1] {
	case "profiles":
		DBShowProfiles()
	case "queries":
		DBShowQueries()
	default:
		PrintErr("Not a command.")
	}
}

func delete(ps []string) {
	switch ps[1] {
	case "profile":
		DBDropIdx(ps[2])
	case "query":
		DBDropQueryIdx(ps[2])
	default:
		PrintErr("Not a command.")
	}
}
