package rule

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LowercaseRuleTestSuite struct {
	suite.Suite
}

func TestLowercaseRuleSuite(t *testing.T) {
	suite.Run(t, new(LowercaseRuleTestSuite))
}

func (s *LowercaseRuleTestSuite) TestLowercaseRule_Validate() {
	tests := []struct {
		name string
		text string
		want bool
	}{
		{
			name: "should return true when text contains a lowercase character",
			text: "a",
			want: true,
		},
		{
			name: "should return false when text contains only uppercase characters",
			text: "ABCDE",
			want: false,
		},
		{
			name: "should return false when text contains only numbers",
			text: "1",
			want: false,
		},
		{
			name: "should return false when text contains special characters only",
			text: "@#$%^&*",
			want: false,
		},
	}

	rule := NewLowercaseRule()

	for _, tt := range tests {
		s.Run(tt.name, func() {
			result := rule.Validate(tt.text)
			s.Equal(tt.want, result)
		})
	}
}

func (s *LowercaseRuleTestSuite) TestLowercaseRule_ErrorMessage() {
	rule := NewLowercaseRule()
	expected := "must contain at least one lowercase letter"

	if got := rule.ErrorMessage(); got != expected {
		s.T().Errorf("LowercaseRule.ErrorMessage() = %v, want %v", got, expected)
	}
}
