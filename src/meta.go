package src

import "github.com/cheynewallace/tabby"

var Version string = "0.5-NotStable"

func Banner() {
	println("######  \033[33mSurrealCLI\033[0m  ######")
	println("Type `.help` for help meu.")
	println("v " + Version)
}

func Help() {
	t := tabby.New()
	t.AddHeader("COMMAND", "DESCRIPTION")
	t.AddLine(".help", "Show help menu")
	t.AddLine(".options", "Env variables")
	t.AddLine(".set", "Set variable")
	t.AddLine(".save", "Save profile|query")
	t.AddLine(".show", "Show profiles|queries")
	t.AddLine(".delete", "Delete profile|query")
	t.AddLine(".run", "Run profile|query")
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

func HelpCMD() {
	t := tabby.New()
	t.AddHeader("COMMAND", "DESCRIPTION", "DEFAULT")
	t.AddLine("-u", "Username", "root")
	t.AddLine("-p", "Password", "hide password")
	t.AddLine("-host", "Database address \"IP:PORT\"", "0.0.0.0:80")
	t.AddLine("-NS", "Namespace", "surr")
	t.AddLine("-DB", "Database", "surr")
	t.AddLine("-sc", "Schema", "http")
	t.AddLine("-profile", "Connect to a profile", "none")
	t.AddLine("-t", "Connection timeout", "5")
	t.AddLine("-pretty", "Pretty output", "true")
	t.AddLine("-comp", "Number of suggestions", "5")
	print("\n")
	t.Print()
	print("\n")
}
