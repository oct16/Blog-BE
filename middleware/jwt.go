package middleware

import "github.com/labstack/echo/middleware"

var JwtAuth = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("waterbear"),
	// AuthScheme: "Bearer",
	TokenLookup: "cookie:token",
})
