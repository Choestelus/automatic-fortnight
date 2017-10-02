package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("such is")
	fmt.Println(AddNumber(2, 3))
	fmt.Fprintf(os.Stdout, "%v\n", "test")
	mux := http.NewServeMux()

	mux.HandleFunc("/foo", FooHandler)
	mux.HandleFunc("/", IndexHandler)

	s := &http.Server{
		Addr:    ":9000",
		Handler: mux,
	}
	log.Printf("Listening on %v", s.Addr)
	log.Fatal(s.ListenAndServe())
}
