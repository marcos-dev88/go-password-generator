package main

import (
	"bufio"
	"github.com/google/uuid"
	"github.com/marcos-dev88/go-password-generator/application"
	"github.com/marcos-dev88/go-password-generator/domain/entity"
	service2 "github.com/marcos-dev88/go-password-generator/domain/service"
	"github.com/marcos-dev88/go-password-generator/infrastructure/persistence"
	"github.com/marcos-dev88/go-password-generator/infrastructure/ui"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	if err := defineEnvs(".env"); err != nil {
		log.Printf("Error to load .env -> %v", err)
	}
}

func main() {
	db := persistence.NewMongoDB("", "", "")
	repo := persistence.NewRepository(db)

	newUuid, err := uuid.NewUUID()

	if err != nil {
		panic(err)
	}

	passwordGen := entity.NewPasswordGen(
		newUuid,
		"",
		0,
		false,
		false,
		false,
	)

	service := service2.NewService(passwordGen)

	app := application.NewApplication(repo, service)

	handler := ui.NewHandler(app)

	log.Printf("\nServer is running at: %s", os.Getenv("API_PORT"))
	http.HandleFunc("/password-gen/", handler.HandlePasswordGenerator)
	log.Fatal(http.ListenAndServe(os.Getenv("API_PORT"), nil))

}

func defineEnvs(fileName string) error {

	envs := make(map[string]string)

	file, err := os.Open(fileName)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("We got an error -> %v", err)
		}
	}(file)

	if err != nil {
		return err
	}

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		envSplit := strings.SplitN(sc.Text(), "=", 2)
		if len(envSplit) > 1 {
			envs[envSplit[0]] = envSplit[1]
		}
	}

	for key, value := range envs {
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	return nil
}
