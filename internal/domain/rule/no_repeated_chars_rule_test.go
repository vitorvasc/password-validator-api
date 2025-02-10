package rule

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type NoRepeatedCharsRuleTestSuite struct {
	suite.Suite
}

func TestNoRepeatedCharsRuleSuite(t *testing.T) {
	suite.Run(t, new(NoRepeatedCharsRuleTestSuite))
}

func (s *NoRepeatedCharsRuleTestSuite) TestNoRepeatedCharsRule_Validate() {
	tests := []struct {
		name string
		text string
		want bool
	}{
		{
			name: "should return true when there are no repeated characters",
			text: "abcdef",
			want: true,
		},
		{
			name: "should return false when there are repeated characters",
			text: "aabbcc",
			want: false,
		},
	}

	rule := NewNoRepatedCharsRule()
	for _, tt := range tests {
		s.Run(tt.name, func() {
			result := rule.Validate(tt.text)
			s.Equal(tt.want, result)
		})
	}
}

func (s *NoRepeatedCharsRuleTestSuite) TestNoRepeatedCharsRule_ErrorMessage() {
	rule := NewNoRepatedCharsRule()
	expected := "must not contain repeated characters"

	if got := rule.ErrorMessage(); got != expected {
		s.T().Errorf("NoRepeatedCharsRule.ErrorMessage() = %v, want %v", got, expected)
	}
}
