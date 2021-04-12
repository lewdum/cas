package main

import (
	"fmt"
	"log"
	"net/http"
)

func info(r *http.Request, message string) {
	log.Printf("[%s] %s\n", r.RemoteAddr, message)
}

func fail(w http.ResponseWriter, r *http.Request, status int, message string) {
	info(r, fmt.Sprintf("Error: %s (%d)", message, status))

	w.WriteHeader(status)
	w.Write([]byte(message))
}

func failInternal(w http.ResponseWriter, r *http.Request, err error) {
	info(r, fmt.Sprintf("Internal Error: %v", err))

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
}
