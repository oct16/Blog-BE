package routes

import (
	"echo-blog/api"
	"echo-blog/middleware"

	"github.com/labstack/echo"
)

// NewRouters ..
func NewRouters(e *echo.Echo) *echo.Echo {
	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.GET("/posts", api.GetPosts)
		v1.GET("/posts/:page", api.GetPosts)
		v1.GET("/post/:id", api.GetPost)
		v1.GET("/post/title/:title", api.GetPostByTitle)

		v1.POST("/post/comment", api.PostComment)

		v1.POST("/login", api.Login)

		v1.POST("/verify", api.Verify)

		v1.POST("/oauth/register", api.RegisterUser)
	}

	// Restricted group
	a := e.Group("/api/v1/admin", middleware.JwtAuth)
	{
		a.GET("/posts", api.GetPosts)
		a.PUT("/post/:id", api.PutPost)
		a.POST("/post", api.NewPost)
		a.DELETE("/post/:id", api.DeletePost)
		a.DELETE("/post/comment/:id", api.DeleteComment)
	}

	all := e.Group("/")
	{
		all.GET("*", api.NotFound)
	}

	return e
}
