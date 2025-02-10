package rule

// NoRepeatedCharsRule validates if the password contains no repeated characters
type NoRepeatedCharsRule struct{}

func NewNoRepatedCharsRule() Rule {
	return &NoRepeatedCharsRule{}
}

func (r *NoRepeatedCharsRule) Validate(password string) bool {
	seen := make(map[rune]bool)
	for _, char := range password {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func (r *NoRepeatedCharsRule) ErrorMessage() string {
	return "The password must not contain repeated characters."
}
