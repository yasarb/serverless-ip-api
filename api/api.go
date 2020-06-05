package api

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func Home(w http.ResponseWriter, req *http.Request) {

	ipAddr := ""
	headers := [9]string{
		"HTTP_CLIENT_IP",
		"HTTP_X_FORWARDED_FOR",
		"HTTP_X_FORWARDED",
		"HTTP_X_CLUSTER_CLIENT_IP",
		"HTTP_FORWARDED_FOR",
		"HTTP_FORWARDED",
		"HTTP_CF_CONNECTING_IP",
		"X-FORWARDED-FOR",
		"REMOTE_ADDR"}

	for _, key := range headers {
		val := req.Header.Get(key)

		if val != "" {
			val = strings.Split(val, ",")[0]
			ipAddr = net.ParseIP(val).String()
			break
		}
	}

	if ipAddr == "" {
		// still not found
		val := strings.Split(req.RemoteAddr, ":")[0]
		ipAddr = net.ParseIP(val).String()
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, ipAddr)
}

func Ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "PONG")
}
