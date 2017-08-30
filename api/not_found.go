package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// NotFound ..
func NotFound(c echo.Context) (err error) {
	return c.String(http.StatusNotFound, "哈哈，啥也没有，地址打错了吧")
}
