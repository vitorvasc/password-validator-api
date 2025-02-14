package validator

import (
	"fmt"

	domain "github.com/vitorvasc/api-password-validator/internal/domain/rule"
)

// Validator handles the password validation
type Validator struct {
	rules []domain.Rule
}

// NewValidator creates a new password validator with the default rules
func NewValidator(opts ...ValidatorOption) *Validator {
	v := &Validator{}
	for _, opt := range opts {
		opt(v)
	}
	return v
}

// AddRule allows the addition of new rules after the Validator is created
func (v *Validator) AddRule(rule domain.Rule) {
	v.rules = append(v.rules, rule)
}

// Validate checks if a password meets all validation rules
func (v *Validator) Validate(password string) (bool, []string) {
	var errors []string
	for _, rule := range v.rules {
		if !rule.Validate(password) {
			errors = append(errors, rule.ErrorMessage())
		}
	}
	fmt.Println(len(errors))
	return len(errors) == 0, errors
}
