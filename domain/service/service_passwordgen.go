package service

import (
	"errors"
	"github.com/marcos-dev88/go-password-generator/domain/entity"
	"math/rand"
	"sync"
	"time"
)

type Service interface {
	GeneratePasswordByLength(length int, passCharacters []rune) (string, error)
	GenerateRandomPassword() string
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

func (s *service) GenerateRandomPassword() string {
	rand.Seed(time.Now().UTC().UnixNano())
	var passwordList []string
	inputCh := make(chan string, 3)
	outputCh := make(chan string, 3)
	ch := make(chan string, 3)
	wg := sync.WaitGroup{}

	wg.Add(1)

	defer wg.Wait()

	// Generates a random number
	random := func(min, max int) int {
		return min + rand.Intn(max-min)
	}

	getPasswords := func(outputChan <-chan string, newWg *sync.WaitGroup) []string {
		defer newWg.Done()
		for v := range outputCh {
			passwordList = append(passwordList, v)
		}
		return passwordList
	}

	// Checking the duplicated passwords and removing them
	removeDuplicatedPasswords := func(inputChan chan string, outputChan chan string) {
		var previousPassword string
		for actualPassword := range inputChan {
			if actualPassword != previousPassword {
				previousPassword = actualPassword
				outputChan <- actualPassword
			}
		}
		close(outputChan)
	}

	go getPasswords(outputCh, &wg)

	go removeDuplicatedPasswords(inputCh, outputCh)

	// Generating a random number one to fifteen
	generatedRandom := random(1, 15)

	go func() {
		for i := 0; i < generatedRandom; i++ {
			pass, err := s.GeneratePasswordByLength(32, entity.AllCharacters)
			if err != nil {
				panic(err)
			}
			inputCh <- pass
		}
		close(inputCh)
	}()

	go func() {
		for {
			select {
			case generatedPassword := <-outputCh:
				ch <- generatedPassword
			case generatedPassword2 := <-outputCh:
				ch <- generatedPassword2
			case generatedPassword3 := <-outputCh:
				ch <- generatedPassword3
			case generatedPassword4 := <-outputCh:
				ch <- generatedPassword4
			case generatedPassword5 := <-outputCh:
				ch <- generatedPassword5
			}
		}
		//close(ch)
	}()

	// Saving return of output in a var to close ch
	returnCh :=  <-ch

	//close(ch)

	return returnCh
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
