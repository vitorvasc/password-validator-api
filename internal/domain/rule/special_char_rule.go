package rule

// SpecialCharRule validates if the string contains at least one special character
type SpecialCharRule struct{}

func NewSpecialCharRule() Rule {
	return &SpecialCharRule{}
}

func (r *SpecialCharRule) Validate(t string) bool {
	specialChars := "!@#$%^&*()-+" // TODO: Convert this to use map
	for _, passwordChar := range t {
		for _, specialChar := range specialChars {
			if passwordChar == specialChar {
				return true
			}
		}
	}
	return false
}

func (r *SpecialCharRule) ErrorMessage() string {
	return "must contain at least one special character (!@#$%^&*()-+)."
}
