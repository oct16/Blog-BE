package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo"
)

func BodyToJson(c echo.Context) (map[string]interface{}, error) {
	s, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return nil, err
	}

	var body map[string]interface{}
	if err := json.Unmarshal(s, &body); err != nil {
		return nil, err
	}

	return body, nil
}
