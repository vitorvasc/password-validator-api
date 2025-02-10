package password

// Validator handles the password validation
type Validator struct {
	rules []Rule
}

// NewValidator creates a new password validator with the default rules
func NewValidator() *Validator {
	return &Validator{
		rules: []Rule{
			MinLengthRule{},
			DigitRule{},
			LowercaseRule{},
			UppercaseRule{},
			SpecialCharRule{},
			NoRepeatedCharsRule{},
		},
	}
}

// AddRule allows the addition of new validation rules
func (v *Validator) AddRule(rule Rule) {
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
