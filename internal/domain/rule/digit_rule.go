package rule

// DigitRule validates if the password contains at least one digit
type DigitRule struct{}

func NewDigitRule() Rule {
	return &DigitRule{}
}

func (r *DigitRule) Validate(password string) bool {
	for _, char := range password {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func (r *DigitRule) ErrorMessage() string {
	return "Password must contain at least one digit."
}
