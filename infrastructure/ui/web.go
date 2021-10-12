package ui

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/marcos-dev88/go-password-generator/application"
	"github.com/marcos-dev88/go-password-generator/domain/entity"
	"github.com/marcos-dev88/go-password-generator/infrastructure/http_response"
	"io"
	"net/http"
)

type Handler interface {
	HandlePasswordGenerator(http.ResponseWriter, *http.Request)
}

type handler struct {
	app application.PasswordGeneratorApp
	json_resp http_response.ResponseHTTP
}

func NewHandler(app application.PasswordGeneratorApp, json_resp http_response.ResponseHTTP) *handler {
	return &handler{app: app, json_resp: json_resp}
}

func (h *handler) HandlePasswordGenerator(rw http.ResponseWriter, req *http.Request){

	body, err := io.ReadAll(req.Body)

	if err := h.app.Validate(body); err != nil{
		h.defaultErrorResponse(rw, err)
		return
	}

	newUuid, err := uuid.NewUUID()

	if err != nil {
		h.defaultErrorResponse(rw, err)
		return
	}

	password := entity.NewPasswordGen(newUuid.String(), "", 0, false, false, false)

	err = json.Unmarshal(body, password)

	if err != nil {
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
	h.json_resp.ErrorJSON(rw, *newErr)
}


