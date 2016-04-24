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
