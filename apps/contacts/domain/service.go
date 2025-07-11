package domain

// ContactService defines application-level use cases for managing contacts.
type ContactService interface {
	// Create creates a new contact.
	Create(contact *Contact) error

	// Get retrieves a contact by its ID.
	Get(id string) (*Contact, error)

	// Update applies partial updates to an existing contact.
	Update(id string, patch *ContactPatch) (*Contact, error)

	// Delete removes a contact by its ID.
	Delete(id string) error

	// List retrieves all contacts.
	List() ([]*Contact, error)
}
