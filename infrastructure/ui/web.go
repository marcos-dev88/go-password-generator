package ui

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/marcos-dev88/go-password-generator/application"
	"github.com/marcos-dev88/go-password-generator/domain/entity"
	"log"
	"net/http"
)

type Handler interface {
	HandlePasswordGenerator(http.ResponseWriter, *http.Request)
}

type handler struct {
	app application.PasswordGeneratorApp
}

func NewHandler(app application.PasswordGeneratorApp) *handler {
	return &handler{app: app}
}

func (h *handler) HandlePasswordGenerator(rw http.ResponseWriter, req *http.Request){
	if req.Method != "POST" {
		log.Println("Method not allowed")
	}

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

	response, _ := json.Marshal(&generatedPassword)

	rw.Header().Set("Content-Type","application-json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(response)
}


