# SurrealCLI
SurrealDB CLI tool.

```bash
git clone https://github.com/farinap5/SurrealCLI.git
cd SurrealCLI
go build SurrealCLI
```

Normal connection:
```
surrcli -host "0.0.0.0:80" -u elf -ns surr -db surr -comp 0              
[password]: 
######  SurrealCLI  ######
Type `.help` for help meu.
v 0.3-NotStable
[OK]- Connection is OK!

[surr]> INFO FOR DB
[
  {
    "time": "273.293µs",
    "status": "OK",
    "result": {
...
```

### Help Menu
```
─$ surrcli --help

COMMAND   DESCRIPTION                 DEFAULT
-------   -----------                 -------
-u        Username                    root
-p        Password                    hide password
-host     Database address "IP:PORT"  0.0.0.0:80
-NS       Namespace                   surr
-DB       Database                    surr
-sc       Schema                      http
-profile  Connect to a profile        none
-t        Connection timeout          5
-pretty   Pretty output               true
-comp     Number of suggestions       5
```


```
[surr]> .help

COMMAND   DESCRIPTION
-------   -----------
.help     Show help menu
.options  Env variables
.set      Set variable
.save     Save profile|query
.show     Show profiles|queries
.delete   Delete profile|query
.run      Run profile|query
```


Authenticate with saved profile using `-profile`
```
╰─$ surrcli -profile anyprofile -p $PXX
######  SurrealCLI  ######
Type `.help` for help meu.
0.3-NotStable
[OK]- Connection is OK!

[surr]> .options

VARIABLE    VALUE
--------    -----
Host        0.0.0.0:80
User        elf
Namespace   surr
Database    surr
Schema      http
Pretty      true
Timeout     5ns
Suggestion  5

[surr]> 
```

Profile does not keep password, it is needed to be passed as parameter.
```
╰─$ surrcli -profile test -p $PXX -q "INFO FOR DB;"
[
  {
    "time": "271.87µs",
    "status": "OK",
    "result": {
      "dl": {},
      "dt": {},
...
```