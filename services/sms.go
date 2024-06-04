package services

import (
    "bytes"
    "fmt"
    "net/http"
    "net/url"
	"os"
	"log"
)

func SendSMS(message string, toNumbers ...string) error {
	// Prepare request data
	data := url.Values{}
	data.Set("username",os.Getenv("SMS_PROVIDER_USERNAME"))
	data.Set("message", message)
	data.Set("from", os.Getenv("SHORT_CODE"))
	for _, num := range toNumbers {
		data.Add("to", num)
	}

	// Create HTTP client
	client := &http.Client{}

	// Prepare request
	req, err := http.NewRequest("POST", "https://api.sandbox.africastalking.com/version1/messaging", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}
	log.Println(os.Getenv("AFRRICAS_TALKING_SANDBOX_API_KEY"))
	log.Println(os.Getenv("SHORT_CODE"))
	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apiKey", os.Getenv("AFRRICAS_TALKING_SANDBOX_API_KEY"))

	// Make request
	resp, err := client.Do(req)
	// response log
	log.Printf("\n Africas Talking API SMS response: %v \n", resp)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to send SMS: %s", resp.Status)
	}

	return nil
}