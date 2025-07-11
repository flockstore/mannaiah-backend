package http

import (
	"github.com/flockstore/mannaiah-backend/apps/contacts/domain"
	"time"
)

// ToDomainContact converts a ContactInput DTO into a domain.Contact entity.
//
// This function is used when creating a new contact from external input.
func ToDomainContact(input ContactInput) *domain.Contact {
	return &domain.Contact{
		DocumentType:   domain.DocumentType(input.DocumentType),
		DocumentNumber: input.DocumentNumber,
		LegalName:      input.LegalName,
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		Address:        input.Address,
		AddressExtra:   input.AddressExtra,
		CityCode:       input.CityCode,
		Phone:          input.Phone,
		Email:          input.Email,
	}
}

// ToDomainPatch converts a ContactPatchInput DTO into a domain.ContactPatch entity.
//
// This function is used to apply partial updates to an existing contact.
func ToDomainPatch(input ContactPatchInput) *domain.ContactPatch {
	return &domain.ContactPatch{
		LegalName:    input.LegalName,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Address:      input.Address,
		AddressExtra: input.AddressExtra,
		CityCode:     input.CityCode,
		Phone:        input.Phone,
		Email:        input.Email,
	}
}

// ToResponseDTO converts a domain.Contact into a ContactResponse DTO.
//
// This function is used when returning contact data in HTTP responses.
func ToResponseDTO(c *domain.Contact) ContactResponse {
	return ContactResponse{
		ID:             c.ID,
		DocumentType:   string(c.DocumentType),
		DocumentNumber: c.DocumentNumber,
		LegalName:      c.LegalName,
		FirstName:      c.FirstName,
		LastName:       c.LastName,
		Address:        c.Address,
		AddressExtra:   c.AddressExtra,
		CityCode:       c.CityCode,
		Phone:          c.Phone,
		Email:          c.Email,
		CreatedAt:      c.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      c.UpdatedAt.Format(time.RFC3339),
	}
}
