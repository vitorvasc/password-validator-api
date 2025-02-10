package rule

// SpecialCharRule validates if the password contains at least one special character
type SpecialCharRule struct{}

func NewSpecialCharRule() Rule {
	return &SpecialCharRule{}
}

func (r *SpecialCharRule) Validate(password string) bool {
	specialChars := "!@#$%^&*()-+" // TODO: Convert this to use map
	for _, passwordChar := range password {
		for _, specialChar := range specialChars {
			if passwordChar == specialChar {
				return true
			}
		}
	}
	return false
}

func (r *SpecialCharRule) ErrorMessage() string {
	return "Password must contain at least one special character (!@#$%^&*()-+)."
}
