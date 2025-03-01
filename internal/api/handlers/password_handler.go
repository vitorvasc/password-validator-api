package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vitorvasc/api-password-validator/internal/domain/validator"
)

type PasswordHandler struct {
	validator *validator.Validator
}

func NewPasswordHandler(validator *validator.Validator) *PasswordHandler {
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

	isValid, errors := h.validator.Validate(req.Password)

	response := ValidatePasswordResponse{
		Valid:  isValid,
		Errors: errors,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
