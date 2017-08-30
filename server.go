package main

import (
	"echo-blog/conf"
	"echo-blog/models"
	"echo-blog/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	models.NewOrm()
	routes.NewRouters(e)
	e.Static("/static", "static")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Start(":" + conf.Config["Port"].(string))
}
