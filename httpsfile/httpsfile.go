package main

import (
	"net/http"
	"io"
	"log"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi, This is an demo of https serve in goLang!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServeTLS(":8090", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}