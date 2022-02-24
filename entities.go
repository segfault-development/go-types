package types

import "github.com/google/uuid"

type IDEntity struct {
	UUID uuid.UUID `json:"uuid"`
}

type StampedEntity struct {
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	DeletedAt *int64 `json:"deletedAt"`
}
