package entity

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

type PasswordGenerator interface {
	Validate(password PasswordGen) error
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

func (p *PasswordGen) Validate(password PasswordGen) error {

	if err := validatePasswordBodyParamsTypes(password); err != nil{
		return err
	}

	return nil
}

func validatePasswordBodyParamsTypes(password PasswordGen) error {

	if reflect.TypeOf(password.Length).Kind() != reflect.Int {
		return errors.New(fmt.Sprintf("error: expected a int type and %v given, use an integer value example: 32", reflect.TypeOf(password.Length).Kind()))
	}

	if reflect.TypeOf(password.HasSpecialChar).Kind() != reflect.Bool {
		return errors.New(fmt.Sprintf("error: expected a boolean type and %v given, use 'true' or 'false'", reflect.TypeOf(password.HasSpecialChar).Kind()))
	}

	if reflect.TypeOf(password.HasNumber).Kind() != reflect.Bool {
		return errors.New(fmt.Sprintf("error: expected a boolean type and %v given, use 'true' or 'false'", reflect.TypeOf(password.HasNumber).Kind()))
	}

	if reflect.TypeOf(password.HasLetter).Kind() != reflect.Bool {
		return errors.New(fmt.Sprintf("error: expected a boolean type and %v given, use 'true' or 'false'", reflect.TypeOf(password.HasLetter).Kind()))
	}

	return nil
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
