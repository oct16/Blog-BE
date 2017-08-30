package api

import (
	"echo-blog/models"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Login ..
func Login(c echo.Context) (err error) {
	u := new(models.SuperUser)
	err = c.Bind(u)
	if err != nil {
		return err
	}

	user := models.GetSuperUser(*u)
	if user.ID != 0 {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = user.Email
		claims["id"] = user.ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("waterbear"))
		if err != nil {
			return err
		}
		json := make(map[string]interface{})
		json["token"] = t
		jsonUser := make(map[string]interface{})
		jsonUser["email"] = user.Email
		jsonUser["id"] = user.ID
		jsonUser["exp"] = claims["exp"]
		json["user"] = jsonUser

		return c.JSON(http.StatusOK, json)
	}

	return echo.ErrUnauthorized
}

// Verify ..
func Verify(c echo.Context) error {
	type Tk struct {
		Token string
	}
	t := new(Tk)
	err := c.Bind(t)
	if err != nil {
		return err
	}

	tokenString := t.Token

	// return echo.ErrUnauthorized
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("waterbear"), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println(claims["email"])
		return c.String(http.StatusOK, "ok")
	}
	return echo.ErrUnauthorized
}
