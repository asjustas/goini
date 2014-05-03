## GoIni
GoIni is simple .INI file parser in golang.

## Installation
============
    go get github.com/asjustas/goini

## Usage
Ini file:

```ini
; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
organization = Acme Widgets Inc.
 
[database]
; use IP address in case network name resolution is not working
server = 192.0.2.62     
port = 143
file = payroll.dat
```
Example program:

```go
package main

import (
	"github.com/asjustas/goini"
	"fmt"
)

func main() {
	conf, err := goini.Load("test.ini")

	if err != nil {
		panic(err)
	}

	fmt.Println( conf.Str("owner", "name") )
	fmt.Println( conf.Str("owner", "organization") )

	fmt.Println( conf.Str("database", "server") )
	fmt.Println( conf.Int("database", "port") )
	fmt.Print( conf.Str("database", "file") )
}
```
## API Docs
[http://godoc.org/github.com/asjustas/goini](http://godoc.org/github.com/asjustas/goini)