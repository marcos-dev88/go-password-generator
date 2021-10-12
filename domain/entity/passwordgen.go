package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
)

type PasswordGenerator interface {
	Validate(body []byte) error
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
	RegexNumbers      = `[0-9]+`
	RegexLetters      = `[a-zA-Z]+`
	RegexEspecialChar = `[^0-9a-zA-Z *]`
	RegexAsterisk     = `[*]`
)

type PasswordGen struct {
	Uuid           string `json:"uuid"`
	Password       string `json:"password"`
	Length         int    `json:"length"`
	HasLetter      bool   `json:"has_letter"`
	HasNumber      bool   `json:"has_number"`
	HasSpecialChar bool   `json:"has_special_char"`
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

func (p *PasswordGen) Validate(body []byte) error {

	expectedBody := map[string]interface{}{
		"length":           0,
		"has_letter":       false,
		"has_number":       false,
		"has_special_char": false,
	}


	if err := json.Unmarshal(body, &expectedBody); err != nil {
		return err
	}

	log.Printf("value -> %v", expectedBody)

	if expectedBody["length"] == 0 {
		return errors.New("error: expected a json key 'length' with a value: 8, 16, 32 or 64")
	}

	if reflect.TypeOf(expectedBody["length"]).Kind() != reflect.Float64 {
		return errors.New(fmt.Sprintf("error: expected a float64 type and %v given, use an integer value example: 32", reflect.TypeOf(expectedBody["length"]).Kind()))
	}

	if reflect.TypeOf(expectedBody["has_special_char"]).Kind() != reflect.Bool {
		return errors.New(fmt.Sprintf("error: expected a boolean type and %v given, use 'true' or 'false'", reflect.TypeOf(expectedBody["has_special_char"]).Kind()))
	}

	if reflect.TypeOf(expectedBody["has_number"]).Kind() != reflect.Bool {
		return errors.New(fmt.Sprintf("error: expected a boolean type and %v given, use 'true' or 'false'", reflect.TypeOf(expectedBody["has_number"]).Kind()))
	}

	if reflect.TypeOf(expectedBody["has_letter"]).Kind() != reflect.Bool {
		return errors.New(fmt.Sprintf("error: expected a boolean type and %v given, use 'true' or 'false'", reflect.TypeOf(expectedBody["has_letter"]).Kind()))
	}


	var conditionsSlice []bool
	lengthValue := expectedBody["length"].(float64)
	hasNumberValue := expectedBody["has_number"].(bool)
	hasLetterValue := expectedBody["has_letter"].(bool)
	hasSpecialCharValue := expectedBody["has_special_char"].(bool)
	conditionsSlice = append(conditionsSlice, hasNumberValue, hasLetterValue, hasSpecialCharValue)

	if err := validateLength(lengthValue); err != nil {
		return err
	}

	if err := validateConditions(conditionsSlice); err != nil {
		return err
	}

	return nil
}

func validateLength (lengthValue float64) error {
	if (lengthValue == 8) ||
		(lengthValue == 16) ||
		(lengthValue == 32) ||
		(lengthValue == 64) {
		return nil
	}else{
		return errors.New(fmt.Sprintf("error: length field must receive one of this numbers: 8, 16, 32, 64 and %v given", lengthValue))
	}
}

func validateConditions(conditions []bool) error {
	if !conditions[0] && !conditions[1] && !conditions[2] {
		return errors.New("error: not condition given, you must define if your password has just numbers, letters or special char. example: 'has_number: true")
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
