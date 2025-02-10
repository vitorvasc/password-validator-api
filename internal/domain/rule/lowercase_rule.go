package rule

// LowercaseRule validates if the string contains at least one lowercase letter
type LowercaseRule struct{}

func NewLowercaseRule() Rule {
	return &LowercaseRule{}
}

func (r *LowercaseRule) Validate(t string) bool {
	for _, char := range t {
		if char >= 'a' && char <= 'z' {
			return true
		}
	}
	return false
}

func (r *LowercaseRule) ErrorMessage() string {
	return "must contain at least one lowercase letter"
}
