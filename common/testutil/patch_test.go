package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type sampleStruct struct {
	Name  string
	Age   int
	Email string
}

// TestEqualStructs verifies that two identical structs are considered equal.
func TestEqualStructs(t *testing.T) {
	a := sampleStruct{Name: "Alice", Age: 30, Email: "alice@example.com"}
	b := sampleStruct{Name: "Alice", Age: 30, Email: "alice@example.com"}

	ok, err := AssertPatchedFieldsEqual(a, b)
	assert.True(t, ok)
	assert.NoError(t, err)
}

// TestDifferentFieldValue checks that a struct with one different field fails comparison.
func TestDifferentFieldValue(t *testing.T) {
	a := sampleStruct{Name: "Alice", Age: 30, Email: "alice@example.com"}
	b := sampleStruct{Name: "Alice", Age: 31, Email: "alice@example.com"}

	ok, err := AssertPatchedFieldsEqual(a, b)
	assert.False(t, ok)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Age")
}

// TestPointerStructs verifies that pointer structs are dereferenced and compared correctly.
func TestPointerStructs(t *testing.T) {
	a := &sampleStruct{Name: "Ana", Age: 20}
	b := &sampleStruct{Name: "Ana", Age: 20}

	ok, err := AssertPatchedFieldsEqual(a, b)
	assert.True(t, ok)
	assert.NoError(t, err)
}

// TestUnexportedFieldsIgnored checks that unexported fields are ignored during comparison.
func TestUnexportedFieldsIgnored(t *testing.T) {
	type privateStruct struct {
		ID    string
		value string // unexported field
	}

	a := privateStruct{ID: "1", value: "hidden"}
	b := privateStruct{ID: "1", value: "other"}

	ok, err := AssertPatchedFieldsEqual(a, b)
	assert.True(t, ok)
	assert.NoError(t, err)
}
