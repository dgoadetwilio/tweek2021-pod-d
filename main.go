package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("[hello]\n")

	//headersLocal(w, req)
	fmt.Fprintf(w, "hello world\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func headersLocal(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}
}

func main() {
	fmt.Printf("[main]\n")

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	fmt.Printf("[main] listening...\n")
	http.ListenAndServe(":8090", nil)

	fmt.Printf("[main] end\n")
}
