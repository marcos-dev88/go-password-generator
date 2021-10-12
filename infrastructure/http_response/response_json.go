package http_response

import (
	"encoding/json"
	"net/http"
)

type ResponseHTTP interface {
	ResponseJSON(rw http.ResponseWriter, resp *responseHttp)
	ErrorJSON(rw http.ResponseWriter, ce customError)
}

type responseHttp struct {
	StatusCode int
	Payload interface{}
}

func NewResponseHTTP(statusCode int, payload interface{}) *responseHttp {
	return &responseHttp{StatusCode: statusCode, Payload: payload}
}

func (rh *responseHttp) ResponseJSON(rw http.ResponseWriter, resp *responseHttp) {

	response, err := json.Marshal(&resp.Payload)

	if err != nil {
		rh.ErrorJSON(rw, *NewCustomError(GetErrorStatus(err), GetErrorMessage(err)))
	}

	rw.Header().Set("Content-Type", "application-json")
	rw.WriteHeader(resp.StatusCode)
	rw.Write(response)
}

func (rh *responseHttp) ErrorJSON(rw http.ResponseWriter, ce customError) {
	ce.DefaultLogResponse()
	rh.ResponseJSON(rw, &responseHttp{
		StatusCode: ce.StatusCode ,
		Payload:    map[string]interface{}{"error": ce.Message, "status": ce.StatusCode},
	})
}
