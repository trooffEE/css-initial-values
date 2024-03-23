package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", viewApp)

	fmt.Println("CSSInitial application server started, serving port 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
