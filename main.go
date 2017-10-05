package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	time.Sleep(20 * time.Second)
	mux := http.NewServeMux()

	mux.HandleFunc("/foo", FooHandler)
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/health", LivenessHandler)
	mux.HandleFunc("/fail", InternalServerErrorhandler)

	s := &http.Server{
		Addr:    ":9000",
		Handler: mux,
	}

	log.Printf("Listening on %v", s.Addr)
	log.Fatal(s.ListenAndServe())
}
