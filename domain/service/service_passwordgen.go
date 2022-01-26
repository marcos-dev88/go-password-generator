package service

import (
	"errors"
	"math/rand"
	"sync"
	"time"

	"github.com/marcos-dev88/go-password-generator/domain/entity"
)

type Service interface {
	GeneratePasswordByLength(length int, passCharacters []rune) (string, error)
	CheckSpecialCharAndLettersQuantity(password *entity.PasswordGen) bool
	CheckSpecialCharAndNumbersQuantity(password *entity.PasswordGen) bool
	CheckLettersAndNumbersQuantity(password *entity.PasswordGen) bool
	CheckAllCharsQuantity(password *entity.PasswordGen) bool
	CheckCharConsiderations(password entity.PasswordGen) []rune
}

type service struct {
	passGen entity.PasswordGenerator
}

func NewService(passGen entity.PasswordGenerator) *service {
	return &service{passGen: passGen}
}

// GeneratePasswordByLength - It generates a random password by defined length and your own consideration of chars to make it
func (s *service) GeneratePasswordByLength(length int, passCharacters []rune) (string, error) {
	rand.Seed(time.Now().UnixNano())
	randomCharArray := make([]rune, length)

	// Channels
	inCH, outCH, returnCH, errorCH, errDoneCH := make(chan string, 3), make(chan string, 3), make(chan string, 3), make(chan error), make(chan error)
	passwordListChannel := make(chan []string)

	wg := sync.WaitGroup{}
	wg.Add(2)
	defer wg.Wait()

	// This var receives a function that gets the password's condition and generates them
	passGenerator := func(passwordLength int, passwordCharacters []rune) (string, error) {
		for i := 0; i < len(randomCharArray); i++ {
			randomCharArray[i] = passCharacters[rand.Int63n(int64(len(passCharacters)))]
		}

		if len(string(randomCharArray)) == 0 {
			return "", errors.New("password is empty")
		}

		return string(randomCharArray), nil
	}

	// Generates a random number
	random := func(min, max int) int {
		return min + rand.Intn(max-min)
	}

	// Generating a random number one to fifteen
	generatedRandom := random(1, 15)

	go func() {
		for i := 0; i < generatedRandom; i++ {
			pass, err := passGenerator(length, passCharacters)
			if err != nil {
				errorCH <- err
			}
			inCH <- pass
		}
		close(inCH)
	}()

	go removeDuplicatedPasswords(inCH, outCH)
	go getPasswords(passwordListChannel, outCH, &wg)
	go returnAllPasswords(passwordListChannel, returnCH, &wg)
	go chanErrorHandler(errorCH, errDoneCH, &wg)

	select {
	case generatedPassword := <-returnCH:
		return generatedPassword, nil
	case generatedPassword2 := <-returnCH:
		return generatedPassword2, nil
	case generatedPassword3 := <-returnCH:
		return generatedPassword3, nil
	case generatedPassword4 := <-returnCH:
		return generatedPassword4, nil
	case generatedPassword5 := <-returnCH:
		return generatedPassword5, nil
	case returnError := <-errDoneCH:
		return "", returnError
	}
}

// Checking the duplicated passwords and removing them
func removeDuplicatedPasswords(inputChannel, outputChannel chan string) {
	var previousPassword string
	for actualPassword := range inputChannel {
		if actualPassword != previousPassword {
			previousPassword = actualPassword
			outputChannel <- actualPassword
		}
	}
	close(outputChannel)
}

// getPasswords: Get all generated passswords and send it to an array
func getPasswords(generatedPasswods chan []string, outputSlice chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	var passwordList []string

	for v := range outputSlice {
		passwordList = append(passwordList, v)
	}

	generatedPasswods <- passwordList
}

func returnAllPasswords(passwordList chan []string, receiveCh chan string, newWg *sync.WaitGroup) {
	defer newWg.Done()

	go func() {
		for _, v := range <-passwordList {
			receiveCh <- v
		}
		close(receiveCh)
	}()
}

// Get errors from CH, this way we could handle errors
func chanErrorHandler(errorInput chan error, errorOutput chan error, newWg *sync.WaitGroup) {
	defer newWg.Done()
	select {
	case err := <-errorInput:
		errorOutput <- err
	}
}

// CheckSpecialCharAndLettersQuantity - It Checks password's special characters and letters according to its length
func (s *service) CheckSpecialCharAndLettersQuantity(password *entity.PasswordGen) bool {
	passwordLetters := s.passGen.GetPasswordLetters(password.Password)
	passwordSpecialChars := s.passGen.GetPasswordSpecialChars(password.Password)

	switch password.Length {
	case 8:
		if len(passwordLetters) < 3 && len(passwordSpecialChars) < 3 {
			return false
		}
	case 16:
		if len(passwordLetters) < 5 && len(passwordSpecialChars) < 5 {
			return false
		}
	case 32:
		if len(passwordLetters) < 6 && len(passwordSpecialChars) < 6 {
			return false
		}
	}
	return true
}

// CheckSpecialCharAndNumbersQuantity - It Checks password's special characters and numbers according to its length
func (s *service) CheckSpecialCharAndNumbersQuantity(password *entity.PasswordGen) bool {
	passwordNumbers := s.passGen.GetPasswordNumbers(password.Password)
	passwordSpecialChars := s.passGen.GetPasswordSpecialChars(password.Password)

	switch password.Length {
	case 8:
		if len(passwordNumbers) < 3 && len(passwordSpecialChars) < 3 {
			return false
		}
	case 16:
		if len(passwordNumbers) < 5 && len(passwordSpecialChars) < 5 {
			return false
		}
	case 32:
		if len(passwordNumbers) < 6 && len(passwordSpecialChars) < 6 {
			return false
		}
	}
	return true
}

// CheckLettersAndNumbersQuantity - It Checks password's letters and numbers according to its length
func (s *service) CheckLettersAndNumbersQuantity(password *entity.PasswordGen) bool {
	passwordNumbers := s.passGen.GetPasswordNumbers(password.Password)
	passwordLetters := s.passGen.GetPasswordLetters(password.Password)

	switch password.Length {
	case 8:
		if len(passwordNumbers) < 3 && len(passwordLetters) < 3 {
			return false
		}
	case 16:
		if len(passwordNumbers) < 5 && len(passwordLetters) < 5 {
			return false
		}
	case 32:
		if len(passwordNumbers) < 6 && len(passwordLetters) < 6 {
			return false
		}
	}
	return true
}

// CheckAllCharsQuantity - It Checks password's criteria of all kind of characters
func (s *service) CheckAllCharsQuantity(password *entity.PasswordGen) bool {
	passwordNumbers := s.passGen.GetPasswordNumbers(password.Password)
	passwordLetters := s.passGen.GetPasswordLetters(password.Password)
	passwordSpecialChars := s.passGen.GetPasswordSpecialChars(password.Password)

	switch password.Length {
	case 8:
		if len(passwordNumbers) < 2 && len(passwordLetters) < 2 && len(passwordSpecialChars) < 2 {
			return false
		}
	case 16:
		if len(passwordNumbers) < 3 && len(passwordLetters) < 3 && len(passwordSpecialChars) < 3 {
			return false
		}
	case 32:
		if len(passwordNumbers) < 6 && len(passwordLetters) < 6 && len(passwordSpecialChars) < 6 {
			return false
		}
	}
	return true
}

// CheckCharConsiderations - It Checks password's considerations, this way, check which characters it will have
func (s *service) CheckCharConsiderations(password entity.PasswordGen) []rune {

	switch {
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
