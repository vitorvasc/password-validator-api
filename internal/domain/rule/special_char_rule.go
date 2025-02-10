package rule

// SpecialCharRule validates if the string contains at least one special character
type SpecialCharRule struct {
	charMap  map[rune]bool
	charList string
}

func NewSpecialCharRule(chars string) Rule {
	charMap := make(map[rune]bool)
	for _, char := range chars {
		charMap[char] = true
	}

	return &SpecialCharRule{
		charMap:  charMap,
		charList: chars,
	}
}

func (r *SpecialCharRule) Validate(t string) bool {
	for _, char := range t {
		if r.charMap[char] {
			return true
		}
	}
	return false
}

func (r *SpecialCharRule) ErrorMessage() string {
	return "must contain at least one special character: " + r.charList
}
