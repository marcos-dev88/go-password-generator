package ui

import (
	"github.com/marcos-dev88/go-password-generator/infrastructure/http_response"
	"net/http"
	"os"
)

type Middleware interface {
	Auth(handler http.HandlerFunc) http.HandlerFunc
	EnablingCORS(handler http.HandlerFunc) http.HandlerFunc
}

type middleware struct{
	json_resp http_response.ResponseHTTP
}

func NewMiddleware(json_resp http_response.ResponseHTTP) *middleware {
	return &middleware{json_resp: json_resp}
}

func (m *middleware) Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

		apiKey := req.Header.Get("api-token")
		if apiKey != os.Getenv("API_KEY") {
			m.json_resp.ErrorJSON(rw, *http_response.NewCustomError(http.StatusUnauthorized, "api-key is missing or it isn't correct"))
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
			m.json_resp.ErrorJSON(rw, *http_response.NewCustomError(http.StatusMethodNotAllowed, "method not allowed"))
			return
		}
		handler.ServeHTTP(rw, req)
	}
}
