package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"serverless-ip/api"
)

func main() {
	host := os.Getenv("LISTEN_ADDR")
	port := os.Getenv("PORT")
	addr := net.JoinHostPort(host, port)

	http.HandleFunc("/ping", api.Ping)

	log.Println("Starting HTTP server at", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
