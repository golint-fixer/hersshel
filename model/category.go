package model

// Category describe a category a Feed could belong to.
type Category struct {
	ID   uint   `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
