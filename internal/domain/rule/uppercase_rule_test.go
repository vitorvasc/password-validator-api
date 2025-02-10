package rule

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UppercaseRuleTestSuite struct {
	suite.Suite
}

func TestUppercaseRuleSuite(t *testing.T) {
	suite.Run(t, new(UppercaseRuleTestSuite))
}

func (s *UppercaseRuleTestSuite) TestUppercaseRule_Validate() {
	tests := []struct {
		name string
		text string
		want bool
	}{
		{
			name: "should return true when there is an uppercase character",
			text: "Password123",
			want: true,
		},
		{
			name: "should return false when there isnt an uppercase character",
			text: "password123",
			want: false,
		},
	}

	rule := NewUppercaseRule()
	for _, tt := range tests {
		s.Run(tt.name, func() {
			result := rule.Validate(tt.text)
			s.Equal(tt.want, result)
		})
	}
}

func (s *UppercaseRuleTestSuite) TestUppercaseRule_ErrorMessage() {
	rule := NewUppercaseRule()
	expected := "must contain at least one uppercase letter"

	if got := rule.ErrorMessage(); got != expected {
		s.T().Errorf("UppercaseRule.ErrorMessage() = %v, want %v", got, expected)
	}
}
