package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

const accessKey = "7PDeVhQpOtoeG2dCxyB7GMCWpnzoOxG0J5UrE2D4"
const secretKey = "lbFvqQ1WUc9vBtIT0QqeyeKdy_tI23SZpZupDQJo"
const bucket = "static"

// GetQiniuAccessToken ..
func GetQiniuAccessToken(c echo.Context) (err error) {
	mac := qbox.NewMac(accessKey, secretKey)
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	m := make(map[string]string)
	m["uptoken"] = upToken
	return c.JSON(http.StatusOK, m)
}
