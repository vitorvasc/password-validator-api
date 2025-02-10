package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/vitorvasc/api-password-validator/internal/domain/rule"
	"github.com/vitorvasc/api-password-validator/internal/domain/validator"
)

type PasswordHandlerTestSuite struct {
	suite.Suite
	validator *validator.Validator
	handler   *PasswordHandler
}

func TestPasswordHandlerSuite(t *testing.T) {
	suite.Run(t, new(PasswordHandlerTestSuite))
}

func (s *PasswordHandlerTestSuite) SetupTest() {
	s.validator = validator.NewValidator(
		validator.WithRules(
			rule.NewMinLengthRule(9),
			rule.NewDigitRule(),
			rule.NewLowercaseRule(),
			rule.NewUppercaseRule(),
			rule.NewSpecialCharRule("!@#$%^&*()-+"),
			rule.NewNoRepatedCharsRule(),
		),
	)
	s.handler = NewPasswordHandler(s.validator)
}

func (s *PasswordHandlerTestSuite) TestValidatePassword() {
	tests := []struct {
		name          string
		password      string
		expectedValid bool
		expectedCode  int
	}{
		{
			name:          "given a valid password, should return true. (>= 9 characters, at least 1 digit, at least 1 lowercase character, at least 1 uppercase character, at least 1 special character, no repeateed character)",
			password:      "AbTp9!fok",
			expectedValid: true,
			expectedCode:  http.StatusOK,
		},
		{
			name:          "given an invalid password, should return false. (< 9 characters)",
			password:      "AbTp9!fo",
			expectedValid: false,
			expectedCode:  http.StatusOK,
		},
		{
			name:          "given an invalid password, should return false. (no digits)",
			password:      "AbTpZ!fok",
			expectedValid: false,
			expectedCode:  http.StatusOK,
		},
		{
			name:          "given an invalid password, should return false. (no uppercase character)",
			password:      "abtp9!fok",
			expectedValid: false,
			expectedCode:  http.StatusOK,
		},
		{
			name:          "given an invalid password, should return false. (no lowercase character)",
			password:      "ABTP9!FOK",
			expectedValid: false,
			expectedCode:  http.StatusOK,
		},
		{
			name:          "given an invalid password, should return flase. (no special character)",
			password:      "AbTp9Zfok",
			expectedValid: false,
			expectedCode:  http.StatusOK,
		},
		{
			name:          "given an invalid password, should return false. (space character)",
			password:      "AbTp9 fok",
			expectedValid: false,
			expectedCode:  http.StatusOK,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			reqBody := ValidatePasswordRequest{
				Password: tt.password,
			}
			bodyBytes, err := json.Marshal(reqBody)
			s.NoError(err)

			req := httptest.NewRequest(http.MethodPost, "/v1/users/validate-password", bytes.NewReader(bodyBytes))
			w := httptest.NewRecorder()

			s.handler.ValidatePassword(w, req)

			s.Equal(tt.expectedCode, w.Code)

			var response ValidatePasswordResponse
			err = json.NewDecoder(w.Body).Decode(&response)
			s.NoError(err)

			s.Equal(tt.expectedValid, response.Valid)

			if tt.expectedValid {
				s.Zero(response.Errors)
			} else {
				s.NotZero(response.Errors)
			}
		})
	}
}

func (s *PasswordHandlerTestSuite) TestValidatePassword_InvalidJSON_ShouldReturnBadRequest() {
	req := httptest.NewRequest(http.MethodPost, "/v1/users/validate-password", bytes.NewReader([]byte("invalid json")))
	w := httptest.NewRecorder()

	s.handler.ValidatePassword(w, req)

	s.Equal(http.StatusBadRequest, w.Code)
}
