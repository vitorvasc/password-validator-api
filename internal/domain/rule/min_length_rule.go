package rule

import "fmt"

// MinLengthRule validates the minimum length of a string
type MinLengthRule struct {
	minLength int
}

func NewMinLengthRule(minLength int) Rule {
	return &MinLengthRule{minLength}
}

func (r *MinLengthRule) Validate(t string) bool {
	return len(t) >= r.minLength
}

func (r *MinLengthRule) ErrorMessage() string {
	return fmt.Sprintf("must be at least %d characters long", r.minLength)
}
