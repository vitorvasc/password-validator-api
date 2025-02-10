package rule

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SpecialCharRuleTestSuite struct {
	suite.Suite
}

func TestSpecialCharRuleSuite(t *testing.T) {
	suite.Run(t, new(SpecialCharRuleTestSuite))
}

func (s *SpecialCharRuleTestSuite) TestSpecialCharRule_Validate() {
	tests := []struct {
		name string
		text string
		want bool
	}{
		{
			name: "should return true when there is at least one special character",
			text: "password@123",
			want: true,
		},
		{
			name: "should return false when there are no special characters",
			text: "password123",
			want: false,
		},
	}

	rule := NewSpecialCharRule()
	for _, tt := range tests {
		s.Run(tt.name, func() {
			result := rule.Validate(tt.text)
			s.Equal(tt.want, result)
		})
	}
}

func (s *SpecialCharRuleTestSuite) TestSpecialCharRule_ErrorMessage() {
	rule := NewSpecialCharRule()
	expected := "must not contain repeated characters"

	if got := rule.ErrorMessage(); got != expected {
		s.T().Errorf("SpecialCharRule.ErrorMessage() = %v, want %v", got, expected)
	}
}
