package service

import "github.com/marcos-dev88/go-password-generator/domain/entity"


type passwordMockTest struct {
	entity.PasswordGenerator
}

var (
	lettersMock               = []rune(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`)
	numbersMock               = []rune(`0123456789`)
	specialCharactersMock     = []rune(`$&*()%_+!@#=-][}{,.><;:\/?|`)
	specialCharAndNumbersMock = append(numbersMock, specialCharactersMock...)
	specialCharAndLettersMock = append(lettersMock, specialCharactersMock...)
	lettersAndNumbersMock     = append(numbersMock, lettersMock...)
	allCharactersMock         = append(specialCharAndLettersMock, numbersMock...)
)

var (
	regexNumbersMock = `[0-9]+`
	regexLettersMock = `[a-zA-Z]+`
	regexEspecialCharMock =`[^0-9a-zA-Z *]`
	regexAsteriskMock = `[*]`
)

/** Struct what implements PasswordGenerator interface **/
var passwordEntityMock = entity.PasswordGen{
	Uuid:           "",
	Password:       "",
	Length:         0,
	HasLetter:      false,
	HasNumber:      false,
	HasSpecialChar: false,
}
//var passwordMock = passwordMockTest{&passwordEntityMock}


