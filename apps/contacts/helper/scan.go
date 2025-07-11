package helper

import (
	"errors"
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
		&c.ID, &c.DocumentType, &c.DocumentNumber, &c.LegalName,
		&c.FirstName, &c.LastName, &c.Address, &c.AddressExtra,
		&c.CityCode, &c.Phone, &c.Email,
		&c.CreatedAt, &c.UpdatedAt, &c.DeletedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrContactNotFound
	}

	if err != nil {
		return nil, err
	}

	return &c, nil
}
