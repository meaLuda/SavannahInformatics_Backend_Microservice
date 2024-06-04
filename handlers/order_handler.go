package handlers

import (
	"customersorders/db"
	"customersorders/services"
	"fmt"
	"net/http"
	"time"
    "log"
	"github.com/labstack/echo/v4"
)

type OrderRequest struct {
    Item       string         `json:"item"`
    Amount     float64        `json:"amount"`
    CustomerID string         `json:"customer_id"`
}


func CreateOrder(c echo.Context) error {
    
    var orderRequest OrderRequest
    if err := c.Bind(&orderRequest); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    // Find customer
    var customer db.Customer
    if err := db.DB.Where("id = ?", orderRequest.CustomerID).First(&customer).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    // Create order
    order := db.Order{
        Item:       orderRequest.Item,
        Amount:     orderRequest.Amount,
        OrderTime:  time.Now(),
        CustomerID: customer.ID,
    }
    if err := db.DB.Create(&order).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    // On successfull send sms to the user
    // Send SMS to customer
    message := fmt.Sprintf("Dear %s, your order for %s has been received and is being processed.", customer.Name, order.Item)
    err := services.SendSMS(message,"+254717478122")
    log.Println(err)
	if err != nil {
		log.Printf("Error sending SMS: %v ", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send SMS"})
    }
    return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Order created and SMS Sent to %s",customer.Name),"data":fmt.Sprintf("Order: %s",order.Item)})}
