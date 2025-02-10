package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vitorvasc/api-password-validator/internal/domain/password"
)

type PasswordHandler struct {
	validator *password.Validator
}

func NewPasswordHandler(validator *password.Validator) *PasswordHandler {
	return &PasswordHandler{validator}
}

type ValidatePasswordRequest struct {
	Password string `json:"password"`
}

type ValidatePasswordResponse struct {
	Valid  bool     `json:"valid"`
	Errors []string `json:"errors,omitempty"`
}

func (h *PasswordHandler) ValidatePassword(w http.ResponseWriter, r *http.Request) {
	var req ValidatePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	isValid := h.validator.Validate(req.Password)
	errors := []string{}
	if !isValid {
		errors = h.validator.GetValidationErrors(req.Password)
	}

	response := ValidatePasswordResponse{
		Valid:  isValid,
		Errors: errors,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
