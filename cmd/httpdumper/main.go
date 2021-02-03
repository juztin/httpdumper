package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/juztin/httpdumper"
)

func main() {
	addr := flag.String("addr", ":4040", "Address to listen on (default :4040)")
	flag.Parse()
	dumper := httpdumper.New(*addr, nil)
	fmt.Printf("listening on %s...\n", *addr)
	log.Fatalln(dumper.Listen())
}
