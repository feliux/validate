package validations

import (
	"strings"
	"unicode"
)

func Validate(in any, out any, fields Fields) bool {
	return true
}

func Rules(rules ...RuleFunc) []RuleSet {
	ruleSets := make([]RuleSet, len(rules))
	for i := 0; i < len(ruleSets); i++ {
		ruleSets[i] = rules[i]()
	}
	return ruleSets
}

func New(data any, fields Fields) *Validator {
	return &Validator{
		fields: fields,
		data:   data,
	}
}

func (v *Validator) Validate(target any) bool {
	ok := true
	for fieldName, ruleSets := range v.fields {
		// reflect panics on un-exported variables.
		if !unicode.IsUpper(rune(fieldName[0])) {
			continue
		}
		fieldValue := getFieldValueByName(v.data, fieldName)
		for _, set := range ruleSets {
			set.FieldValue = fieldValue
			set.FieldName = fieldName
			if set.Name == "message" {
				setErrorMessage(target, fieldName, set.RuleValue.(string))
				continue
			}
			if !set.ValidateFunc(set) {
				msg := set.MessageFunc(set)
				setErrorMessage(target, fieldName, msg)
				ok = false
			}
		}
	}
	return ok
}

// ValidatePassword checks if the password is strong and meets the criteria:
// - At least 8 characters long
// - Contains at least one digit
// - Contains at least one lowercase letter
// - Contains at least one uppercase letter
// - Contains at least one special character
func ValidatePassword(password string) (string, bool) {
	var (
		hasUpper     = false
		hasLower     = false
		hasNumber    = false
		hasSpecial   = false
		specialRunes = "!@#$%^&*"
	)

	if len(password) < 8 {
		return "Password must contain at least 8 characters", false
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char) || strings.ContainsRune(specialRunes, char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return "Password must contain at least 1 uppercase character", false
	}
	if !hasLower {
		return "Password must contain at least 1 lowercase character", false
	}
	if !hasNumber {
		return "Password must contain at least 1 numeric character (0, 1, 2, ...)", false
	}
	if !hasSpecial {
		return "Password must contain at least 1 special character (@, ;, _, ...)", false
	}
	return "", true
}
