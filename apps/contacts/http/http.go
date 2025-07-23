package http

import (
	"errors"
	"fmt"
	"github.com/flockstore/mannaiah-backend/apps/contacts/domain"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strings"
)

// Handler manages HTTP routes for contact operations.
type Handler struct {
	logger   *zap.SugaredLogger
	service  domain.ContactService
	validate *validator.Validate
}

// New creates a new Handler with the given ContactService.
func New(service domain.ContactService, l *zap.SugaredLogger) *Handler {
	return &Handler{
		logger:   l,
		service:  service,
		validate: validator.New(),
	}
}

// RegisterRoutes mounts contact routes on the given router group.
func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/", h.CreateContact)
	router.Get("/", h.ListContacts)
	router.Get("/:id", h.GetContact)
	router.Patch("/:id", h.PatchContact)
	router.Delete("/:id", h.DeleteContact)
}

// CreateContact handles POST /contacts to create a new contact.
func (h *Handler) CreateContact(c *fiber.Ctx) error {
	var input ContactInput

	if err := c.BodyParser(&input); err != nil {
		h.logger.Debug("Failed to parse body", zap.Error(err))
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON")
	}

	if err := h.validate.Struct(&input); err != nil {
		me := mapValidationErrors(err)
		h.logger.Debug("Failed to parse body", zap.Error(me))
		return me
	}

	domainContact := ToDomainContact(input)
	if err := h.service.Create(domainContact); err != nil {
		return MapDomainErrorToFiber(err)
	}

	return c.Status(fiber.StatusCreated).JSON(ToResponseDTO(domainContact))
}

// GetContact handles GET /contacts/:id to retrieve a contact by ID.
func (h *Handler) GetContact(c *fiber.Ctx) error {
	id := c.Params("id")
	contact, err := h.service.Get(id)
	if err != nil {
		return MapDomainErrorToFiber(err)
	}
	return c.JSON(ToResponseDTO(contact))
}

// DeleteContact handles DELETE /contacts/:id to remove a contact by ID.
func (h *Handler) DeleteContact(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(id); err != nil {
		return MapDomainErrorToFiber(err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// ListContacts handles GET /contacts to retrieve all contacts.
func (h *Handler) ListContacts(c *fiber.Ctx) error {
	contacts, err := h.service.List()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	response := make([]ContactResponse, len(contacts))
	for i, contact := range contacts {
		response[i] = ToResponseDTO(contact)
	}
	return c.JSON(response)
}

// PatchContact handles PATCH /contacts/:id to partially update a contact.
func (h *Handler) PatchContact(c *fiber.Ctx) error {
	id := c.Params("id")
	var patch ContactPatchInput
	if err := c.BodyParser(&patch); err != nil {
		h.logger.Debug("Failed to parse body", zap.Error(err))
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON")
	}
	if err := h.validate.Struct(&patch); err != nil {
		me := mapValidationErrors(err)
		h.logger.Debug("Failed to parse body", zap.Error(me))
		return me
	}

	domainPatch := ToDomainPatch(patch)
	updated, err := h.service.Update(id, domainPatch)
	if err != nil {
		return MapDomainErrorToFiber(err)
	}
	return c.JSON(ToResponseDTO(updated))
}

// mapValidationErrors converts validator errors into readable messages.
func mapValidationErrors(err error) error {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		var msg []string
		for _, e := range errs {
			msg = append(msg, fmt.Sprintf("field '%s' failed on '%s'", e.Field(), e.Tag()))
		}
		return fiber.NewError(fiber.StatusBadRequest, strings.Join(msg, ", "))
	}
	return fiber.NewError(fiber.StatusBadRequest, "validation failed")
}
