## HTTPDumper

Simple HTTP server that dumps request payloads to `stdout`

**install**
```
go get install github.com/juztin/httpdumper/cmd/httpdumper
```

**serve**
```
httpdumper

httpdumper -addr :4040
httpdumper -addr 0.0.0.0:4040
httpdumper -addr 127.0.0.1:4040
```

#### Custom

If you'd like to print some additional, custom information, create a server with 

```
func customFn1(r *http.Request) {
  // Do some custom request printing, like validating API signature, etc.
}

func customFn2(r *http.Request) {
  // More custom stuff
}


dumper, err := httpdumper.New(":4040", customFn1, customFn2)
```
