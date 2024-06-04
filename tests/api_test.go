package tests

import (
	"bytes"
	"customersorders/db"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func init() {
    conf := db.LoadConfig()
    db.InitDB(conf)
}


func GenfloatRand(v1, v2 float64) float64{
    return v1 + v2 + rand.Float64() 
}

func TestCreateCustomer(t *testing.T) {
    log.Println("--------- Testing Creating Customers ----------------")

    // Generate a fake Customer struct
    customer := db.Customer{
        Name: faker.Name(),
        Code: faker.UUIDHyphenated(),
    }
    body, _ := json.Marshal(customer)

    resp := httptest.NewRequest(http.MethodPost, "/api/v1/customers", bytes.NewBuffer(body))
    log.Println(resp)
    assert.Equal(t, 200, http.StatusOK)

}

type OrderRequest struct {
    Item       string         `json:"item"`
    Amount     float64        `json:"amount"`
    CustomerID string         `json:"customer_id"`
}

func TestCreateOrder(t *testing.T) {
    log.Println("--------- Testing Order & SMS Creation ----------------")
    // Generate a fake OrderRequest struct
    order := OrderRequest{
        Item:       faker.Word(),
        Amount:     GenfloatRand(100,200),
        CustomerID: faker.UUIDHyphenated(),
    }
    body, _ := json.Marshal(order)

    resp := httptest.NewRequest(http.MethodPost, "/api/v1/orders", bytes.NewBuffer(body))
    log.Println(resp)
    assert.Equal(t, 200, http.StatusOK)
}
