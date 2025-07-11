package http

import (
	bdomain "github.com/flockstore/mannaiah-backend/common/domain"
	"testing"
	"time"

	"github.com/flockstore/mannaiah-backend/apps/contacts/domain"
	"github.com/flockstore/mannaiah-backend/common/testutil"
)

// TestToDomainContact validates that ToDomainContact correctly maps a ContactInput
// into a domain.Contact object with matching fields.
func TestToDomainContact(t *testing.T) {
	input := ContactInput{
		DocumentType:   "CC",
		DocumentNumber: "123456789",
		LegalName:      "John Doe",
		FirstName:      "John",
		LastName:       "Doe",
		Address:        "Main St",
		AddressExtra:   "Apt 101",
		CityCode:       "05001",
		Phone:          "3001234567",
		Email:          "john@example.com",
	}

	expected := &domain.Contact{
		DocumentType:   "CC",
		DocumentNumber: "123456789",
		LegalName:      "John Doe",
		FirstName:      "John",
		LastName:       "Doe",
		Address:        "Main St",
		AddressExtra:   "Apt 101",
		CityCode:       "05001",
		Phone:          "3001234567",
		Email:          "john@example.com",
	}

	actual := ToDomainContact(input)

	ok, err := testutil.AssertPatchedFieldsEqual(expected, actual)
	if !ok {
		t.Errorf("ToDomainContact failed: %v", err)
	}
}

// TestToResponseDTO checks that ToResponseDTO properly converts a domain.Contact
// into a ContactResponse DTO, including time formatting.
func TestToResponseDTO(t *testing.T) {
	now := time.Now()
	contact := &domain.Contact{
		ID:             "uuid-123",
		DocumentType:   "CC",
		DocumentNumber: "987654321",
		LegalName:      "Jane Smith",
		FirstName:      "Jane",
		LastName:       "Smith",
		Address:        "2nd Ave",
		AddressExtra:   "Suite 202",
		CityCode:       "11001",
		Phone:          "3012345678",
		Email:          "jane@example.com",
		Auditable: bdomain.Auditable{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	expected := ContactResponse{
		ID:             "uuid-123",
		DocumentType:   "CC",
		DocumentNumber: "987654321",
		LegalName:      "Jane Smith",
		FirstName:      "Jane",
		LastName:       "Smith",
		Address:        "2nd Ave",
		AddressExtra:   "Suite 202",
		CityCode:       "11001",
		Phone:          "3012345678",
		Email:          "jane@example.com",
		CreatedAt:      now.Format(time.RFC3339),
		UpdatedAt:      now.Format(time.RFC3339),
	}

	actual := ToResponseDTO(contact)

	ok, err := testutil.AssertPatchedFieldsEqual(expected, actual)
	if !ok {
		t.Errorf("ToResponseDTO failed: %v", err)
	}
}

// TestToDomainPatch validates that ToDomainPatch correctly maps a ContactPatchInput
// into a domain.ContactPatch struct, preserving pointer semantics.
func TestToDomainPatch(t *testing.T) {
	name := "Alice"
	last := "Doe"
	phone := "3001112233"
	legalName := "Alice Legal"
	address := "Calle 123"
	email := "alice@example.com"

	input := ContactPatchInput{
		FirstName: &name,
		LastName:  &last,
		Phone:     &phone,
		LegalName: &legalName,
		Address:   &address,
		Email:     &email,
	}

	expected := &domain.ContactPatch{
		FirstName: &name,
		LastName:  &last,
		Phone:     &phone,
		LegalName: &legalName,
		Address:   &address,
		Email:     &email,
	}

	actual := ToDomainPatch(input)

	ok, err := testutil.AssertPatchedFieldsEqual(expected, actual)
	if !ok {
		t.Errorf("ToDomainPatch failed: %v", err)
	}
}
