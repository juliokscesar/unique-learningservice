package utils

import (
	"log"
	"net/http"
)

func LogRequest(r *http.Request) {
	log.Println(r.Method, r.URL.Path, "by", r.RemoteAddr)
}
