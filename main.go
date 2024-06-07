package main

import (
    "customersorders/routes"
    "customersorders/db"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    // Setup db connection and migrate db
    conf := db.LoadConfig() // load .env variables
    db.InitDB(conf) // create global db and migrate data.

    e := routes.SetupRouter()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowHeaders: []string{echo.HeaderOrigin,echo.HeaderContentType,echo.HeaderAccept},
    }))
    e.Use(AuthMiddleware())

    e.Logger.Fatal(e.Start(":8080"))
}

func AuthMiddleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            // Authentication code using auth0
            return next(c)
        }
    }
}
