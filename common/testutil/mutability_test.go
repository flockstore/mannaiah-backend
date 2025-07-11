package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// baseStruct simulates a domain model with a mix of mutable and immutable fields.
type baseStruct struct {
	Name      string
	Age       int
	Immutable string
}

// validPatch is a simulated DTO that only includes allowed mutable fields.
type validPatch struct {
	Name *string
	Age  *int
}

// invalidPatch includes a forbidden field (Immutable), simulating a bad update attempt.
type invalidPatch struct {
	Name      *string
	Age       *int
	Immutable *string
}

// TestAssertPatchFieldsMatch_ValidPatch ensures validPatch passes the mutability check.
func TestAssertPatchFieldsMatch_ValidPatch(t *testing.T) {
	base := baseStruct{}
	patch := validPatch{}

	// Should pass: patch only includes mutable fields.
	AssertPatchFieldsMatch(t, base, patch, []string{"Immutable"})
}

// TestAssertPatchFieldsMatch_InvalidField ensures invalidPatch fails due to disallowed field.
func TestAssertPatchFieldsMatch_InvalidField(t *testing.T) {
	base := baseStruct{}
	patch := invalidPatch{}

	// Use a mock testing.T to verify failure without halting the test suite.
	mockT := &testing.T{}
	AssertPatchFieldsMatch(mockT, base, patch, []string{"Immutable"})
	assert.True(t, mockT.Failed())
}

// TestAssertPatchFieldsMatch_WithPointerTypes ensures the check works with pointer struct types.
func TestAssertPatchFieldsMatch_WithPointerTypes(t *testing.T) {
	base := &baseStruct{}
	patch := &validPatch{}

	// Should pass: works with pointers to structs.
	AssertPatchFieldsMatch(t, base, patch, []string{"Immutable"})
}
