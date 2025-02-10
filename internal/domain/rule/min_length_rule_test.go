package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MinLengthRuleTestSuite struct {
	suite.Suite
}

func TestMinLengthRuleSuite(t *testing.T) {
	suite.Run(t, new(MinLengthRuleTestSuite))
}

func (s *MinLengthRuleTestSuite) TestMinLengthRule_Validate() {
	tests := []struct {
		name      string
		minLength int
		text      string
		want      bool
	}{
		{
			name:      "should return true when text meets minimum length",
			minLength: 9,
			text:      "password123",
			want:      true,
		},
		{
			name:      "should return flase when text is shorter than minimum length",
			minLength: 9,
			text:      "pass",
			want:      false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			rule := NewMinLengthRule(tt.minLength)
			result := rule.Validate(tt.text)
			s.Equal(tt.want, result)
		})
	}
}

func (s *MinLengthRuleTestSuite) TestMinLengthRule_ErrorMessage() {
	minLength := 9
	rule := NewMinLengthRule(minLength)
	expected := fmt.Sprintf("must be at least %d characters long", minLength)

	if got := rule.ErrorMessage(); got != expected {
		s.T().Errorf("MinLengthRule.ErrorMessage() = %v, want %v", got, expected)
	}
}
