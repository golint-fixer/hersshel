package model

import "time"

// Item describe an entry in a RSS Feed.
type Item struct {
	ID        uint       `db:"id"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	Title     string     `db:"title"`
	Author    string     `db:"author"`
	Content   string     `db:"content"`
	Link      string     `db:"link"`
	Read      bool       `db:"read"`
	Starred   bool       `db:"starred"`
	FeedID    uint       `db:"feed_id"`
}
