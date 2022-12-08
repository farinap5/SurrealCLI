package src

import (
	"database/sql"
	"github.com/cheynewallace/tabby"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var DBconn *sql.DB

func DBFileInit() {
	filename, err := os.UserHomeDir()
	if err != nil {
		PrintErr("Error locating home dir.")
		HandErrs(err)
	}
	filename = filename + "/.local/surrcli.db"

	_, err = os.Open(filename)
	if err != nil {
		PrintErr("Error locating local database. Creating one!")
		_, err = os.Create(filename)
		if err != nil {
			PrintErr("Error creating database: " + err.Error())
		} else {
			PrintSuc("Database created: " + filename)
		}
	}
	tmp, err := sql.Open("sqlite3", filename)
	if err != nil {
		HandErrs(err)
	}
	DBconn = tmp
	DBTableSet()
}

func DBResetAll() {
	dbdropall()
	DBTableSet()
}

func DBTableSet() {
	sttm, err := DBconn.Prepare(`
	CREATE TABLE IF NOT EXISTS Profile (
	    pid 	INTEGER PRIMARY KEY AUTOINCREMENT,
	    Idx 	TEXT NOT NULL,
	    Host	TEXT NOT NULL,
	    Sch 	TEXT NOT NULL,
	    DBUser	TEXT NOT NULL,
	    NS 		TEXT NOT NULL,
	    DB		TEXT NOT NULL,
	    Date 	TEXT NOT NULL
	);
	`)
	if err != nil {
		HandErrs(err)
	} else {
		sttm.Exec()
	}

	sttm, err = DBconn.Prepare(`
	CREATE TABLE IF NOT EXISTS Sess (
	    sid INTEGER PRIMARY KEY AUTOINCREMENT
	);
	`)
	if err != nil {
		HandErrs(err)
	} else {
		sttm.Exec()
	}

	// Save query
	sttm, err = DBconn.Prepare(`
	CREATE TABLE IF NOT EXISTS SQuery (
	    qid 	INTEGER PRIMARY KEY AUTOINCREMENT,
	    Idx 	TEXT,
	    Query 	TEXT
	);
	`)
	if err != nil {
		HandErrs(err)
	} else {
		sttm.Exec()
	}

	// Save output
	sttm, err = DBconn.Prepare(`
	CREATE TABLE IF NOT EXISTS SOut (
	    qid 	INTEGER PRIMARY KEY AUTOINCREMENT,
	    Idx 	TEXT,
	    Query 	TEXT
	)
	`)
	if err != nil {
		HandErrs(err)
	} else {
		sttm.Exec()
	}
}

func dbdropall() {
	sttm, err := DBconn.Prepare(`
	DROP TABLE Profile;
	DROP TABLE Sess;
	DROP TABLE SOut;
	DROP TABLE SQuery;
	`)
	if err != nil {
		HandErrs(err)
	} else {
		sttm.Exec()
	}
}

func (s SurrDB) DBSaveProfile(name string) {
	x := DBValidIndex(name)
	if x {
		PrintErr("Profile exists.")
		return
	}
	sttm, err := DBconn.Prepare(`
		INSERT INTO Profile (Idx,Host,Sch,DBUser,NS,DB,Date) VALUES (?,?,?,?,?,?,datetime('now','localtime'));
	`)
	if err != nil {
		HandErrs(err)
	}
	_, err = sttm.Exec(name, s.Host, s.Schema, s.User, s.Namespace, s.Database)
	if err != nil {
		HandErrs(err)
	} else {
		PrintSuc("Profile saved.")
	}
}

func DBShowProfiles() {
	rw, err := DBconn.Query(`SELECT pid,Idx,Host,Sch,DBUser,NS,DB,Date FROM Profile;`)
	if err != nil {
		HandErrs(err)
	} else {
		t := tabby.New()
		t.AddHeader("ID", "NAME", "HOST", "PROTOCOL", "USER", "NAMESPACE", "DATABASE", "DATE")
		for rw.Next() {
			var pid int
			var Idx, Host, Sch, DBUser, NS, DB, Date string
			rw.Scan(&pid, &Idx, &Host, &Sch, &DBUser, &NS, &DB, &Date)
			t.AddLine(pid, Idx, Host, Sch, DBUser, NS, DB, Date)
		}
		print("\n")
		t.Print()
		print("\n")
	}
}

func DBValidIndex(i string) bool {
	var pid int = 0
	rw := DBconn.QueryRow("SELECT pid FROM Profile WHERE Idx=?;", i)
	err := rw.Scan(&pid)
	if err != nil {
		return false
	} else {
		if pid == 0 {
			return false
		} else {
			return true
		}
	}
}

func DBDropIdx(i string) {
	x := DBValidIndex(i)
	if !x {
		PrintErr("Profile do not exists.")
		return
	}
	sttm, err := DBconn.Prepare("DELETE FROM Profile WHERE Idx=?")
	if err != nil {
		HandErrs(err)
	} else {
		sttm.Exec(i)
		PrintSuc(i + " deleted.")
	}
}
