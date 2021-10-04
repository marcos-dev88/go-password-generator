package ui

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/marcos-dev88/go-password-generator/application"
	"github.com/marcos-dev88/go-password-generator/domain/entity"
	"github.com/marcos-dev88/go-password-generator/infrastructure/http_response"
	"net/http"
)

type Handler interface {
	HandlePasswordGenerator(http.ResponseWriter, *http.Request)
}

type handler struct {
	app application.PasswordGeneratorApp
	json_resp http_response.ResponseHTTP
	json_err http_response.CustomError
}

func NewHandler(app application.PasswordGeneratorApp, json_resp http_response.ResponseHTTP, json_err http_response.CustomError) *handler {
	return &handler{app: app, json_resp: json_resp, json_err: json_err}
}

func (h *handler) HandlePasswordGenerator(rw http.ResponseWriter, req *http.Request){

	decoder := json.NewDecoder(req.Body)

	newUuid, err := uuid.NewUUID()

	if err != nil {
		h.defaultErrorResponse(rw, err)
		return
	}

	password := entity.NewPasswordGen(newUuid.String(), "", 0, false, false, false)

	if err := decoder.Decode(&password); err != nil {
		h.defaultErrorResponse(rw, err)
		return
	}

	generatedPassword, err := h.app.GeneratePassword(password)

	if err != nil {
		h.defaultErrorResponse(rw, err)
		return
	}

	h.json_resp.ResponseJSON(rw, http_response.NewResponseHTTP(http.StatusOK, generatedPassword))
}

func (h *handler) defaultErrorResponse(rw http.ResponseWriter, err error) {
	newErr := http_response.NewCustomError(http.StatusInternalServerError, err.Error())
	h.json_err.DefaultLogResponse()
	h.json_resp.ErrorJSON(rw, *newErr)
}


