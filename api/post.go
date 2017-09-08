package api

import (
	"echo-blog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// GetPosts ..
func GetPosts(c echo.Context) (err error) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		page = 1
	}
	return c.JSON(http.StatusOK, models.GetPosts(page))
}

// GetPost ..
func GetPost(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("id"))
	post, err := models.GetPost(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, post)
}

// GetPostByTitle ..
func GetPostByTitle(c echo.Context) error {
	var title = c.Param("title")
	post, err := models.GetPostByTitle(title)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if post.ID == 0 {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, post)
}

// PutPost ..
func PutPost(c echo.Context) (err error) {
	ID, _ := strconv.Atoi(c.Param("id"))

	p := new(models.Post)
	err = c.Bind(p)
	if err != nil {
		return err
	}

	res, err := models.PutPost(ID, *p)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// NewPost ..
func NewPost(c echo.Context) (err error) {
	p := new(models.Post)
	c.Bind(p)
	post, err := models.NewPost(*p)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, post)
}

// DeletePost ..
func DeletePost(c echo.Context) (err error) {
	ID, _ := strconv.Atoi(c.Param("id"))
	post, err := models.DeletePost(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, post)
}
