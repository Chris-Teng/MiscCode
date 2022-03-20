package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json") //为了做前后端分离，设置后端的返回值格式为json
	w.Write([]byte(fmt.Sprintf(`{"time":"%s"}`, time.Now().String())))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", Hello)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
