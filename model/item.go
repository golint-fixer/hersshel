package model

import "time"

// Item describe an entry in a RSS Feed.
type Item struct {
	ID        uint       `db:"id" json:"id"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	Title     string     `db:"title" json:"title,omitempty"`
	Link      string     `db:"link" json:"link,omitempty"`
	Author    *string    `db:"author" json:"author,omitempty"`
	Content   *string    `db:"content" json:"content,omitempty"`
	Read      bool       `db:"read" json:"read,omitempty"`
	Starred   bool       `db:"starred" json:"starred,omitempty"`
	FeedID    uint       `db:"feed_id" json:"feed_id,omitempty"`
}
