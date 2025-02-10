package validator

import domain "github.com/vitorvasc/api-password-validator/internal/domain/rule"

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
func (v *Validator) Validate(password string) bool {
	for _, rule := range v.rules {
		if !rule.Validate(password) {
			return false
		}
	}
	return true
}

// GetValidationErrors returns all validation error messages
func (v *Validator) GetValidationErrors(password string) []string {
	var errors []string // TODO: Get error message on Validate() method, avoid going through the loop twice
	for _, rule := range v.rules {
		if !rule.Validate(password) {
			errors = append(errors, rule.ErrorMessage())
		}
	}
	return errors
}
