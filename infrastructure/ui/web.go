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
	jsonresp http_response.ResponseHTTP
}

func NewHandler(app application.PasswordGeneratorApp, jsonresp http_response.ResponseHTTP) *handler {
	return &handler{app: app, jsonresp: jsonresp}
}

func (h *handler) HandlePasswordGenerator(rw http.ResponseWriter, req *http.Request){

	decoder := json.NewDecoder(req.Body)

	newUuid, err := uuid.NewUUID()

	if err != nil {
		panic(err)
	}

	password := entity.NewPasswordGen(newUuid.String(), "", 0, false, false, false)

	if err := decoder.Decode(&password); err != nil {
		panic(err)
	}

	generatedPassword, err := h.app.GeneratePassword(password)

	if err != nil {
		panic(err)
	}

	h.jsonresp.ResponseJSON(rw, http_response.NewResponseHTTP(http.StatusOK, generatedPassword))
}


