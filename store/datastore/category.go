package datastore

import "github.com/hersshel/hersshel/model"

// CreateCategory insert a new model.Category in the datastore.
func (ds *datastore) CreateCategory(category *model.Category) error {
	return ds.Insert(category)
}
