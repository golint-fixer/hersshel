package store

import (
	"github.com/hersshel/hersshel/model"
	"golang.org/x/net/context"
)

// Store carries methods to manipulate data in a storage.
type Store interface {
	// CreateFeed creates a new feed.
	CreateFeed(*model.Feed) error
	// GetAllFeeds returns all feeds from the store.
	GetAllFeeds() ([]*model.Feed, error)
	// GetFeedItems return a list of model.Item belonging to a given feed.
	GetFeedItems(id uint) ([]*model.Item, error)
	// GetCategoryItems return a list of model.Item
	// belonging to a given category.
	GetCategoryItems(id uint) ([]*model.Item, error)
	// CreateItems inserts multiple model.Item into the store.
	CreateItems(items []*model.Item) error
	// GetAllItems returns all the model.Item from the store.
	GetAllItems() ([]*model.Item, error)
	// CreateCategory creates a new category.
	CreateCategory(*model.Category) error
}

// CreateFeed creates a new feed.
func CreateFeed(c context.Context, feed *model.Feed) error {
	return FromContext(c).CreateFeed(feed)
}

// GetAllFeeds returns all feeds from the store.
func GetAllFeeds(c context.Context) ([]*model.Feed, error) {
	return FromContext(c).GetAllFeeds()
}

// GetFeedItems return a list of model.Item belonging to a given feed.
func GetFeedItems(c context.Context, id uint) ([]*model.Item, error) {
	return FromContext(c).GetFeedItems(id)
}

// GetCategoryItems return a list of model.Item belonging to a given feed.
func GetCategoryItems(c context.Context, id uint) ([]*model.Item, error) {
	return FromContext(c).GetCategoryItems(id)
}

// CreateItems inserts multiple model.Item into the store.
func CreateItems(c context.Context, items []*model.Item) error {
	return FromContext(c).CreateItems(items)
}

// GetAllItems returns all the model.Item from the store.
func GetAllItems(c context.Context) ([]*model.Item, error) {
	return FromContext(c).GetAllItems()
}

// CreateCategory creates a new category.
func CreateCategory(c context.Context, category *model.Category) error {
	return FromContext(c).CreateCategory(category)
}
