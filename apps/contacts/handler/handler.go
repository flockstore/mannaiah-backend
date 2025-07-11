package handler

import (
	"github.com/flockstore/mannaiah-backend/apps/contacts/domain"
	"github.com/gofiber/fiber/v2"
)

// Handler manages HTTP routes for contact operations.
type Handler struct {
	service domain.ContactService
}

// New creates a new Handler with the given ContactService.
func New(service domain.ContactService) *Handler {
	return &Handler{service: service}
}

// RegisterRoutes mounts contact routes on the given router group.
func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/", h.CreateContact)
	router.Get("/", h.ListContacts)
	router.Get("/:id", h.GetContact)
	router.Delete("/:id", h.DeleteContact)
}

// CreateContact handles POST /contacts to create a new contact.
func (h *Handler) CreateContact(c *fiber.Ctx) error {
	var contact domain.Contact
	if err := c.BodyParser(&contact); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
	}

	if err := h.service.Create(&contact); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(contact)
}

// GetContact handles GET /contacts/:id to retrieve a contact by ID.
func (h *Handler) GetContact(c *fiber.Ctx) error {
	id := c.Params("id")
	contact, err := h.service.Get(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.JSON(contact)
}

// DeleteContact handles DELETE /contacts/:id to remove a contact by ID.
func (h *Handler) DeleteContact(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// ListContacts handles GET /contacts to retrieve all contacts.
func (h *Handler) ListContacts(c *fiber.Ctx) error {
	contacts, err := h.service.List()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(contacts)
}
