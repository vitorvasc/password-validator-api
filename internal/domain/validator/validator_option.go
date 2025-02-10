package validator

import domain "github.com/vitorvasc/api-password-validator/internal/domain/rule"

// ValidatorOption defines a function type for configuring the Validtor
type ValidatorOption func(*Validator)

// WithRules allows setting custom rules when initializing a Validator
func WithRules(rules ...domain.Rule) ValidatorOption {
	return func(v *Validator) {
		v.rules = append(v.rules, rules...)
	}
}
