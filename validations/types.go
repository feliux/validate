package validations

type RuleFunc func() RuleSet

type RuleSet struct {
	Name         string
	RuleValue    any
	FieldValue   any
	FieldName    any
	MessageFunc  func(RuleSet) string
	ValidateFunc func(RuleSet) bool
}

type Fields map[string][]RuleSet

type Messages map[string]string

type Validator struct {
	data   any
	fields Fields
}
