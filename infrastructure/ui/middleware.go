package ui

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Middleware interface {
	Auth(handler http.HandlerFunc) http.HandlerFunc
	EnablingCORS(handler http.HandlerFunc) http.HandlerFunc
}

type middleware struct{}

func NewMiddleware() *middleware {
	return &middleware{}
}

func (m *middleware) Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

		apiKey := req.Header.Get("api-token")
		if apiKey != os.Getenv("API_KEY") {
			returnMessage, _ := json.Marshal(map[string]string{"error": "api-key is missing or it isn't correct"})
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(returnMessage)
			return
		}
		handler.ServeHTTP(rw, req)
	}
}

func (m *middleware) EnablingCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		rw.Header().Set("Content-Type", "application-json")

		if req.Method != "POST" {
			log.Printf("method not allowed")
			returnMessage, _ := json.Marshal(map[string]string{"error": "method not allowed"})
			rw.WriteHeader(http.StatusMethodNotAllowed)
			rw.Write(returnMessage)
			return
		}

		rw.WriteHeader(http.StatusOK)
		handler.ServeHTTP(rw, req)
	}
}
