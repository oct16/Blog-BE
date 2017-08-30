package api

import (
	"echo-blog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// PostComment ..
// @Params [post_id, user_id, content ]
func PostComment(c echo.Context) (err error) {
	cm := new(models.Comment)
	err = c.Bind(cm)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	cm.IP = c.RealIP()
	comment, err := models.NewComment(*cm)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, comment)
}

// DeleteComment ..
func DeleteComment(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	comment, err := models.DeleteComment(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, comment)
}
