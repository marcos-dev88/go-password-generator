package main

import (
	"bufio"
	"fmt"
	"github.com/marcos-dev88/go-password-generator/application"
	"github.com/marcos-dev88/go-password-generator/domain/entity"
	"github.com/marcos-dev88/go-password-generator/domain/service"
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

var mongodbURI string = fmt.Sprintf(
	"mongodb://%v:%v@%v",
	os.Getenv("MONGO_USERNAME"),
	os.Getenv("MONGO_PASSWORD"),
	os.Getenv("MONGO_URL"),
)

func main() {
	db := persistence.NewMongoDB(
		os.Getenv("MONGO_DB"),
		os.Getenv("MONGO_TABLE"),
		mongodbURI,
	)

	repo := persistence.NewRepository(db)

	passwordGen := entity.NewPasswordGen(
		"",
		"",
		0,
		false,
		false,
		false,
	)

	servicePass := service.NewService(passwordGen)

	app := application.NewApplication(repo, servicePass)

	cli := ui.NewCli(app)
	cli.GeneratePassword()

	router := http.NewServeMux()
	handler := ui.NewHandler(app)
	middleware := ui.NewMiddleware()

	log.Printf("\nServer is running at: %s", os.Getenv("API_PORT"))
	router.HandleFunc("/password-gen/", middleware.EnablingCORS(middleware.Auth(handler.HandlePasswordGenerator)))
	log.Fatal(http.ListenAndServe(os.Getenv("API_PORT"), router))

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
