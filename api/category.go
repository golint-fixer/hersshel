package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hersshel/hersshel/errors"
	"github.com/hersshel/hersshel/model"
	"github.com/hersshel/hersshel/store"
	"github.com/lib/pq"
)

type newCategory struct {
	Name string `json:"name"`
}

// PostCategory allows users to add a new category.
func PostCategory(c *gin.Context) {
	var in = &newCategory{}

	err := BindJSON(c, in)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Status:  http.StatusBadRequest,
			Code:    "bad_format",
			Message: "payload is not valid or well-formated.",
		})
		return
	}
	category := &model.Category{
		Name: in.Name,
	}
	err = store.CreateCategory(c, category)
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
			Message: "the server cannot fulfill your request",
		})
		return
	}
	c.JSON(http.StatusCreated, category)
}

// GetCategoryItems returns a list of all the items from a specific category.
func GetCategoryItems(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, errors.Error{
			Status:  http.StatusNotFound,
			Code:    "not_found",
			Message: "feed not found",
		})
		return
	}
	items, err := store.GetCategoryItems(c, uint(id))
	if err != nil {
		if driverErr, ok := err.(*pq.Error); ok {
			c.JSON(http.StatusNotFound, errors.Error{
				Status:  http.StatusNotFound,
				Code:    "not_found",
				Message: driverErr.Detail,
			})
			return
		}
	}
	c.JSON(http.StatusOK, items)
}
