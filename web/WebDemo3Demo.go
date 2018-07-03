package main

import (
	"net/http"
	"fmt"
)

type MyMux struct {
}
//实现这个Handler接口   位于在server.go 包中
//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}


