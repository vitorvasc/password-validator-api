package rule

// LowercaseRule validates if the password contains at least one lowercase letter
type LowercaseRule struct{}

func NewLowercaseRule() Rule {
	return &LowercaseRule{}
}

func (r *LowercaseRule) Validate(password string) bool {
	for _, char := range password {
		if char >= 'a' && char <= 'z' {
			return true
		}
	}
	return false
}

func (r *LowercaseRule) ErrorMessage() string {
	return "Password must contain at least one lowercase letter."
}
