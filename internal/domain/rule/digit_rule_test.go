package rule

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DigitRuleTestSuite struct {
	suite.Suite
}

func TestDigitRuleSuite(t *testing.T) {
	suite.Run(t, new(DigitRuleTestSuite))
}

func (s *DigitRuleTestSuite) TestDigitRule_Validate() {
	tests := []struct {
		name string
		text string
		want bool
	}{
		{
			name: "should return true when text contains a digit",
			text: "Password123",
			want: true,
		},
		{
			name: "should return true when text contains only digits",
			text: "12345",
			want: true,
		},
		{
			name: "should return false when text contains no digits",
			text: "Password",
			want: false,
		},
		{
			name: "should return false when text is empty",
			text: "",
			want: false,
		},
		{
			name: "should return false when text contains only special characters",
			text: "@#$%^&*",
			want: false,
		},
	}

	rule := NewDigitRule()

	for _, tt := range tests {
		s.Run(tt.name, func() {
			result := rule.Validate(tt.text)
			s.Equal(tt.want, result)
		})
	}
}

func (s *DigitRuleTestSuite) TestDigitRule_ErrorMessage() {
	rule := NewDigitRule()
	expected := "must contain at least one digit"

	if got := rule.ErrorMessage(); got != expected {
		s.T().Errorf("DigitRule.ErrorMessage() = %v, want %v", got, expected)
	}
}
