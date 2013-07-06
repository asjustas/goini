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

	fmt.Print( conf.Int("database", "port") )

}