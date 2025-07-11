package testutil

import (
	"fmt"
	"reflect"
)

// AssertPatchedFieldsEqual compares all exported fields of a and b for equality.
// Returns false with an error if any field differs or types mismatch.
func AssertPatchedFieldsEqual[T any](a, b T) (bool, error) {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	if va.Kind() == reflect.Ptr {
		va = va.Elem()
	}
	if vb.Kind() == reflect.Ptr {
		vb = vb.Elem()
	}

	if va.Type() != vb.Type() {
		return false, fmt.Errorf("types are different: %s vs %s", va.Type(), vb.Type())
	}

	if va.Kind() != reflect.Struct {
		return false, fmt.Errorf("expected struct type but got %s", va.Kind())
	}

	for i := 0; i < va.NumField(); i++ {
		field := va.Type().Field(i)
		if !field.IsExported() {
			continue
		}

		valueA := va.Field(i)
		valueB := vb.Field(i)

		if !reflect.DeepEqual(valueA.Interface(), valueB.Interface()) {
			return false, fmt.Errorf("field %s: values not equal (%v vs %v)", field.Name, valueA.Interface(), valueB.Interface())
		}
	}

	return true, nil
}
