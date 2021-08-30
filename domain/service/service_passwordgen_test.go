package service

import (
	"regexp"
	"testing"
)

func TestService_GeneratePasswordByLength(t *testing.T) {
	testService := NewService(&passwordEntityMock)

	t.Run("Generate_Password_Letters_Only", func(t *testing.T) {
		password, err := testService.GeneratePasswordByLength(64, lettersMock)

		if err != nil {
			t.Fatalf("We got an error to generate password to length -> %v", err)
		}

		findLetters, err := regexp.MatchString(regexLettersMock, password)

		if err != nil {
			t.Fatalf("We got an error to match strings in password -> %v", err)
		}

		if !findLetters{
			t.Fatalf("We expected a password with Letters, and we got %v", password)
		}

		if len(password) != 64 {
			t.Fatalf("We expected a password with length 64, and we got %v", len(password))
		}
	})

	t.Run("Generate_Password_Numbers_Only", func(t *testing.T) {
		password, err := testService.GeneratePasswordByLength(64, numbersMock)

		if err != nil {
			t.Fatalf("We got an error to generate password to length -> %v", err)
		}

		findNumbers, err := regexp.MatchString(regexNumbersMock, password)

		if err != nil {
			t.Fatalf("We got an error to match strings in password -> %v", err)
		}

		if !findNumbers{
			t.Fatalf("We expected a password with Letters, and we got %v", password)
		}

		if len(password) != 64 {
			t.Fatalf("We expected a password with length 64, and we got %v", len(password))
		}
	})

	t.Run("Generate_Password_SpecialChars_Only", func(t *testing.T) {
		password, err := testService.GeneratePasswordByLength(64, specialCharactersMock)

		if err != nil {
			t.Fatalf("We got an error to generate password to length -> %v", err)
		}

		findSpecialChars, err := regexp.MatchString(regexEspecialCharMock, password)

		if err != nil {
			t.Fatalf("We got an error to match strings in password -> %v", err)
		}

		if !findSpecialChars{
			t.Fatalf("We expected a password with Special Chars, and we got %v", password)
		}

		if len(password) != 64 {
			t.Fatalf("We expected a password with length 64, and we got %v", len(password))
		}
	})

	t.Run("Generate_Password_Letters_Numbers", func(t *testing.T) {
		password, err := testService.GeneratePasswordByLength(64, lettersAndNumbersMock)

		if err != nil {
			t.Fatalf("We got an error to generate password to length -> %v", err)
		}

		findLetters, err := regexp.MatchString(regexLettersMock, password)

		if err != nil {
			t.Fatalf("We got an error to match strings in password -> %v", err)
		}

		findNumbers, err := regexp.MatchString(regexNumbersMock, password)

		if err != nil {
			t.Fatalf("We got an error to match strings in password -> %v", err)
		}

		if !findLetters{
			t.Fatalf("We expected a password with letters, and we got %v", password)
		}

		if !findNumbers{
			t.Fatalf("We expected a password with Numbers, and we got %v", password)
		}

		if len(password) != 64 {
			t.Fatalf("We expected a password with length 64, and we got %v", len(password))
		}
	})

	t.Run("Generate_Password_SpecialChars_Letters", func(t *testing.T) {
		password, err := testService.GeneratePasswordByLength(64, specialCharAndLettersMock)

		if err != nil {
			t.Fatalf("We got an error to generate password to length -> %v", err)
		}

		findLetters, err := regexp.MatchString(regexLettersMock, password)

		if err != nil {
			t.Fatalf("We got an error to match strings in password -> %v", err)
		}

		findSpecialChars, err := regexp.MatchString(regexEspecialCharMock, password)
		findAsterisk, err := regexp.MatchString(regexAsteriskMock, password)

		if err != nil {
			t.Fatalf("We got an error to match strings in password -> %v", err)
		}

		if !findSpecialChars && !findAsterisk{
			t.Fatalf("We expected a password with Special Chars, and we got %v", password)
		}

		if !findLetters{
			t.Fatalf("We expected a password with letters, and we got %v", password)
		}

		if len(password) != 64 {
			t.Fatalf("We expected a password with length 64, and we got %v", len(password))
		}
	})

	t.Run("Generate_Password_SpecialChars_Numbers", func(t *testing.T) {
		password, err := testService.GeneratePasswordByLength(64, specialCharAndNumbersMock)

		if err != nil {
			t.Fatalf("We got an error to generate password to length -> %v", err)
		}

		findNumbers, err := regexp.MatchString(regexNumbersMock, password)

		if err != nil {
			t.Fatalf("We got an error to match strings in password -> %v", err)
		}

		findSpecialChars, err := regexp.MatchString(regexEspecialCharMock, password)
		findAsterisk, err := regexp.MatchString(regexAsteriskMock, password)

		if err != nil {
			t.Fatalf("We got an error to match strings in password -> %v", err)
		}

		if !findSpecialChars && !findAsterisk{
			t.Fatalf("We expected a password with Special Chars, and we got %v", password)
		}

		if !findNumbers{
			t.Fatalf("We expected a password with numbers, and we got %v", password)
		}

		if len(password) != 64 {
			t.Fatalf("We expected a password with length 64, and we got %v", len(password))
		}
	})

	t.Run("Generate_Password_All_Characters", func(t *testing.T) {
		password, err := testService.GeneratePasswordByLength(64, allCharactersMock)

		if err != nil {
			t.Fatalf("We got an error to generate password to length -> %v", err)
		}

		findLetters, err := regexp.MatchString(regexLettersMock, password)

		if err != nil {
			t.Fatalf("We got an error to match strings in password -> %v", err)
		}

		findNumbers, err := regexp.MatchString(regexNumbersMock, password)

		if err != nil {
			t.Fatalf("We got an error to match strings in password -> %v", err)
		}

		findSpecialChars, err := regexp.MatchString(regexEspecialCharMock, password)
		findAsterisk, err := regexp.MatchString(regexAsteriskMock, password)

		if err != nil {
			t.Fatalf("We got an error to match strings in password -> %v", err)
		}

		if !findSpecialChars && !findAsterisk{
			t.Fatalf("We expected a password with Special Chars, and we got %v", password)
		}

		if !findNumbers{
			t.Fatalf("We expected a password with numbers, and we got %v", password)
		}

		if !findLetters{
			t.Fatalf("We expected a password with letters, and we got %v", password)
		}

		if len(password) != 64 {
			t.Fatalf("We expected a password with length 64, and we got %v", len(password))
		}
	})

	t.Run("Test_Generate_Password_error", func(t *testing.T) {
		_, err := testService.GeneratePasswordByLength(0, lettersMock)

		if err == nil {
			t.Fatalf("We expected an error here")
		}
	})
}

func TestService_CheckSpecialCharAndLettersQuantity(t *testing.T) {

}

func TestService_CheckSpecialCharAndNumbersQuantity(t *testing.T) {

}

func TestService_CheckLettersAndNumbersQuantity(t *testing.T) {

}

func TestService_CheckAllCharsQuantity(t *testing.T) {

}

func TestService_CheckCharConsiderations(t *testing.T) {

}