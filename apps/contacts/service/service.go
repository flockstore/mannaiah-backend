package service

import (
	"errors"
	"time"

	"github.com/flockstore/mannaiah-backend/apps/contacts/domain"
	"github.com/google/uuid"
)

// contactService provides the business logic for managing contacts.
type contactService struct {
	repo domain.ContactRepository
}

// NewContactService creates a new instance of ContactService.
func NewContactService(repo domain.ContactRepository) domain.ContactService {
	return &contactService{repo: repo}
}

// Create creates a new contact, generating the ID and timestamps.
func (s *contactService) Create(c *domain.Contact) error {
	c.ID = uuid.NewString()
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
	return s.repo.Save(c)
}

// Get retrieves a contact by its ID.
func (s *contactService) Get(id string) (*domain.Contact, error) {
	return s.repo.GetByID(id)
}

// Delete removes a contact by its ID.
func (s *contactService) Delete(id string) error {
	return s.repo.Delete(id)
}

// List retrieves all contacts.
func (s *contactService) List() ([]*domain.Contact, error) {
	return s.repo.List()
}

// Update applies a patch to a contact and updates its timestamp.
func (s *contactService) Update(id string, patch *domain.ContactPatch) error {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("contact not found")
	}

	domain.ApplyPatch(existing, patch)
	existing.UpdatedAt = time.Now()

	return s.repo.Save(existing)
}
