package httpdumper

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type RequestFunc func(http.ResponseWriter, *http.Request)

type Dumper struct {
	requestFuncs []RequestFunc
}

func (d *Dumper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
			fn(w, r)
		}
	}
}

func New(funcs ...RequestFunc) *Dumper {
	return &Dumper{funcs}
}
