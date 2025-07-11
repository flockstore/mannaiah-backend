package service

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/flockstore/mannaiah-backend/apps/contacts/domain"
	"github.com/flockstore/mannaiah-backend/apps/contacts/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// newValidContact returns a natural person with first and last name.
func newValidContact() *domain.Contact {
	return &domain.Contact{
		DocumentType:   "CC",
		DocumentNumber: "123456",
		FirstName:      "Ana",
		LastName:       "Gomez",
	}
}

// newLegalEntity returns a legal entity with legal name only.
func newLegalEntity() *domain.Contact {
	return &domain.Contact{
		DocumentType:   "CC",
		DocumentNumber: "654321",
		LegalName:      "Empresa S.A.",
	}
}

// pointer returns a pointer to the given value.
func pointer[T any](v T) *T {
	return &v
}

// TestCreate_Success ensures a valid contact is saved correctly.
func TestCreate_Success(t *testing.T) {
	repo := mocks.NewContactRepository(t)
	svc := NewContactService(repo)
	contact := newValidContact()

	repo.On("GetByDocument", contact.DocumentType, contact.DocumentNumber).Return(nil, domain.ErrContactNotFound)
	repo.On("Save", mock.AnythingOfType("*domain.Contact")).Return(nil)

	err := svc.Create(contact)
	assert.NoError(t, err)
	assert.NotEmpty(t, contact.ID)
	assert.WithinDuration(t, time.Now(), contact.CreatedAt, time.Second)
	assert.Equal(t, contact.CreatedAt, contact.UpdatedAt)
}

// TestCreate_Duplicate checks rejection of duplicate document numbers.
func TestCreate_Duplicate(t *testing.T) {
	repo := mocks.NewContactRepository(t)
	svc := NewContactService(repo)
	contact := newValidContact()

	repo.On("GetByDocument", contact.DocumentType, contact.DocumentNumber).Return(&domain.Contact{}, nil)

	err := svc.Create(contact)
	assert.ErrorIs(t, err, domain.ErrDuplicateDocument)
}

// TestCreate_InvalidNameCombination checks if legal + natural name fails.
func TestCreate_InvalidNameCombination(t *testing.T) {
	repo := mocks.NewContactRepository(t)
	svc := NewContactService(repo)
	contact := newValidContact()
	contact.LegalName = "Empresa S.A."

	repo.On("GetByDocument", contact.DocumentType, contact.DocumentNumber).Return(nil, domain.ErrContactNotFound)

	err := svc.Create(contact)
	assert.ErrorIs(t, err, domain.ErrInvalidNameCombination)
}

// TestCreate_MissingName checks that missing name values are rejected.
func TestCreate_MissingName(t *testing.T) {
	repo := mocks.NewContactRepository(t)
	svc := NewContactService(repo)
	contact := &domain.Contact{
		DocumentType:   "CC",
		DocumentNumber: "999",
	}

	repo.On("GetByDocument", contact.DocumentType, contact.DocumentNumber).Return(nil, domain.ErrContactNotFound)

	err := svc.Create(contact)
	assert.ErrorIs(t, err, domain.ErrMissingName)
}

// TestGet_ReturnsContact validates fetching a contact by ID.
func TestGet_ReturnsContact(t *testing.T) {
	repo := mocks.NewContactRepository(t)
	svc := NewContactService(repo)
	expected := newValidContact()
	expected.ID = "abc"

	repo.On("GetByID", "abc").Return(expected, nil)

	result, err := svc.Get("abc")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

// TestDelete_CallsRepo ensures delete by ID delegates to repo.
func TestDelete_CallsRepo(t *testing.T) {
	repo := mocks.NewContactRepository(t)
	svc := NewContactService(repo)

	repo.On("Delete", "abc").Return(nil)

	err := svc.Delete("abc")
	assert.NoError(t, err)
}

// TestList_ReturnsContacts checks all contacts are returned.
func TestList_ReturnsContacts(t *testing.T) {
	repo := mocks.NewContactRepository(t)
	svc := NewContactService(repo)
	expected := []*domain.Contact{newValidContact()}

	repo.On("List").Return(expected, nil)

	result, err := svc.List()
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

// TestUpdate_Success checks if valid patch updates the contact.
func TestUpdate_Success(t *testing.T) {
	repo := mocks.NewContactRepository(t)
	svc := NewContactService(repo)
	id := "abc"
	existing := newValidContact()
	existing.ID = id

	ph := "5551234"

	patch := &domain.ContactPatch{
		Phone: &ph,
	}

	repo.On("GetByID", id).Return(existing, nil)
	repo.On("Save", mock.AnythingOfType("*domain.Contact")).Return(nil)

	updated, err := svc.Update(id, patch)
	assert.NoError(t, err)
	assert.Equal(t, "5551234", updated.Phone)
}

// TestCreate_LegalEntity_Success checks if legal entity is created correctly.
func TestCreate_LegalEntity_Success(t *testing.T) {
	repo := mocks.NewContactRepository(t)
	svc := NewContactService(repo)
	contact := newLegalEntity()

	repo.On("GetByDocument", contact.DocumentType, contact.DocumentNumber).
		Return(nil, domain.ErrContactNotFound)
	repo.On("Save", mock.AnythingOfType("*domain.Contact")).
		Return(nil)

	err := svc.Create(contact)
	assert.NoError(t, err)
	assert.NotEmpty(t, contact.ID)
	assert.WithinDuration(t, time.Now(), contact.CreatedAt, time.Second)
	assert.Equal(t, contact.CreatedAt, contact.UpdatedAt)
}

// TestUpdate_InvalidCombination checks invalid patch combination.
func TestUpdate_InvalidCombination(t *testing.T) {
	repo := mocks.NewContactRepository(t)
	svc := NewContactService(repo)
	id := "abc"
	existing := newValidContact()
	existing.ID = id

	patch := &domain.ContactPatch{
		LegalName: pointer("Empresa"),
		FirstName: pointer("Carlos"),
	}

	_, err := svc.Update(id, patch)
	fmt.Println(err)
	assert.ErrorIs(t, err, domain.ErrInvalidNameCombination)
}

// TestUpdate_NotFound checks that nil entity without error triggers ErrContactNotFound.
func TestUpdate_NotFound(t *testing.T) {
	repo := mocks.NewContactRepository(t)
	svc := NewContactService(repo)

	repo.On("GetByID", "abc").Return(nil, nil)

	_, err := svc.Update("abc", &domain.ContactPatch{})
	assert.ErrorIs(t, err, domain.ErrContactNotFound)
}

// TestCreate_UnexpectedRepoError ensures repo errors (not ContactNotFound) are propagated.
func TestCreate_UnexpectedRepoError(t *testing.T) {
	repo := mocks.NewContactRepository(t)
	svc := NewContactService(repo)
	contact := newValidContact()

	repo.On("GetByDocument", contact.DocumentType, contact.DocumentNumber).
		Return(nil, assert.AnError)

	err := svc.Create(contact)
	assert.ErrorIs(t, err, assert.AnError)
}

// TestUpdate_SaveFails checks if repo.Save errors are propagated.
func TestUpdate_SaveFails(t *testing.T) {
	repo := mocks.NewContactRepository(t)
	svc := NewContactService(repo)

	id := "abc"
	existing := newValidContact()
	existing.ID = id
	patch := &domain.ContactPatch{
		Email: pointer("new@example.com"),
	}

	repo.On("GetByID", id).Return(existing, nil)
	repo.On("Save", mock.AnythingOfType("*domain.Contact")).Return(assert.AnError)

	_, err := svc.Update(id, patch)
	assert.ErrorIs(t, err, assert.AnError)
}

// TestUpdate_GetByIDError returns early if repository.GetByID fails.
func TestUpdate_GetByIDError(t *testing.T) {
	repo := mocks.NewContactRepository(t)
	svc := NewContactService(repo)

	expectedErr := errors.New("db unavailable")

	repo.On("GetByID", "abc").Return(nil, expectedErr)

	_, err := svc.Update("abc", &domain.ContactPatch{})
	assert.ErrorIs(t, err, expectedErr)
}
