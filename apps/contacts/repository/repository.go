package repository

import (
	"context"

	"github.com/flockstore/mannaiah-backend/apps/contacts/domain"
	"github.com/flockstore/mannaiah-backend/apps/contacts/helper"
	"github.com/flockstore/mannaiah-backend/common/database"
)

// postgresContactRepository implements domain.ContactRepository using PostgreSQL and pgx.
type postgresContactRepository struct {
	db database.DB
}

// NewPostgresContactRepository creates a new instance of ContactRepository using PostgreSQL.
func NewPostgresContactRepository(db database.PgxClient) domain.ContactRepository {
	return &postgresContactRepository{db: &db}
}

// Save inserts or updates a Contact in the database.
// Assumes the Contact entity has already been fully constructed (ID, timestamps, etc.) by the domain/service layer.
func (r *postgresContactRepository) Save(c *domain.Contact) error {
	query := `
		INSERT INTO contacts (
			id, doc_type, doc_number, legal_name,
			first_name, last_name, address, address_extra,
			city_code, phone, email,
			created_at, updated_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)
		ON CONFLICT (id) DO UPDATE SET
			doc_type=$2, doc_number=$3, legal_name=$4,
			first_name=$5, last_name=$6, address=$7, address_extra=$8,
			city_code=$9, phone=$10, email=$11, created_at=$12, updated_at=$13
	`

	_, err := r.db.Exec(context.Background(), query,
		c.ID, c.DocumentType, c.DocumentNumber, c.LegalName,
		c.FirstName, c.LastName, c.Address, c.AddressExtra,
		c.CityCode, c.Phone, c.Email,
		c.CreatedAt, c.UpdatedAt,
	)
	return err
}

// GetByID retrieves a Contact by its ID.
func (r *postgresContactRepository) GetByID(id string) (*domain.Contact, error) {
	query := `
		SELECT id, doc_type, doc_number, legal_name, first_name, last_name,
		       address, address_extra, city_code, phone, email,
		       created_at, updated_at
		FROM contacts
		WHERE id = $1
	`

	row := r.db.QueryRow(context.Background(), query, id)
	return helper.ScanContact(row)
}

// GetByDocument retrieves a Contact by its document type and number.
func (r *postgresContactRepository) GetByDocument(docType domain.DocumentType, docNumber string) (*domain.Contact, error) {
	query := `
		SELECT id, doc_type, doc_number, legal_name, first_name, last_name,
		       address, address_extra, city_code, phone, email,
		       created_at, updated_at
		FROM contacts
		WHERE doc_type = $1 AND doc_number = $2
	`

	row := r.db.QueryRow(context.Background(), query, docType, docNumber)
	return helper.ScanContact(row)
}

// Delete removes a Contact by ID from the database.
func (r *postgresContactRepository) Delete(id string) error {
	query := `DELETE FROM contacts WHERE id = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}

// List returns all Contacts from the database.
func (r *postgresContactRepository) List() ([]*domain.Contact, error) {
	query := `
		SELECT id, doc_type, doc_number, legal_name, first_name, last_name,
		       address, address_extra, city_code, phone, email,
		       created_at, updated_at
		FROM contacts
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []*domain.Contact
	for rows.Next() {
		c, err := helper.ScanContact(rows)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}
	return contacts, nil
}
