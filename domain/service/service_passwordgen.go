package service

import (
	"errors"
	"github.com/marcos-dev88/go-password-generator/domain/entity"
	"math/rand"
	"time"
)

type Service interface {
	GeneratePasswordByLength(length int, passCharacters []rune) (string, error)
	CheckSpecialCharAndLettersQuantity(password *entity.PasswordGen) error
	CheckSpecialCharAndNumbersQuantity(password *entity.PasswordGen) error
	CheckLettersAndNumbersQuantity(password *entity.PasswordGen) error
	CheckAllCharsQuantity(password *entity.PasswordGen) error
	CheckCharConsiderations(password entity.PasswordGen) []rune
}

type service struct {
	passGen entity.PasswordGenerator
}

func NewService(passGen entity.PasswordGenerator) *service {
	return &service{passGen: passGen}
}

func (s *service) GeneratePasswordByLength(length int, passCharacters []rune) (string, error) {
	rand.Seed(time.Now().UnixNano())

	randomCharArray := make([]rune, length)

	for i := 0; i < len(randomCharArray); i++ {
		randomCharArray[i] = passCharacters[rand.Int63n(int64(len(passCharacters)))]
	}

	if len(string(randomCharArray)) == 0 {
		return "", errors.New("password is empty")
	}

	return string(randomCharArray), nil
}

func (s *service) CheckSpecialCharAndLettersQuantity(password *entity.PasswordGen) error {
	passwordLetters := s.passGen.GetPasswordLetters(password.Password)
	passwordSpecialChars := s.passGen.GetPasswordSpecialChars(password.Password)

	switch password.Length {
	case 8:
		if len(passwordLetters) < 3 {
			return errors.New("password has small quantity of letters")
		}
		if len(passwordSpecialChars) < 3 {
			return errors.New("password has small quantity of specialChars")
		}
	case 16:
		if len(passwordLetters) < 5 {
			return errors.New("password has small quantity of letters")
		}
		if len(passwordSpecialChars) < 5 {
			return errors.New("password has small quantity of specialChars")
		}
	case 32:
		if len(passwordLetters) < 6 {
			return errors.New("password has small quantity of letters")
		}
		if len(passwordSpecialChars) < 6 {
			return errors.New("password has small quantity of specialChars")
		}
	}

	return nil
}

func (s *service) CheckSpecialCharAndNumbersQuantity(password *entity.PasswordGen) error {
	passwordNumbers := s.passGen.GetPasswordNumbers(password.Password)
	passwordSpecialChars := s.passGen.GetPasswordSpecialChars(password.Password)

	switch password.Length {
	case 8:
		if len(passwordNumbers) < 3 {
			return errors.New("password has small quantity of numbers")
		}
		if len(passwordSpecialChars) < 3 {
			return errors.New("password has small quantity of specialChars")
		}
	case 16:
		if len(passwordNumbers) < 5 {
			return errors.New("password has small quantity of numbers")
		}
		if len(passwordSpecialChars) < 5 {
			return errors.New("password has small quantity of specialChars")
		}
	case 32:
		if len(passwordNumbers) < 6 {
			return errors.New("password has small quantity of numbers")
		}
		if len(passwordSpecialChars) < 6 {
			return errors.New("password has small quantity of specialChars")
		}
	}
	return nil
}

func (s *service) CheckLettersAndNumbersQuantity(password *entity.PasswordGen) error {
	passwordNumbers := s.passGen.GetPasswordNumbers(password.Password)
	passwordLetters := s.passGen.GetPasswordLetters(password.Password)

	switch password.Length {
	case 8:
		if len(passwordNumbers) < 3 {
			return errors.New("password has small quantity of letters")
		}
		if len(passwordLetters) < 3 {
			return errors.New("password has small quantity of letters")
		}
	case 16:
		if len(passwordNumbers) < 5 {
			return errors.New("password has small quantity of letters")
		}
		if len(passwordLetters) < 5 {
			return errors.New("password has small quantity of letters")
		}
	case 32:
		if len(passwordNumbers) < 6 {
			return errors.New("password has small quantity of letters")
		}
		if len(passwordLetters) < 6 {
			return errors.New("password has small quantity of letters")
		}
	}
	return nil

}

func (s *service) CheckAllCharsQuantity(password *entity.PasswordGen) error {
	passwordNumbers := s.passGen.GetPasswordNumbers(password.Password)
	passwordLetters := s.passGen.GetPasswordLetters(password.Password)
	passwordSpecialChars := s.passGen.GetPasswordSpecialChars(password.Password)

	switch password.Length {
	case 8:
		if len(passwordNumbers) < 2 {
			return errors.New("password has small quantity of letters")
		}
		if len(passwordLetters) < 2 {
			return errors.New("password has small quantity of letters")
		}
		if len(passwordSpecialChars) < 2 {
			return errors.New("password has small quantity of specialChars")
		}
	case 16:
		if len(passwordNumbers) < 3 {
			return errors.New("password has small quantity of letters")
		}
		if len(passwordLetters) < 3 {
			return errors.New("password has small quantity of letters")
		}
		if len(passwordSpecialChars) < 3 {
			return errors.New("password has small quantity of specialChars")
		}
	case 32:
		if len(passwordNumbers) < 6 {
			return errors.New("password has small quantity of letters")
		}
		if len(passwordLetters) < 6 {
			return errors.New("password has small quantity of letters")
		}
		if len(passwordSpecialChars) < 6 {
			return errors.New("password has small quantity of specialChars")
		}
	}
	return nil
}

func (s *service) CheckCharConsiderations(password entity.PasswordGen) []rune {

	switch true {
	case password.HasSpecialChar && password.HasLetter && password.HasNumber:
		return entity.AllCharacters
	case password.HasSpecialChar && password.HasLetter:
		return entity.SpecialCharAndLetters
	case password.HasSpecialChar && password.HasNumber:
		return entity.SpecialCharAndNumbers
	case password.HasLetter && password.HasNumber:
		return entity.LettersAndNumbers
	case password.HasLetter:
		return entity.Letters
	case password.HasNumber:
		return entity.Numbers
	case password.HasSpecialChar:
		return entity.SpecialCharacters
	}

	return nil
}
