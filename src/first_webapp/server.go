package main

import (
	"fmt"
	"net/http"
)

func handler(writter http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writter, "Hellow Word, %s!", request.URL.Path[1:])

}


func main() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":9090", nil)
}