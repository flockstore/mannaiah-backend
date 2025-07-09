package domain

import "time"

// Auditable defines common metadata fields for all persistent entities.
type Auditable struct {
	// ID is the unique identifier for the entity.
	ID string `json:"id"`

	// CreatedAt is the timestamp when the entity was first created.
	CreatedAt time.Time `json:"createdAt"`

	// UpdatedAt is the timestamp of the last update to the entity.
	UpdatedAt time.Time `json:"updatedAt"`

	// DeletedAt is the soft-deletion timestamp; nil if the entity is active.
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}
