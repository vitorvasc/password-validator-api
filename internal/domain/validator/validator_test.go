package validator

import (
	"testing"

	"github.com/stretchr/testify/suite"
	domain "github.com/vitorvasc/api-password-validator/internal/domain/rule"
)

type ValidatorTestSuite struct {
	suite.Suite
}

func TestValidatorSuite(t *testing.T) {
	suite.Run(t, new(ValidatorTestSuite))
}

func (s *ValidatorTestSuite) TestNewValidator() {
	// Create a mock rule for testing
	mockRule := struct{ domain.Rule }{}

	validator := NewValidator(
		WithRules(mockRule),
	)

	s.NotNil(validator)
	s.Len(validator.rules, 1)
}

func (s *ValidatorTestSuite) TestAddRule() {
	validator := NewValidator()
	initialRulesCount := len(validator.rules)

	// Create a mock rule for testing
	mockRule := struct{ domain.Rule }{}

	validator.AddRule(mockRule)

	if len(validator.rules) != initialRulesCount+1 {
		s.T().Errorf("Expected %d rules after adding one, got %d", initialRulesCount+1, len(validator.rules))
	}
}
