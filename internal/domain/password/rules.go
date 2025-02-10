package password

// Rule represents a password validation rule
type Rule interface {
	Validate(password string) bool
	ErrorMessage() string
}

// MinLengthRule validates the minimum length of a password
type MinLengthRule struct{}

func (r MinLengthRule) Validate(password string) bool {
	return len(password) >= 9
}

func (r MinLengthRule) ErrorMessage() string {
	return "Password must be at least 9 characters long."
}

// DigitRule validates if the password contains at least one digit
type DigitRule struct{}

func (r DigitRule) Validate(password string) bool {
	for _, char := range password {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func (r DigitRule) ErrorMessage() string {
	return "Password must contain at least one digit."
}

// LowercaseRule validates if the password contains at least one lowercase letter
type LowercaseRule struct{}

func (r LowercaseRule) Validate(password string) bool {
	for _, char := range password {
		if char >= 'a' && char <= 'z' {
			return true
		}
	}
	return false
}

func (r LowercaseRule) ErrorMessage() string {
	return "Password must contain at least one lowercase letter."
}

// UppercaseRule validates if the password contains at least one uppercase letter
type UppercaseRule struct{}

func (r UppercaseRule) Validate(password string) bool {
	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func (r UppercaseRule) ErrorMessage() string {
	return "The password must contain at least one uppercase letter."
}

// SpecialCharRule validates if the password contains at least one special character
type SpecialCharRule struct{}

func (r SpecialCharRule) Validate(password string) bool {
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

func (r SpecialCharRule) ErrorMessage() string {
	return "Password must contain at least one special character (!@#$%^&*()-+)."
}

// NoRepeatedCharsRule validates if the password contains no repeated characters
type NoRepeatedCharsRule struct{}

func (r NoRepeatedCharsRule) Validate(password string) bool {
	seen := make(map[rune]bool)
	for _, char := range password {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func (r NoRepeatedCharsRule) ErrorMessage() string {
	return "The password must not contain repeated characters."
}
