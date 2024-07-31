package main

import (
	"flag"

	"github.com/go-the-way/cpt"
)

var (
	addr = flag.String("addr", ":9988", "The cpt embedded http server bind address")
)

func main() {
	flag.Parse()
	cpt.Serve(*addr)
}
