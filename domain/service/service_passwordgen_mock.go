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
	regexNumbersMock      = `[0-9]+`
	regexLettersMock      = `[a-zA-Z]+`
	regexEspecialCharMock = `[^0-9a-zA-Z *]`
	regexAsteriskMock     = `[*]`
)

/** Struct what implements PasswordGenerator interface **/
var passwordEntityMock = entity.NewPasswordGen("", "", 0, false, false, false)
var passwordMock = passwordMockTest{passwordEntityMock}

// Mocking passwords with only letters, numbers or special chars
var (

	// Letters
	passwordLettersOnlyMock8 = entity.NewPasswordGen(
		"",
		`LKTDLBQf`,
		8,
		true,
		false,
		false,
	)

	// Numbers
	passwordNumbersOnlyMock32 = entity.NewPasswordGen(
		"",
		`58659395404719625295901915395462`,
		32,
		false,
		true,
		false,
	)


	passwordNumbersOnlyMock16 = entity.NewPasswordGen(
		"",
		`9727203247530244`,
		16,
		false,
		true,
		false,
	)

	passwordNumbersOnlyMock8 = entity.NewPasswordGen(
		"",
		`95033633`,
		8,
		false,
		true,
		false,
	)

	// Special Char
	passwordSpecialCharsOnlyMock8 = entity.NewPasswordGen(
		"",
		`,?{#}%=@`,
		8,
		false,
		false,
		true,
	)
)

// Mocking passwords with 32 characters
var (
	passwordLettersAndSpecialCharMock32 = entity.NewPasswordGen(
		"",
		`]+G:AR_fk=]qMeM.xIcMxlMH)=jlDGiJ`,
		32,
		true,
		false,
		true,
	)
	passwordNumbersAndSpecialCharMock32 = entity.NewPasswordGen(
		"",
		`5-{,$2021_8|[7|$.}6<48@>43921@[-`,
		32,
		false,
		true,
		true,
	)

	passwordLettersAndNumbersMock32 = entity.NewPasswordGen(
		"",
		`XadV00x10w1WcgZstVbnff2R8ID1lBTM`,
		32,
		true,
		true,
		false,
	)

	passwordAllMock32 = entity.NewPasswordGen(
		"",
		`?dqJ7q-CkyEJ%J|9a1VTNyRq8*3bC%p5`,
		32,
		true,
		true,
		true,
	)
)

// Mocking passwords with 16 characters
var (
	passwordLettersAndSpecialCharMock16 = entity.NewPasswordGen(
		"",
		`B:+b+Yqn?UUo$%dt`,
		16,
		true,
		false,
		true,
	)
	passwordNumbersAndSpecialCharMock16 = entity.NewPasswordGen(
		"",
		`8+:.2]15{+4?8[=2`,
		16,
		false,
		true,
		true,
	)

	passwordLettersAndNumbersMock16 = entity.NewPasswordGen(
		"",
		`9E6QPt0W80FMDC1g`,
		16,
		true,
		true,
		false,
	)

	passwordAllMock16 = entity.NewPasswordGen(
		"",
		`aQ0G+eJ%7O5GiK#,`,
		16,
		true,
		true,
		true,
	)
)

// Mocking passwords with 8 characters
var (
	passwordLettersAndSpecialCharMock8 = entity.NewPasswordGen(
		"",
		`(G@B&PzZ`,
		8,
		true,
		false,
		true,
	)
	passwordNumbersAndSpecialCharMock8 = entity.NewPasswordGen(
		"",
		`*5#%0-4?`,
		8,
		false,
		true,
		true,
	)

	passwordLettersAndNumbersMock8 = entity.NewPasswordGen(
		"",
		`z5LV6P7z`,
		8,
		true,
		true,
		false,
	)

	passwordAllMock8 = entity.NewPasswordGen(
		"",
		`)i1@/b97`,
		8,
		true,
		true,
		true,
	)
)

// Mock consideration all false
var passwordWithoutConsiderationsMock = entity.NewPasswordGen(
	"",
	`)i1@/b97`,
	8,
	false,
	false,
	false,
)
