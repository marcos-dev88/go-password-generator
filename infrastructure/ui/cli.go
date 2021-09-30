package ui

import (
	"flag"
	"fmt"
	"github.com/marcos-dev88/go-password-generator/application"
	"github.com/marcos-dev88/go-password-generator/domain/entity"
	"log"
	"os"
)

type CLI interface {
	GeneratePassword()
}

type cliGeneratePassword struct {
	app application.PasswordGeneratorApp
}

func NewCli(app application.PasswordGeneratorApp) *cliGeneratePassword {
	return &cliGeneratePassword{app: app}
}

func (c *cliGeneratePassword) GeneratePassword() {

	if len(os.Args) <= 1 {
		log.Fatal("error: input not valid, you must use: generate or server")
	}

	strongPass := flag.NewFlagSet("genstrong", flag.ExitOnError)
	generate := flag.NewFlagSet("generate", flag.ExitOnError)
	passLength := generate.Int("l", 0, "inform the length of your password")

	switch os.Args[1] {
	case "genstrong":
		if err := strongPass.Parse(os.Args[2:]); err != nil {
			log.Printf("error: %v", err)
			os.Exit(1)
		}

		generatedStrongPass, _ := c.app.GeneratePasswordByLength(64, entity.AllCharacters)

		log.Printf("generated strong password: \n\n\t%v", generatedStrongPass)
		fmt.Println("")
		os.Exit(0)

	case "generate":
		if err := generate.Parse(os.Args[2:]); err != nil {
			log.Printf("error: %v", err)
			os.Exit(1)
		}

		pass := entity.NewPasswordGen("", "", *passLength, true, true, true)

		result, err := c.app.GeneratePassword(pass)

		if err != nil {
			log.Printf("error: %v", err)
			os.Exit(1)
		}

		log.Printf("generated password: \n\n\t%v", result.Password)
		fmt.Println("")
		os.Exit(0)

	case "server":
		return
	default:
		log.Println("error: input not valid, you must use: search or server")
		os.Exit(1)
	}
}