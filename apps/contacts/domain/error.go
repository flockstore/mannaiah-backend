package domain

import "errors"

// ErrContactNotFound is returned when a contact is not found in the repository.
var ErrContactNotFound = errors.New("contact not found")

// ErrDuplicateDocument is returned when a contact with the same document number already exists.
var ErrDuplicateDocument = errors.New("duplicate document number")

// ErrInvalidNameCombination is returned when both legal_name and (first_name + last_name) are set.
var ErrInvalidNameCombination = errors.New("invalid name combination: choose either legal_name or first+last name")

// ErrMissingName is returned when neither legal_name nor first_name+last_name are provided.
var ErrMissingName = errors.New("missing required name: provide legal_name or first+last name")
