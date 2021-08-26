package application

import (
	"github.com/marcos-dev88/go-password-generator/domain/entity"
	"github.com/marcos-dev88/go-password-generator/domain/repository"
	"github.com/marcos-dev88/go-password-generator/domain/service"
)

type PasswordGeneratorApp interface {
	GetPasswordGen(password string) (*entity.PasswordGen, error)
	CreatePasswordGen(*entity.PasswordGen) (*entity.PasswordGen, error)
	PasswordExists(entity.PasswordGen) bool
	GeneratePasswordByLength(length int, passCharacters []rune) (string, error)
	CheckSpecialCharAndLettersQuantity(password *entity.PasswordGen) error
	CheckSpecialCharAndNumbersQuantity(password *entity.PasswordGen) error
	CheckAllCharsQuantity(password *entity.PasswordGen) error
	CheckCharConsiderations(password entity.PasswordGen) []rune
	GeneratePassword(password *entity.PasswordGen) (*entity.PasswordGen, error)
}

type passwordGeneratorApp struct {
	passGenEntity entity.PasswordGenerator
	passGenRepo   repository.Repository
	passService   service.Service
}

func NewApplication(passGenRepo repository.Repository, passService service.Service) *passwordGeneratorApp {
	return &passwordGeneratorApp{passGenRepo: passGenRepo, passService: passService}
}

func (p *passwordGeneratorApp) GeneratePassword(password *entity.PasswordGen) (*entity.PasswordGen, error) {
	//Getting password configs of kind of password what client want
	passwordChars := p.passService.CheckCharConsiderations(*password)

	//Generating a password by length
	generatedPass, err := p.passService.GeneratePasswordByLength(password.Length, passwordChars)
	
	password.Password = generatedPass
	
	if err != nil {
		return nil, err
	}

	// Checking whether this password is secure
	switch {
	case password.HasLetter && password.HasNumber && password.HasSpecialChar:
		err := p.passService.CheckAllCharsQuantity(password)
		if err != nil {
			p.GeneratePassword(password)
		}
	case password.HasSpecialChar && password.HasLetter:
		err := p.passService.CheckSpecialCharAndLettersQuantity(password)
		if err != nil {
			p.GeneratePassword(password)
		}
	case password.HasSpecialChar && password.HasNumber:
		err := p.passService.CheckSpecialCharAndNumbersQuantity(password)
		if err != nil {
			p.GeneratePassword(password)
		}
	case password.HasLetter && password.HasNumber:
		err := p.passService.CheckLettersAndNumbersQuantity(password)
		if err != nil {
			p.GeneratePassword(password)
		}
	}

	return password, nil
}

func (p *passwordGeneratorApp) GetPasswordGen(password string) (*entity.PasswordGen, error) {
	return p.passGenRepo.GetPasswordGen(password)
}

func (p *passwordGeneratorApp) CreatePasswordGen(password *entity.PasswordGen) (*entity.PasswordGen, error) {
	return p.passGenRepo.CreatePasswordGen(password)
}

func (p *passwordGeneratorApp) PasswordExists(password entity.PasswordGen) bool {
	return p.passGenRepo.PasswordExists(password)
}

func (p *passwordGeneratorApp) GeneratePasswordByLength(length int, passCharacters []rune) (string, error) {
	return p.passService.GeneratePasswordByLength(length, passCharacters)
}

func (p *passwordGeneratorApp) CheckSpecialCharAndLettersQuantity(password *entity.PasswordGen) error {
	return p.passService.CheckSpecialCharAndLettersQuantity(password)
}

func (p *passwordGeneratorApp) CheckSpecialCharAndNumbersQuantity(password *entity.PasswordGen) error {
	return p.passService.CheckSpecialCharAndNumbersQuantity(password)
}

func (p *passwordGeneratorApp) CheckAllCharsQuantity(password *entity.PasswordGen) error {
	return p.passService.CheckAllCharsQuantity(password)
}

func (p *passwordGeneratorApp) CheckCharConsiderations(password entity.PasswordGen) []rune {
	return p.passService.CheckCharConsiderations(password)
}
