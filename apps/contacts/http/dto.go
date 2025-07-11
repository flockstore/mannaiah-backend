package http

// ContactInput represents the data required to create a new contact.
type ContactInput struct {
	DocumentType   string `json:"documentType" validate:"required"`           // Document type (e.g. "CC", "TI")
	DocumentNumber string `json:"documentNumber" validate:"required"`         // Unique document identifier
	LegalName      string `json:"legalName"`                                  // Legal name (for legal entities)
	FirstName      string `json:"firstName"`                                  // First name (for individuals)
	LastName       string `json:"lastName"`                                   // Last name (for individuals)
	Address        string `json:"address" validate:"required"`                // Main address (mandatory)
	AddressExtra   string `json:"addressExtra"`                               // Additional address details (optional)
	CityCode       string `json:"cityCode" validate:"required,len=5,numeric"` // 5-digit city code from catalog
	Phone          string `json:"phone" validate:"required,min=8,numeric"`    // Minimum 8-digit phone number
	Email          string `json:"email" validate:"required,email"`            // Valid email address
}

// ContactPatchInput represents a partial update payload for a contact.
type ContactPatchInput struct {
	LegalName    *string `json:"legalName,omitempty"`                                   // Updated legal name
	FirstName    *string `json:"firstName,omitempty"`                                   // Updated first name
	LastName     *string `json:"lastName,omitempty"`                                    // Updated last name
	Address      *string `json:"address,omitempty" validate:"omitempty,required"`       // If present, must not be empty
	AddressExtra *string `json:"addressExtra,omitempty"`                                // Updated extra address (optional)
	CityCode     *string `json:"cityCode,omitempty" validate:"omitempty,len=5,numeric"` // Must be 5 digits if present
	Phone        *string `json:"phone,omitempty" validate:"omitempty,min=8,numeric"`    // Must be at least 8 digits if present
	Email        *string `json:"email,omitempty" validate:"omitempty,email"`            // Must be valid if present
}

// ContactResponse represents the contact data returned to the client.
type ContactResponse struct {
	ID             string `json:"id"`             // Unique contact identifier (UUID)
	DocumentType   string `json:"documentType"`   // Document type (e.g. "CC")
	DocumentNumber string `json:"documentNumber"` // Document number (e.g. 123456789)
	LegalName      string `json:"legalName"`      // Legal name (for legal entities)
	FirstName      string `json:"firstName"`      // First name (for individuals)
	LastName       string `json:"lastName"`       // Last name (for individuals)
	Address        string `json:"address"`        // Main address
	AddressExtra   string `json:"addressExtra"`   // Extra address details
	CityCode       string `json:"cityCode"`       // 5-digit city code
	Phone          string `json:"phone"`          // Phone number
	Email          string `json:"email"`          // Email address
	CreatedAt      string `json:"createdAt"`      // ISO 8601 creation timestamp
	UpdatedAt      string `json:"updatedAt"`      // ISO 8601 last update timestamp
}
