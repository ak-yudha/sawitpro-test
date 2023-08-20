package handler

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
)

func ValidateRequest() (*validator.Validate, error) {
	validate := validator.New()
	err := validate.RegisterValidation("custom_prefix", validateCustomPrefix)
	if err != nil {
		return nil, err
	}
	err = validate.RegisterValidation("custom_password", complexPasswordValidator)
	if err != nil {
		return nil, err
	}

	return validate, nil
}

func ExtractJWTToken(ctx echo.Context) ([]string, error) {
	headers := ctx.Request().Header
	authHader := headers.Get("Authorization")
	t := strings.Split(authHader, " ")
	if len(t) != 2 {
		return nil, errors.New("invalid token")
	}
	return t, nil

}

func HashPassword(password string) (string, error) {
	// Generate a random salt with a cost of 12 (higher cost takes more time)

	salt, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return "", err
	}

	// Combine the password hash and salt, then return it as a string
	return string(salt), nil
}

func ComparePasswords(hashedPassword, password string) error {
	// Compare the stored hashed password with the entered password
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Custom validation function for a complex password
func complexPasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Check if the password contains at least one uppercase letter, one numeric character,
	// and one special (non-alphanumeric) character using regular expressions.
	uppercasePattern := `[A-Z]`
	numericPattern := `[0-9]`
	specialPattern := `[^A-Za-z0-9]`

	matchUppercase, _ := regexp.MatchString(uppercasePattern, password)
	matchNumeric, _ := regexp.MatchString(numericPattern, password)
	matchSpecial, _ := regexp.MatchString(specialPattern, password)

	return matchUppercase && matchNumeric && matchSpecial
}

// Custom validation function for the specified prefix
func validateCustomPrefix(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()
	prefix := fl.Param() // Get the specified prefix from the validation tag

	return phoneNumber[0:len(prefix)] == prefix
}
