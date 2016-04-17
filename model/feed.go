package model

import (
	"time"

	"gopkg.in/gorp.v1"
)

// Feed describe an RSS/Atom feed.
type Feed struct {
	ID          uint      `db:"id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	URL         string    `db:"url"`
	Name        string    `db:"name"`
	Website     string    `db:"website"`
	Description string    `db:"description"`
	Image       string    `db:"image"`
}

// PreInsert is a hook called before inserting into the DB.
func (feed *Feed) PreInsert(s gorp.SqlExecutor) error {
	feed.CreatedAt = time.Now()
	feed.UpdatedAt = feed.CreatedAt
	return nil
}
