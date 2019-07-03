package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	port := ":8081"
	log.Println("listen on", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
		panic(err)
	} 
}
