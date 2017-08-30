package api

import (
	"echo-blog/models"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// RegisterUser ..
func RegisterUser(c echo.Context) (err error) {

	u := new(models.User)
	err = c.Bind(u)
	if err != nil {
		return err
	}
	user, err := models.RegisterUser(*u)
	if err != nil {
		return c.String(http.StatusNotAcceptable, err.Error())
	}

	json := make(map[string]interface{})
	json["name"] = user.Name
	json["email"] = user.Email
	json["avatar"] = user.Avatar
	jwtToken, _ := MakeJwtToken(user)
	json["token"] = jwtToken

	return c.JSON(http.StatusOK, json)
}

// MakeJwtToken with common user ..
func MakeJwtToken(u models.User) (jwtToken string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = u.Name
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("waterbear_tp"))
	if err != nil {
		return "", err
	}
	return t, nil
}
