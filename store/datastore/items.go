package datastore

import (
	"github.com/Sirupsen/logrus"
	"github.com/hersshel/hersshel/model"
)

// CreateItems inserts multiple model.Item into the store.
func (ds *datastore) CreateItems(items []*model.Item) error {
	tx, err := ds.Begin()
	if err != nil {
		logrus.Infof("error transaction: %v", err)
		return err
	}
	for _, item := range items {
		err = tx.Insert(item)
		if err != nil {
			logrus.Infof("error inserting: %v", err)
		}
	}
	err = tx.Commit()
	if err != nil {
		logrus.Infof("error committing: %v", err)
		return err
	}
	return nil
}
