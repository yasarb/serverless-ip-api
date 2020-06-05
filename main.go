package main

import (
	"github.com/yasarb/serverless-ip-api/api"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	host := os.Getenv("LISTEN_ADDR")
	port := os.Getenv("PORT")
	addr := net.JoinHostPort(host, port)

	http.HandleFunc("/ping", api.Ping)
	http.HandleFunc("/", api.Home)

	log.Println("Starting HTTP server at", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
