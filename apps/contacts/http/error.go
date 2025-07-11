package http

import (
	"errors"
	"github.com/flockstore/mannaiah-backend/apps/contacts/domain"
	"github.com/gofiber/fiber/v2"
)

// MapDomainErrorToFiber translates the errors presented in the domain to a fiber error.
func MapDomainErrorToFiber(err error) error {
	switch {
	case errors.Is(err, domain.ErrContactNotFound):
		return fiber.NewError(fiber.StatusNotFound, "contact not found")
	case errors.Is(err, domain.ErrDuplicateDocument):
		return fiber.NewError(fiber.StatusConflict, "duplicate document")
	case errors.Is(err, domain.ErrInvalidNameCombination):
		return fiber.NewError(fiber.StatusBadRequest, "invalid name combination")
	case errors.Is(err, domain.ErrMissingName):
		return fiber.NewError(fiber.StatusBadRequest, "missing name")
	default:
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
}
