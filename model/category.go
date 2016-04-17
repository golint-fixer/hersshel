package model

// Category describe a category a Feed could belong to.
type Category struct {
	ID   uint   `db:"id"`
	Name string `db:"name"`
}
