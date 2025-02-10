package rule

import "fmt"

// MinLengthRule validates the minimum length of a password
type MinLengthRule struct {
	minLength int
}

func NewMinLengthRule(minLength int) Rule {
	return &MinLengthRule{minLength}
}

func (r *MinLengthRule) Validate(password string) bool {
	return len(password) >= r.minLength
}

func (r *MinLengthRule) ErrorMessage() string {
	return fmt.Sprintf("Password must be at least %d characters long.", r.minLength)
}
