package datastore

import "github.com/hersshel/hersshel/model"

// CreateFeed creates a new feed.
func (ds *datastore) CreateFeed(feed *model.Feed) error {
	return ds.Insert(feed)
}

func (ds *datastore) GetAllFeeds() ([]*model.Feed, error) {
	var feeds []*model.Feed
	_, err := ds.Select(&feeds, "SELECT * FROM feed")
	return feeds, err
}
