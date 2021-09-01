package entity

import (
	"regexp"
)

type PasswordGenerator interface {
	GetPasswordNumbers(password string) []string
	GetPasswordLetters(password string) []string
	GetPasswordSpecialChars(password string) []string
}

var (
	Letters               = []rune(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`)
	Numbers               = []rune(`0123456789`)
	SpecialCharacters     = []rune(`$&*()%_+!@#=-][}{,.><;:\/?|`)
	SpecialCharAndNumbers = append(Numbers, SpecialCharacters...)
	SpecialCharAndLetters = append(Letters, SpecialCharacters...)
	LettersAndNumbers     = append(Numbers, Letters...)
	AllCharacters         = append(SpecialCharAndLetters, Numbers...)
)

var (
	RegexNumbers = `[0-9]+`
	RegexLetters = `[a-zA-Z]+`
	RegexEspecialChar =`[^0-9a-zA-Z *]`
	RegexAsterisk = `[*]`
)

type PasswordGen struct {
	Uuid           string `json:"uuid"`
	Password       string    `json:"password"`
	Length         int       `json:"length"`
	HasLetter      bool      `json:"has_letter"`
	HasNumber      bool      `json:"has_number"`
	HasSpecialChar bool      `json:"has_special_char"`
}

func NewPasswordGen(uuid string, password string, length int, hasLetter, hasNumber, hasSpecialChar bool) *PasswordGen {
	return &PasswordGen{
		Uuid:           uuid,
		Password:       password,
		Length:         length,
		HasLetter:      hasLetter,
		HasNumber:      hasNumber,
		HasSpecialChar: hasSpecialChar,
	}
}

func (p *PasswordGen) GetPasswordNumbers(password string) []string {
	checkNum := regexp.MustCompile(RegexNumbers)
	return checkNum.FindAllString(password, -1)
}

func (p *PasswordGen) GetPasswordLetters(password string) []string {
	checkLetters := regexp.MustCompile(RegexLetters)
	return checkLetters.FindAllString(password, -1)
}

func (p *PasswordGen) GetPasswordSpecialChars(password string) []string {
	checkSpecialChar := regexp.MustCompile(RegexEspecialChar)
	addSpecialChar := regexp.MustCompile(RegexAsterisk)

	checkResult := checkSpecialChar.FindAllString(password, -1)
	checkResult2 := addSpecialChar.FindAllString(password, -1)
	checkResult = append(checkResult, checkResult2...)

	return checkResult
}
