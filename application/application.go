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
	CheckSpecialCharAndLettersQuantity(password *entity.PasswordGen) bool
	CheckSpecialCharAndNumbersQuantity(password *entity.PasswordGen) bool
	CheckLettersAndNumbersQuantity(password *entity.PasswordGen) bool
	CheckAllCharsQuantity(password *entity.PasswordGen) bool
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
		if !p.passService.CheckAllCharsQuantity(password) {
			p.GeneratePassword(password)
		}
	case password.HasSpecialChar && password.HasLetter:
		if !p.passService.CheckSpecialCharAndLettersQuantity(password) {
			p.GeneratePassword(password)
		}
	case password.HasSpecialChar && password.HasNumber:
		if !p.passService.CheckSpecialCharAndNumbersQuantity(password) {
			p.GeneratePassword(password)
		}
	case password.HasLetter && password.HasNumber:
		if !p.passService.CheckLettersAndNumbersQuantity(password) {
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

func (p *passwordGeneratorApp) CheckSpecialCharAndLettersQuantity(password *entity.PasswordGen) bool {
	return p.passService.CheckSpecialCharAndLettersQuantity(password)
}

func (p *passwordGeneratorApp) CheckSpecialCharAndNumbersQuantity(password *entity.PasswordGen) bool {
	return p.passService.CheckSpecialCharAndNumbersQuantity(password)
}

func (p *passwordGeneratorApp) CheckLettersAndNumbersQuantity(password *entity.PasswordGen) bool {
	return p.passService.CheckLettersAndNumbersQuantity(password)
}

func (p *passwordGeneratorApp) CheckAllCharsQuantity(password *entity.PasswordGen) bool {
	return p.passService.CheckAllCharsQuantity(password)
}

func (p *passwordGeneratorApp) CheckCharConsiderations(password entity.PasswordGen) []rune {
	return p.passService.CheckCharConsiderations(password)
}
