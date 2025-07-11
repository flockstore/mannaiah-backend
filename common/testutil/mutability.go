package testutil

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertPatchFieldsMatch checks that the patch struct only contains a subset of allowed fields from the base struct.
// - `base` is the domain model (e.g., Contact)
// - `patch` is the patch struct (e.g., ContactPatch)
// - `excluded` is a list of field names in base that must NOT be patchable (e.g., ID, CreatedAt)
func AssertPatchFieldsMatch(t *testing.T, base, patch any, excluded []string) {
	t.Helper()

	baseFields := map[string]bool{}
	excludedMap := map[string]bool{}
	for _, field := range excluded {
		excludedMap[field] = true
	}

	baseType := reflect.TypeOf(base)
	if baseType.Kind() == reflect.Ptr {
		baseType = baseType.Elem()
	}

	for i := 0; i < baseType.NumField(); i++ {
		name := baseType.Field(i).Name
		if !excludedMap[name] {
			baseFields[name] = true
		}
	}

	patchType := reflect.TypeOf(patch)
	if patchType.Kind() == reflect.Ptr {
		patchType = patchType.Elem()
	}

	for i := 0; i < patchType.NumField(); i++ {
		name := patchType.Field(i).Name
		_, exists := baseFields[name]
		assert.Truef(t, exists, "Field %s exists in patch but is not a modifiable field in the base struct", name)
	}
}
