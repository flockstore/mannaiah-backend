package domain

// ContactPatch represents partial updates to a contact.
// Fields that are nil will not be updated.
type ContactPatch struct {

	// LegalName is the new legal name (optional).
	LegalName *string

	// FirstName is the new first name (optional).
	FirstName *string

	// LastName is the new last name (optional).
	LastName *string

	// Address is the main address (optional).
	Address *string

	// AddressExtra is a complementary address field (optional).
	AddressExtra *string

	// CityCode is the city or town code (optional).
	CityCode *string

	// Phone is the updated phone number (optional).
	Phone *string

	// Email is the updated email address (optional).
	Email *string
}

// ApplyPatch applies only the non-nil fields from a ContactPatch into the given Contact.
// This ensures safe, explicit mutation and prevents accidental field changes.
func ApplyPatch(c *Contact, patch *ContactPatch) {

	if patch.LegalName != nil {
		c.LegalName = *patch.LegalName
	}
	if patch.FirstName != nil {
		c.FirstName = *patch.FirstName
	}
	if patch.LastName != nil {
		c.LastName = *patch.LastName
	}
	if patch.Address != nil {
		c.Address = *patch.Address
	}
	if patch.AddressExtra != nil {
		c.AddressExtra = *patch.AddressExtra
	}
	if patch.CityCode != nil {
		c.CityCode = *patch.CityCode
	}
	if patch.Phone != nil {
		c.Phone = *patch.Phone
	}
	if patch.Email != nil {
		c.Email = *patch.Email
	}

}
