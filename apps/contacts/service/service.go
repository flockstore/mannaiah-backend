package service

import (
	"errors"

	"github.com/flockstore/mannaiah-backend/apps/contacts/domain"
)

// contactService is the default implementation of ContactService.
type contactService struct {
	repo domain.ContactRepository
}

// NewContactService creates a new instance of ContactService.
func NewContactService(repo domain.ContactRepository) domain.ContactService {
	return &contactService{repo: repo}
}

func (s *contactService) Create(c *domain.Contact) error {
	if c == nil {
		return errors.New("contact is nil")
	}
	// TODO: Add validations for domain fields if needed
	return s.repo.Save(c)
}

func (s *contactService) Get(id string) (*domain.Contact, error) {
	return s.repo.GetByID(id)
}

func (s *contactService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *contactService) List() ([]*domain.Contact, error) {
	return s.repo.List()
}
