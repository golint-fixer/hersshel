package model

import (
	"time"

	"gopkg.in/gorp.v1"
)

// Feed describe an RSS/Atom feed in the database.
type Feed struct {
	ID          uint      `db:"id" json:"id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	URL         string    `db:"url" json:"url"`
	Name        string    `db:"name" json:"name"`
	Website     *string   `db:"website" json:"website,omitempty"`
	Description *string   `db:"description" json:"description,omitempty"`
	Image       *string   `db:"image" json:"image,omitempty"`
	CategoryID  uint      `db:"category_id" json:"category_id,omitempty"`
}

// PreInsert is a hook called before inserting into the DB.
func (feed *Feed) PreInsert(s gorp.SqlExecutor) error {
	feed.CreatedAt = time.Now()
	feed.UpdatedAt = feed.CreatedAt
	return nil
}
