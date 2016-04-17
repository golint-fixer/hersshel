package model

import (
	"time"

	"gopkg.in/gorp.v1"
)

// Item describe an entry in a RSS Feed.
type Item struct {
	ID        uint      `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Title     string    `db:"title"`
	Author    string    `db:"author"`
	Content   string    `db:"content"`
	link      string    `db:"link"`
	Read      bool      `db:"read"`
	Starred   bool      `db:"starred"`
}

// PreInsert is a hook called before inserting into the DB.
func (item *Item) PreInsert(s gorp.SqlExecutor) error {
	item.CreatedAt = time.Now()
	item.UpdatedAt = item.CreatedAt
	return nil
}
