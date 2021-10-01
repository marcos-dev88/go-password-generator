package ui

import (
	"encoding/json"
	"net/http"
	"os"
)

type Middleware interface{
	Auth(handler http.HandlerFunc) http.HandlerFunc
}

type middleware struct {}

func NewMiddleware() *middleware{
	return &middleware{}
}

func (m *middleware) Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func (rw http.ResponseWriter, req *http.Request) {
		apiKey := req.Header.Get("api-token")
		if apiKey != os.Getenv("API_KEY") {

			returnMessage, _ := json.Marshal(map[string]string{"error": "api-key is missing or isn't correct"})
			rw.Header().Set("Content-Type", "application-json")
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(returnMessage)
			return
		}
		handler.ServeHTTP(rw, req)
	}
}
