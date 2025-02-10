package rule

// UppercaseRule validates if the string contains at least one uppercase letter
type UppercaseRule struct{}

func NewUppercaseRule() Rule {
	return &UppercaseRule{}
}

func (r *UppercaseRule) Validate(t string) bool {
	for _, char := range t {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func (r *UppercaseRule) ErrorMessage() string {
	return "must contain at least one uppercase letter"
}
