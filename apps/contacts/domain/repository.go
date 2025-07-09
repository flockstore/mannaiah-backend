package domain

// ContactRepository defines the behavior required to persist and retrieve contacts.
type ContactRepository interface {
	// Save inserts a new contact or updates it if it already exists.
	Save(contact *Contact) error

	// GetByID fetches a contact by its unique ID.
	GetByID(id string) (*Contact, error)

	// GetByDocument finds a contact using document type and number.
	GetByDocument(docType DocumentType, docNumber string) (*Contact, error)

	// Delete removes a contact by its ID.
	Delete(id string) error

	// List returns all contacts (could be paginated later).
	List() ([]*Contact, error)
}
