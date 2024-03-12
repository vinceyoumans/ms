package regserv

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func RegisterService00() {
	// assume localhost for now
	//url00 := "http://example.com/v1/services"      // Replace with your actual endpoint
	//port := "8080"                                 // can adjust this in other code
	//url01 := "http://example.com:8080/v1/services" // Replace with your actual endpoint and port
	//url := "http://example.com:8080/v1/services"   // Replace with your actual endpoint and port
	url := "http://localhost:8080/v1/services" // Assume running on localhost and port 8080

	serviceRequest := ServiceRequest{
		ServiceID:          "123",
		ServiceName:        "ExampleService",
		ServiceDescription: "Description of the service",
		ServiceCategory:    "Category",
		ServiceVersion:     "1.0",
		Active:             true,
		ServiceInstance: struct {
			Name   string `json:"name"`
			Active bool   `json:"active"`
		}{
			Name:   "MapsUI - 1",
			Active: true,
		},
	}

	err := Reg_service01(url, serviceRequest)
	if err != nil {
		handleError(0, "Registration error:", err)
	}

}

func Reg_service01(url string, serviceRequest ServiceRequest) error {
	jsonPayload, err := json.Marshal(serviceRequest)
	if err != nil {
		return fmt.Errorf("Error marshaling JSON: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("Error making POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var serviceResponse ServiceResponse
		err := json.NewDecoder(resp.Body).Decode(&serviceResponse)
		if err != nil {
			return fmt.Errorf("Error decoding response JSON: %v", err)
		}
		fmt.Println("Service registered successfully:", serviceResponse)
		return nil
	} else {
		var errorResponse ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return fmt.Errorf("Error decoding error response JSON: %v", err)
		}
		return errors.New(fmt.Sprintf("Error (Status Code %d): %s", errorResponse.StatusCode, errorResponse.Message))
	}
}

func handleError(statusCode int, message string, err error) {
	fmt.Printf("%s %v\n", message, err)
	// You can handle the error here or log it as needed
}
