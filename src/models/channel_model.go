package models

import "time"

type Channel struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	OwnerID          int64     `json:"owner_id"`
	SubscribersCount int64     `json:"subscribers_count"`
	InviteLink       string    `json:"invite_link,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at,omitempty"`
}
