package bcserv01

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

func BCServ01(urlEndPoint string) {
	// Create an example update request
	updateRequest := UpdateStatusRequest{
		ServiceStatuses: []ServiceStatus{
			{
				ServiceID:         "exampleID",
				ServiceName:       "exampleService",
				HmsHSI:            1,
				ServiceInstanceID: "instance123",
			},
			// Add more services as needed
		},
	}

	// Convert the request to JSON
	requestBody, err := json.Marshal(updateRequest)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Define the API endpoint
	//apiURL := "http://example.com/v1/status"
	apiURL := urlEndPoint

	// Create a custom Dialer with a timeout
	dialer := &net.Dialer{
		Timeout:   5 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	// Create a custom Transport with the Dialer
	transport := &http.Transport{
		DialContext: dialer.DialContext,
	}

	// Create an HTTP client with the custom Transport
	client := &http.Client{
		Transport: transport,
	}

	// Create a PUT request
	request, err := http.NewRequest("PUT", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the request headers
	request.Header.Set("Content-Type", "application/json")

	// Make the request
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode == http.StatusOK {
		fmt.Println("Status updated successfully")
	} else if response.StatusCode == http.StatusNotFound {
		// Handle 404 response
		fmt.Println("Services not found")

		// Parse the error response JSON if needed
		// var errorResponse ErrorResponse
		// json.NewDecoder(response.Body).Decode(&errorResponse)
		// fmt.Println("Error message:", errorResponse.Message)
	} else {
		// Handle other response codes
		fmt.Println("Unexpected error:", response.Status)
	}
}
