package domain

import (
	"github.com/flockstore/mannaiah-backend/common/domain"
)

// DocumentType represents the type of identification document.
type DocumentType string

const (
	DocumentCC       DocumentType = "CC"    // Cédula de Ciudadanía
	DocumentCE       DocumentType = "CE"    // Cédula de Extranjería
	DocumentTI       DocumentType = "TI"    // Tarjeta de Identidad
	DocumentPassport DocumentType = "PAS"   // Pasaporte
	DocumentNIT      DocumentType = "NIT"   // Número de Identificación Tributaria
	DocumentOther    DocumentType = "OTHER" // Otro
)

// Contact is the aggregate root representing a legal or natural entity.
type Contact struct {
	domain.Auditable

	// ID is the unique identifier in the system.
	ID string

	// DocumentType defines the kind of document used.
	DocumentType DocumentType

	// DocumentNumber is the unique ID number (digits only).
	DocumentNumber string

	// LegalName is used for entities with NIT (e.g. business).
	LegalName string

	// FirstName is used if the contact is a natural person.
	FirstName string

	// LastName is used if the contact is a natural person.
	LastName string

	// Email is the main contact email.
	Email string

	// Phone is the main contact phone number.
	Phone string

	// Address is the main physical address.
	Address string

	// AddressExtra is a complement (e.g. apartment, floor).
	AddressExtra string

	// CityCode represents the city or town code (e.g. DANE).
	CityCode string
}

// ValidateNames  check if name combination is correct
func ValidateNames(legal, first, last string) error {

	if legal != "" && (first != "" || last != "") {
		return ErrInvalidNameCombination
	}
	if legal == "" && (first == "" || last == "") {
		return ErrMissingName
	}
	return nil
}
