package api

import (
	"fmt"
	"net/http"
)

func Ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "PONG")
}
