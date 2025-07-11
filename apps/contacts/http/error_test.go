package http

import (
	"errors"
	"testing"

	"github.com/flockstore/mannaiah-backend/apps/contacts/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestMapDomainErrorToFiber checks error testing conversion.
func TestMapDomainErrorToFiber(t *testing.T) {
	tests := []struct {
		name     string
		inputErr error
		wantCode int
		wantMsg  string
	}{
		{
			name:     "Contact not found",
			inputErr: domain.ErrContactNotFound,
			wantCode: fiber.StatusNotFound,
			wantMsg:  "contact not found",
		},
		{
			name:     "Duplicate document",
			inputErr: domain.ErrDuplicateDocument,
			wantCode: fiber.StatusConflict,
			wantMsg:  "duplicate document",
		},
		{
			name:     "Invalid name combination",
			inputErr: domain.ErrInvalidNameCombination,
			wantCode: fiber.StatusBadRequest,
			wantMsg:  "invalid name combination",
		},
		{
			name:     "Missing name",
			inputErr: domain.ErrMissingName,
			wantCode: fiber.StatusBadRequest,
			wantMsg:  "missing name",
		},
		{
			name:     "Generic internal error",
			inputErr: errors.New("something went wrong"),
			wantCode: fiber.StatusInternalServerError,
			wantMsg:  "something went wrong",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MapDomainErrorToFiber(tt.inputErr)
			var fiberErr *fiber.Error
			ok := errors.As(err, &fiberErr)
			assert.True(t, ok, "must return a fiber.Error")

			assert.Equal(t, tt.wantCode, fiberErr.Code)
			assert.Equal(t, tt.wantMsg, fiberErr.Message)
		})
	}
}
