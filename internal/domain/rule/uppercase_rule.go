package rule

// UppercaseRule validates if the password contains at least one uppercase letter
type UppercaseRule struct{}

func NewUppercaseRule() Rule {
	return &UppercaseRule{}
}

func (r *UppercaseRule) Validate(password string) bool {
	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func (r *UppercaseRule) ErrorMessage() string {
	return "The password must contain at least one uppercase letter."
}
