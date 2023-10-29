// Package hdlutil containers decorator functions to validate common http handler
package hdlutil

import (
	"log"
	"net/http"
)

type validateFunc func(http.ResponseWriter, *http.Request)

func (v validateFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	v(rw, r)
}

func ValidateGetMethod(next http.Handler) http.Handler {
	var fn validateFunc = func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Println("Invalid METHOD")
			http.Error(rw, "Invalid HTTP method. Only POST requests are accepted.", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(rw, r)
	}
	return fn
}

func ValidatePostMethod(next http.Handler) http.Handler {
	var fn validateFunc = func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println("Invalid METHOD")
			http.Error(rw, "Invalid HTTP method. Only POST requests are accepted.", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(rw, r)
	}
	return fn
}

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

func MethodGetValidate(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Println("Invalid METHOD")
			http.Error(rw, "Invalid HTTP method. Only GET requests are accepted.", http.StatusMethodNotAllowed)
			return
		}
		next(rw, r)
	}
}
