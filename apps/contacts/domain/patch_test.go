package domain

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/flockstore/mannaiah-backend/common/testutil"
)

// TestPatchContactMatchesDomain ensures PatchContact matches all the allowed fields
// to be updated from base domain model.
func TestPatchContactMatchesDomain(t *testing.T) {
	testutil.AssertPatchFieldsMatch(t, Contact{}, ContactPatch{}, []string{
		"ID", "CreatedAt", "UpdatedAt", "DocumentNumber", "DocumentType",
	})
}

func TestApplyPatchContact(t *testing.T) {
	original := Contact{
		DocumentType:   "CC",
		DocumentNumber: "123456789",
		LegalName:      "Flock S.A.S.",
		FirstName:      "Carlos",
		LastName:       "Ramírez",
		Address:        "Calle 1",
		AddressExtra:   "Apto 202",
		CityCode:       "05001",
		Phone:          "3001234567",
		Email:          "carlos@flock.com",
	}

	// Patch with some updated fields
	newFirstName := "Juan"
	newEmail := "juan@flock.com"
	newPhone := "3009876543"

	patch := &ContactPatch{
		FirstName: &newFirstName,
		Email:     &newEmail,
		Phone:     &newPhone,
	}

	expected := Contact{
		DocumentType:   "CC",
		DocumentNumber: "123456789",
		LegalName:      "Flock S.A.S.",
		FirstName:      newFirstName,
		LastName:       "Ramírez",
		Address:        "Calle 1",
		AddressExtra:   "Apto 202",
		CityCode:       "05001",
		Phone:          newPhone,
		Email:          newEmail,
	}

	// Apply patch and assert field-level equality
	ApplyPatch(&original, patch)
	equal, err := testutil.AssertPatchedFieldsEqual(original, expected)
	require.True(t, equal, "Patch application failed: %v", err)
}
