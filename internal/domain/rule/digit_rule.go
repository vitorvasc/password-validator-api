package rule

// DigitRule validates if the string contains at least one digit
type DigitRule struct{}

func NewDigitRule() Rule {
	return &DigitRule{}
}

func (r *DigitRule) Validate(t string) bool {
	for _, c := range t {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func (r *DigitRule) ErrorMessage() string {
	return "must contain at least one digit"
}
