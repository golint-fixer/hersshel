package datastore

import "github.com/hersshel/hersshel/model"

// CreateItems inserts multiple model.Item into the store.
func (ds *datastore) CreateItems(items []*model.Item) error {
	tx, err := ds.Begin()
	if err != nil {
		return err
	}
	for _, item := range items {
		_ = tx.Insert(item)
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// GetAllItems returns all the model.Item from the store.
func (ds *datastore) GetAllItems() ([]*model.Item, error) {
	var items []*model.Item
	_, err := ds.Select(&items, "SELECT * FROM item")
	return items, err
}

// GetFeedItems return a list of model.Item belonging to a given feed.
func (ds *datastore) GetFeedItems(id uint) ([]*model.Item, error) {
	var items []*model.Item
	_, err := ds.Select(&items, "SELECT * FROM item WHERE feed_id = :id",
		map[string]interface{}{"id": id})
	return items, err
}
