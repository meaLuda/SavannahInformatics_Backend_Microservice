package routes

import (
    "customersorders/handlers"
    "github.com/labstack/echo/v4"
)

func SetupRouter() *echo.Echo {
    e := echo.New()

    // Group api
    v1 := e.Group("/api/v1")
    {
        v1.POST("/customers", handlers.CreateCustomer)
        v1.POST("/orders", handlers.CreateOrder)
    }
    return e
}