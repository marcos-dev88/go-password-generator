package service

import (
	"reflect"
	"regexp"
	"testing"
)

func TestService_GeneratePasswordByLength(t *testing.T) {
	testService := NewService(passwordEntityMock)

	t.Run("Generate_Password_Letters_Only", func(t *testing.T) {
		password, err := testService.GeneratePasswordByLength(64, lettersMock)

		if err != nil {
			t.Fatalf("We got an error to generate password to length -> %v", err)
		}

		findLetters, err := regexp.MatchString(regexLettersMock, password)

		if err != nil {
			t.Fatalf("We got an error to match strings in password -> %v", err)
		}

		if !findLetters {
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

		if !findNumbers {
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

		if !findSpecialChars {
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

		if !findLetters {
			t.Fatalf("We expected a password with letters, and we got %v", password)
		}

		if !findNumbers {
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

		if !findSpecialChars && !findAsterisk {
			t.Fatalf("We expected a password with Special Chars, and we got %v", password)
		}

		if !findLetters {
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

		if !findSpecialChars && !findAsterisk {
			t.Fatalf("We expected a password with Special Chars, and we got %v", password)
		}

		if !findNumbers {
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

		if !findSpecialChars && !findAsterisk {
			t.Fatalf("We expected a password with Special Chars, and we got %v", password)
		}

		if !findNumbers {
			t.Fatalf("We expected a password with numbers, and we got %v", password)
		}

		if !findLetters {
			t.Fatalf("We expected a password with letters, and we got %v", password)
		}

		if len(password) != 64 {
			t.Fatalf("We expected a password with length 64, and we got %v", len(password))
		}
	})
}

func TestService_CheckError(t *testing.T) {
	testService := NewService(passwordEntityMock)

	t.Run("Test_Generate_Password_error", func(t *testing.T) {
		_, err := testService.GeneratePasswordByLength(0, lettersMock)

		if err == nil {
			t.Fatalf("We expected an error here")
		}
	})
}

func TestService_CheckSpecialCharAndLettersQuantity(t *testing.T) {
	testService := NewService(passwordEntityMock)

	t.Run("Test_CheckQuantity_SpecialChar_and_Letters_32", func(t *testing.T) {

		if len(passwordLettersAndSpecialCharMock32.Password) != 32 {
			t.Fatalf("We expected an password with length 32 and we got %x", len(passwordLettersAndSpecialCharMock32.Password))
		}

		if !testService.CheckSpecialCharAndLettersQuantity(passwordLettersAndSpecialCharMock32) {
			t.Fatalf("We expected a true value and got false")
		}
	})

	t.Run("Test_CheckQuantity_SpecialChar_and_Letters_32 - Fail", func(t *testing.T) {
		if len(passwordNumbersOnlyMock32.Password) != 32 {
			t.Fatalf("We expected an password with length 32 and we got %x", len(passwordNumbersOnlyMock32.Password))
		}

		if testService.CheckSpecialCharAndLettersQuantity(passwordNumbersOnlyMock32) {
			t.Fatalf("We expected a false value and got true")
		}
	})

	t.Run("Test_CheckQuantity_SpecialChar_and_Letters_16", func(t *testing.T) {

		if len(passwordLettersAndSpecialCharMock16.Password) != 16 {
			t.Fatalf("We expected an password with length 16 and we got %x", len(passwordLettersAndSpecialCharMock16.Password))
		}

		if !testService.CheckSpecialCharAndLettersQuantity(passwordLettersAndSpecialCharMock16) {
			t.Fatalf("We expected a true value and got false")
		}
	})

	t.Run("Test_CheckQuantity_SpecialChar_and_Letters_16 - Fail", func(t *testing.T) {
		if len(passwordNumbersOnlyMock16.Password) != 16 {
			t.Fatalf("We expected an password with length 16 and we got %x", len(passwordNumbersOnlyMock16.Password))
		}

		if testService.CheckSpecialCharAndLettersQuantity(passwordNumbersOnlyMock16) {
			t.Fatalf("We expected a false value and got true")
		}
	})

	t.Run("Test_CheckQuantity_SpecialChar_and_Letters_8", func(t *testing.T) {

		if len(passwordLettersAndSpecialCharMock8.Password) != 8 {
			t.Fatalf("We expected an password with length 8 and we got %x", len(passwordLettersAndSpecialCharMock8.Password))
		}

		if !testService.CheckSpecialCharAndLettersQuantity(passwordLettersAndSpecialCharMock8) {
			t.Fatalf("We expected a true value and got false")
		}
	})

	t.Run("Test_CheckQuantity_SpecialChar_and_Letters_8 - Fail", func(t *testing.T) {
		if len(passwordNumbersOnlyMock8.Password) != 8 {
			t.Fatalf("We expected an password with length 8 and we got %x", len(passwordNumbersOnlyMock8.Password))
		}

		if testService.CheckSpecialCharAndLettersQuantity(passwordNumbersOnlyMock8) {
			t.Fatalf("We expected a false value and got true")
		}
	})
}

func TestService_CheckSpecialCharAndNumbersQuantity(t *testing.T) {
	testService := NewService(passwordEntityMock)

	t.Run("Test_CheckQuantity_SpecialChar_and_Numbers_32", func(t *testing.T) {

		if len(passwordNumbersAndSpecialCharMock32.Password) != 32 {
			t.Fatalf("We expected an password with length 32 and we got %x", len(passwordNumbersAndSpecialCharMock32.Password))
		}

		if !testService.CheckSpecialCharAndNumbersQuantity(passwordNumbersAndSpecialCharMock32) {
			t.Fatalf("We expected a true value and got false")
		}
	})

	t.Run("Test_CheckQuantity_SpecialChar_and_Numbers_32 - Fail", func(t *testing.T) {
		if len(passwordNumbersOnlyMock32.Password) != 32 {
			t.Fatalf("We expected an password with length 32 and we got %x", len(passwordNumbersOnlyMock32.Password))
		}

		if testService.CheckSpecialCharAndNumbersQuantity(passwordNumbersOnlyMock32) {
			t.Fatalf("We expected a false value and got true")
		}
	})

	t.Run("Test_CheckQuantity_SpecialChar_and_Numbers_16", func(t *testing.T) {

		if len(passwordNumbersAndSpecialCharMock16.Password) != 16 {
			t.Fatalf("We expected an password with length 16 and we got %x", len(passwordNumbersAndSpecialCharMock16.Password))
		}

		if !testService.CheckSpecialCharAndNumbersQuantity(passwordNumbersAndSpecialCharMock16) {
			t.Fatalf("We expected a true value and got false")
		}
	})

	t.Run("Test_CheckQuantity_SpecialChar_and_Numbers_16 - Fail", func(t *testing.T) {
		if len(passwordNumbersOnlyMock16.Password) != 16 {
			t.Fatalf("We expected an password with length 16 and we got %x", len(passwordNumbersOnlyMock16.Password))
		}

		if testService.CheckSpecialCharAndNumbersQuantity(passwordNumbersOnlyMock16) {
			t.Fatalf("We expected a false value and got true")
		}
	})

	t.Run("Test_CheckQuantity_SpecialChar_and_Numbers_8", func(t *testing.T) {

		if len(passwordNumbersAndSpecialCharMock8.Password) != 8 {
			t.Fatalf("We expected an password with length 8 and we got %x", len(passwordNumbersAndSpecialCharMock8.Password))
		}

		if !testService.CheckSpecialCharAndNumbersQuantity(passwordNumbersAndSpecialCharMock8) {
			t.Fatalf("We expected a true value and got false")
		}
	})

	t.Run("Test_CheckQuantity_SpecialChar_and_Numbers_8 - Fail", func(t *testing.T) {
		if len(passwordNumbersOnlyMock8.Password) != 8 {
			t.Fatalf("We expected an password with length 8 and we got %x", len(passwordNumbersOnlyMock8.Password))
		}

		if testService.CheckSpecialCharAndNumbersQuantity(passwordNumbersOnlyMock8) {
			t.Fatalf("We expected a false value and got true")
		}
	})
}

func TestService_CheckLettersAndNumbersQuantity(t *testing.T) {
	testService := NewService(passwordEntityMock)

	t.Run("Test_CheckQuantity_Letters_and_Numbers_32", func(t *testing.T) {

		if len(passwordLettersAndNumbersMock32.Password) != 32 {
			t.Fatalf("We expected an password with length 32 and we got %x", len(passwordLettersAndNumbersMock32.Password))
		}

		if !testService.CheckLettersAndNumbersQuantity(passwordLettersAndNumbersMock32) {
			t.Fatalf("We expected a true value and got false")
		}
	})

	t.Run("Test_CheckQuantity_Letters_and_Numbers_32 - Fail", func(t *testing.T) {
		if len(passwordNumbersOnlyMock32.Password) != 32 {
			t.Fatalf("We expected an password with length 32 and we got %x", len(passwordNumbersOnlyMock32.Password))
		}

		if testService.CheckLettersAndNumbersQuantity(passwordNumbersOnlyMock32) {
			t.Fatalf("We expected a false value and got true")
		}
	})

	t.Run("Test_CheckQuantity_Letters_and_Numbers_16", func(t *testing.T) {

		if len(passwordLettersAndNumbersMock16.Password) != 16 {
			t.Fatalf("We expected an password with length 16 and we got %x", len(passwordLettersAndNumbersMock16.Password))
		}

		if !testService.CheckLettersAndNumbersQuantity(passwordLettersAndNumbersMock16) {
			t.Fatalf("We expected a true value and got false")
		}
	})

	t.Run("Test_CheckQuantity_Letters_and_Numbers_16 - Fail", func(t *testing.T) {
		if len(passwordNumbersOnlyMock16.Password) != 16 {
			t.Fatalf("We expected an password with length 16 and we got %x", len(passwordNumbersOnlyMock16.Password))
		}

		if testService.CheckLettersAndNumbersQuantity(passwordNumbersOnlyMock16) {
			t.Fatalf("We expected a false value and got true")
		}
	})

	t.Run("Test_CheckQuantity_Letters_and_Numbers_8", func(t *testing.T) {

		if len(passwordLettersAndNumbersMock8.Password) != 8 {
			t.Fatalf("We expected an password with length 8 and we got %x", len(passwordLettersAndNumbersMock8.Password))
		}

		if !testService.CheckLettersAndNumbersQuantity(passwordLettersAndNumbersMock8) {
			t.Fatalf("We expected a true value and got false")
		}
	})

	t.Run("Test_CheckQuantity_Letters_and_Numbers_8 - Fail", func(t *testing.T) {
		if len(passwordNumbersOnlyMock8.Password) != 8 {
			t.Fatalf("We expected an password with length 16 and we got %x", len(passwordNumbersOnlyMock8.Password))
		}

		if testService.CheckLettersAndNumbersQuantity(passwordNumbersOnlyMock8) {
			t.Fatalf("We expected a false value and got true")
		}
	})
}

func TestService_CheckAllCharsQuantity(t *testing.T) {
	testService := NewService(passwordEntityMock)

	t.Run("Test_CheckQuantity_All_Characters_32", func(t *testing.T) {

		if len(passwordAllMock32.Password) != 32 {
			t.Fatalf("We expected an password with length 32 and we got %x", len(passwordAllMock32.Password))
		}

		if !testService.CheckAllCharsQuantity(passwordAllMock32) {
			t.Fatalf("We expected a true value and got false")
		}
	})

	t.Run("Test_CheckQuantity_All_Characters_32 - Fail", func(t *testing.T) {
		if len(passwordNumbersOnlyMock32.Password) != 32 {
			t.Fatalf("We expected an password with length 32 and we got %x", len(passwordNumbersOnlyMock32.Password))
		}

		if testService.CheckAllCharsQuantity(passwordNumbersOnlyMock32) {
			t.Fatalf("We expected a false value and got true")
		}
	})

	t.Run("Test_CheckQuantity_All_Characters_16", func(t *testing.T) {

		if len(passwordAllMock16.Password) != 16 {
			t.Fatalf("We expected an password with length 16 and we got %x", len(passwordAllMock16.Password))
		}

		if !testService.CheckAllCharsQuantity(passwordAllMock16) {
			t.Fatalf("We expected a true value and got false")
		}
	})

	t.Run("Test_CheckQuantity_All_Characters_16 - Fail", func(t *testing.T) {
		if len(passwordNumbersOnlyMock16.Password) != 16 {
			t.Fatalf("We expected an password with length 16 and we got %x", len(passwordNumbersOnlyMock16.Password))
		}

		if testService.CheckAllCharsQuantity(passwordNumbersOnlyMock16) {
			t.Fatalf("We expected a false value and got true")
		}
	})

	t.Run("Test_CheckQuantity_All_Characters_8", func(t *testing.T) {

		if len(passwordAllMock8.Password) != 8 {
			t.Fatalf("We expected an password with length 8 and we got %x", len(passwordAllMock8.Password))
		}

		if !testService.CheckAllCharsQuantity(passwordAllMock8) {
			t.Fatalf("We expected a true value and got false")
		}
	})

	t.Run("Test_CheckQuantity_All_Characters_8 - Fail", func(t *testing.T) {
		if len(passwordNumbersOnlyMock8.Password) != 8 {
			t.Fatalf("We expected an password with length 8 and we got %x", len(passwordNumbersOnlyMock8.Password))
		}

		if testService.CheckAllCharsQuantity(passwordNumbersOnlyMock8) {
			t.Fatalf("We expected a false value and got true")
		}
	})
}

func TestService_CheckCharConsiderations(t *testing.T) {
	testService := NewService(passwordEntityMock)

	t.Run("Test_All_Characters_considerations", func(t *testing.T) {
		passwordConsiderations := testService.CheckCharConsiderations(*passwordAllMock8)

		if !reflect.DeepEqual(passwordConsiderations, allCharactersMock) {
			t.Fatalf("We expected %v and we got %v", allCharactersMock, passwordConsiderations)
		}
	})

	t.Run("Test_SpecialChar_and_Letters_considerations", func(t *testing.T) {
		passwordConsiderations := testService.CheckCharConsiderations(*passwordLettersAndSpecialCharMock8)

		if !reflect.DeepEqual(passwordConsiderations, specialCharAndLettersMock) {
			t.Fatalf("We expected %v and we got %v", specialCharAndLettersMock, passwordConsiderations)
		}
	})

	t.Run("Test_SpecialChar_and_Numbers_considerations", func(t *testing.T) {
		passwordConsiderations := testService.CheckCharConsiderations(*passwordNumbersAndSpecialCharMock8)

		if !reflect.DeepEqual(passwordConsiderations, specialCharAndNumbersMock) {
			t.Fatalf("We expected %v and we got %v", specialCharAndNumbersMock, passwordConsiderations)
		}
	})

	t.Run("Test_Letters_and_Numbers_considerations", func(t *testing.T) {
		passwordConsiderations := testService.CheckCharConsiderations(*passwordLettersAndNumbersMock8)

		if !reflect.DeepEqual(passwordConsiderations, lettersAndNumbersMock) {
			t.Fatalf("We expected %v and we got %v", lettersAndNumbersMock, passwordConsiderations)
		}
	})

	t.Run("Test_Letters_considerations", func(t *testing.T) {
		passwordConsiderations := testService.CheckCharConsiderations(*passwordLettersOnlyMock8)

		if !reflect.DeepEqual(passwordConsiderations, lettersMock) {
			t.Fatalf("We expected %v and we got %v", lettersMock, passwordConsiderations)
		}
	})

	t.Run("Test_Numbers_considerations", func(t *testing.T) {
		passwordConsiderations := testService.CheckCharConsiderations(*passwordNumbersOnlyMock8)

		if !reflect.DeepEqual(passwordConsiderations, numbersMock) {
			t.Fatalf("We expected %v and we got %v", numbersMock, passwordConsiderations)
		}
	})

	t.Run("Test_Special_Characters_considerations", func(t *testing.T) {
		passwordConsiderations := testService.CheckCharConsiderations(*passwordSpecialCharsOnlyMock8)

		if !reflect.DeepEqual(passwordConsiderations, specialCharactersMock) {
			t.Fatalf("We expected %v and we got %v", specialCharactersMock, passwordConsiderations)
		}
	})

	t.Run("Test_return_null_considerations", func(t *testing.T) {
		passwordConsiderations := testService.CheckCharConsiderations(*passwordWithoutConsiderationsMock)

		if passwordConsiderations != nil {
			t.Fatalf("We expected %v and we got %v", nil, passwordConsiderations)
		}
	})
}
