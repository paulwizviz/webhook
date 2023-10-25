// Package hdlutil containers decorator functions to validate common http handler
package hdlutil

import (
	"log"
	"net/http"
)

func MethodPostValidate(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println("Invalid METHOD")
			http.Error(rw, "Invalid HTTP method. Only POST requests are accepted.", http.StatusMethodNotAllowed)
			return
		}
		next(rw, r)
	}
}
