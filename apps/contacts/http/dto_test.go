package http

import (
	"github.com/flockstore/mannaiah-backend/apps/contacts/domain"
	"github.com/flockstore/mannaiah-backend/common/testutil"
	"testing"
)

// TestDTOFieldParity ensures that DTOs mirrors the allowed fields from the domain.Contact.
// Fields like CreatedAt and UpdatedAt are excluded due to type differences (string vs time.Time).
func TestDTOFieldParity(t *testing.T) {
	testutil.AssertPatchFieldsMatch(t, domain.Contact{}, ContactInput{}, []string{
		"CreatedAt",
		"UpdatedAt",
	})
	testutil.AssertPatchFieldsMatch(t, domain.ContactPatch{}, ContactPatchInput{}, []string{
		"CreatedAt",
		"UpdatedAt",
	})
}
