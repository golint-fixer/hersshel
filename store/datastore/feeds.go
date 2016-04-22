package datastore

import "github.com/hersshel/hersshel/model"

// CreateFeed creates a new feed.
func (ds *datastore) CreateFeed(feed *model.Feed) error {
	return ds.Insert(feed)
}
