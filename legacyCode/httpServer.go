package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello web!" + time.Now().String()))
	fmt.Fprintf(w, "hello web:%s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
