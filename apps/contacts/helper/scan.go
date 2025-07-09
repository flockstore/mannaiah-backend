package helper

import (
	"github.com/flockstore/mannaiah-backend/apps/contacts/domain"
	"github.com/jackc/pgx/v5"
)

// ScanContact reads database columns into a Contact entity.
//
// It expects the columns to follow the exact order defined in the SELECT statement.
// Returns a pointer to Contact and any scan error.
func ScanContact(scanner pgx.Row) (*domain.Contact, error) {
	var c domain.Contact

	err := scanner.Scan(
		&c.ID,
		&c.CreatedAt,
		&c.UpdatedAt,
		&c.DeletedAt,
		&c.DocumentType,
		&c.DocumentType,
		&c.DocumentNumber,
		&c.FirstName,
		&c.LastName,
		&c.LegalName,
		&c.Address,
		&c.AddressExtra,
		&c.CityCode,
		&c.DepartmentCode,
		&c.Phone,
		&c.Email,
	)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
