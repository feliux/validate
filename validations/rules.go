package validations

import (
	"fmt"
	"regexp"
)

var (
	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,8}$`)
	urlRegex   = regexp.MustCompile(`^(http(s)?://)?([\da-z\.-]+)\.([a-z\.]{2,6})([/\w \.-]*)*/?$`)
)

func Password() RuleSet {
	return RuleSet{
		Name: "password",
		//RuleValue: n,
		ValidateFunc: func(set RuleSet) bool {
			str, ok := set.FieldValue.(string)
			if !ok {
				return false
			}
			_, ok = ValidatePassword(str)
			return ok
		},
		MessageFunc: func(set RuleSet) string {
			return fmt.Sprintf("%s should be valid", set.FieldName)
		},
	}
}

func Message(msg string) RuleFunc {
	return func() RuleSet {
		return RuleSet{
			Name:      "message",
			RuleValue: msg,
		}
	}
}

func Required() RuleSet {
	return RuleSet{
		Name: "required",
		MessageFunc: func(set RuleSet) string {
			return fmt.Sprintf("%s is a required field", set.FieldName)
		},
		ValidateFunc: func(rule RuleSet) bool {
			str, ok := rule.FieldValue.(string)
			if !ok {
				return false
			}
			return len(str) > 0
		},
	}
}

func Url() RuleSet {
	return RuleSet{
		Name: "url",
		MessageFunc: func(set RuleSet) string {
			return "not a valid url"
		},
		ValidateFunc: func(set RuleSet) bool {
			u, ok := set.FieldValue.(string)
			if !ok {
				return false
			}
			return urlRegex.MatchString(u)
		},
	}
}

func Email() RuleSet {
	return RuleSet{
		Name: "email",
		MessageFunc: func(set RuleSet) string {
			return "email address is invalid"
		},
		ValidateFunc: func(set RuleSet) bool {
			email, ok := set.FieldValue.(string)
			if !ok {
				return false
			}
			return emailRegex.MatchString(email)
		},
	}
}

func Equal(s string) RuleFunc {
	return func() RuleSet {
		return RuleSet{
			Name:      "equal",
			RuleValue: s,
			ValidateFunc: func(set RuleSet) bool {
				str, ok := set.FieldValue.(string)
				if !ok {
					return false
				}
				return str == s
			},
			MessageFunc: func(set RuleSet) string {
				return fmt.Sprintf("%s should be equal %s", set.FieldName, s)
			},
		}
	}
}

func Max(n int) RuleFunc {
	return func() RuleSet {
		return RuleSet{
			Name:      "max",
			RuleValue: n,
			ValidateFunc: func(set RuleSet) bool {
				str, ok := set.FieldValue.(string)
				if !ok {
					return false
				}
				return len(str) <= n
			},
			MessageFunc: func(set RuleSet) string {
				return fmt.Sprintf("%s should be maximum %d characters long", set.FieldName, n)
			},
		}
	}
}

func Min(n int) RuleFunc {
	return func() RuleSet {
		return RuleSet{
			Name:      "min",
			RuleValue: n,
			ValidateFunc: func(set RuleSet) bool {
				str, ok := set.FieldValue.(string)
				if !ok {
					return false
				}
				return len(str) >= n
			},
			MessageFunc: func(set RuleSet) string {
				return fmt.Sprintf("%s should be at least %d characters long", set.FieldName, n)
			},
		}
	}
}
