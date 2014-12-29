package main

import (
	"github.com/codegangsta/negroni"
)

func main() {
	r := router()
	n := negroni.Classic()
	n.UseHandler(r)

	n.Run(":3000")
}
