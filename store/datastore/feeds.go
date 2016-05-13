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

// DeleteFeed will delete in the store the given feed.
func (ds *datastore) DeleteFeed(id uint) error {
	var feed model.Feed
	err := ds.SelectOne(&feed, "SELECT * FROM feed WHERE id = $1", id)
	if err != nil {
		return err
	}
	_, err = ds.Delete(&feed)
	if err != nil {
		return err
	}
	return nil
}
