package rule

// NoRepeatedCharsRule validates if the string contains no repeated characters
type NoRepeatedCharsRule struct{}

func NewNoRepatedCharsRule() Rule {
	return &NoRepeatedCharsRule{}
}

func (r *NoRepeatedCharsRule) Validate(t string) bool {
	seen := make(map[rune]bool)
	for _, char := range t {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func (r *NoRepeatedCharsRule) ErrorMessage() string {
	return "must not contain repeated characters"
}
