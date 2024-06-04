package handlers


import (
    "net/http"
    "customersorders/db"
    "github.com/labstack/echo/v4"
    "fmt"
)

func CreateCustomer(c echo.Context) error {
    var customer db.Customer
    if err := c.Bind(&customer); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    if err := db.DB.Create(&customer).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Customer  %s created",customer.Name)})
}

