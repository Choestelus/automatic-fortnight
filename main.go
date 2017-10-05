package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Printf("Initializing...")
	mux := http.NewServeMux()
	time.Sleep(20 * time.Second)
	fmt.Printf("done\n")

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
