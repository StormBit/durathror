package main

import (
	"flag"
	"log"
)

var (
	verbose = flag.Bool("verbose", false, "Verbosity")
)

func main() {

	flag.Parse()
	log.Printf("%s v%s starting up.", NAME, VERSION)
}
