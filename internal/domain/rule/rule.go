package rule

// Rule represents a password validation rule
type Rule interface {
	Validate(password string) bool
	ErrorMessage() string
}
