package http

// ContactInput represents the data required to create a new contact.
type ContactInput struct {
	DocumentType   string `json:"documentType"`   // Type of document (e.g. "CC", "TI", etc.)
	DocumentNumber string `json:"documentNumber"` // Unique document identifier (e.g. ID number)
	LegalName      string `json:"legalName"`      // Full legal name (used for legal entities)
	FirstName      string `json:"firstName"`      // First name (used for individuals)
	LastName       string `json:"lastName"`       // Last name (used for individuals)
	Address        string `json:"address"`        // Main address of the contact
	AddressExtra   string `json:"addressExtra"`   // Additional address information (optional)
	CityCode       string `json:"cityCode"`       // City code following standard catalog
	Phone          string `json:"phone"`          // Contact phone number
	Email          string `json:"email"`          // Contact email address
}

// ContactPatchInput allows partial updates to a contact.
// All fields are optional and represented as pointers to detect changes.
type ContactPatchInput struct {
	LegalName    *string `json:"legalName,omitempty"`    // Updated legal name
	FirstName    *string `json:"firstName,omitempty"`    // Updated first name
	LastName     *string `json:"lastName,omitempty"`     // Updated last name
	Address      *string `json:"address,omitempty"`      // Updated main address
	AddressExtra *string `json:"addressExtra,omitempty"` // Updated address extra info
	CityCode     *string `json:"cityCode,omitempty"`     // Updated city code
	Phone        *string `json:"phone,omitempty"`        // Updated phone number
	Email        *string `json:"email,omitempty"`        // Updated email address
}

// ContactResponse represents the data returned to the client after operations.
type ContactResponse struct {
	ID             string `json:"id"`             // Unique identifier of the contact
	DocumentType   string `json:"documentType"`   // Document type (e.g. "CC")
	DocumentNumber string `json:"documentNumber"` // Document number (e.g. 123456789)
	LegalName      string `json:"legalName"`      // Legal name (for legal entities)
	FirstName      string `json:"firstName"`      // First name (for individuals)
	LastName       string `json:"lastName"`       // Last name (for individuals)
	Address        string `json:"address"`        // Main address
	AddressExtra   string `json:"addressExtra"`   // Additional address details
	CityCode       string `json:"cityCode"`       // City code
	Phone          string `json:"phone"`          // Phone number
	Email          string `json:"email"`          // Email address
	CreatedAt      string `json:"createdAt"`      // ISO8601 timestamp of creation
	UpdatedAt      string `json:"updatedAt"`      // ISO8601 timestamp of last update
}
