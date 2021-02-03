package httpdumper

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type RequestFunc func(*http.Request)

type Dumper struct {
	addr         string
	requestFuncs []RequestFunc
}

func (d *Dumper) Listen() error {
	return http.ListenAndServe(d.addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("------------------------------------Request-------------------------------------")
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		fmt.Print(string(dump))
		for _, fn := range d.requestFuncs {
			if fn != nil {
				fn(r)
			}
		}
	}))
}

func New(addr string, funcs ...RequestFunc) *Dumper {
	return &Dumper{addr, funcs}
}
