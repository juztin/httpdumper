package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/juztin/httpdumper"
)

func main() {
	addr := flag.String("addr", ":4040", "Address to listen on (default :4040)")
	flag.Parse()
	dumper := httpdumper.New(nil)
	fmt.Printf("listening on %s...\n", *addr)
	http.ListenAndServe(*addr, dumper)
}
