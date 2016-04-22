package store

import (
	"github.com/hersshel/hersshel/model"
	"golang.org/x/net/context"
)

// Store carries methods to manipulate data in a storage.
type Store interface {
	// CreateFeed creates a new feed.
	CreateFeed(*model.Feed) error
}

// CreateFeed creates a new feed.
func CreateFeed(c context.Context, feed *model.Feed) error {
	return FromContext(c).CreateFeed(feed)
}
