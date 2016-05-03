package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"github.com/hersshel/hersshel/engine"
	"github.com/hersshel/hersshel/errors"
	"github.com/hersshel/hersshel/model"
	"github.com/hersshel/hersshel/store"
)

type newFeed struct {
	URL         string  `json:"url" binding:"required,url"`
	Name        string  `json:"name" binding:"required,gt=0"`
	Website     *string `json:"website" binding:"gt=0,omitempty"`
	Description *string `json:"description" binding:"gt=0,omitempty"`
	Image       *string `json:"image" binding:"url,omitempty"`
}

// PostFeed adds a new RSS feed in the list of feed to aggregate.
// The user sends a JSON described by model.Feed.
func PostFeed(c *gin.Context) {
	var in = &newFeed{}
	var e = engine.FromContext(c)

	err := BindJSON(c, in)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Status:  http.StatusBadRequest,
			Code:    "bad_format",
			Message: "Your payload is not valid or well-formated.",
		})
		return
	}
	feed := &model.Feed{
		URL:         in.URL,
		Name:        in.Name,
		Website:     in.Website,
		Description: in.Description,
		Image:       in.Image,
	}
	err = store.CreateFeed(c, feed)
	if err != nil {
		if driverErr, ok := err.(*pq.Error); ok {
			c.JSON(http.StatusConflict, errors.Error{
				Status:  http.StatusConflict,
				Code:    "conflict_error",
				Message: driverErr.Detail,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, errors.Error{
			Status:  http.StatusInternalServerError,
			Code:    "internal_error",
			Message: err.Error(),
		})
		return
	}
	go e.Schedule(c, feed)
	c.JSON(http.StatusCreated, feed)
}

// GetAllItems returns a list of all the items in the store.
func GetAllItems(c *gin.Context) {
	items, _ := store.GetAllItems(c)
	c.JSON(http.StatusOK, items)
}
