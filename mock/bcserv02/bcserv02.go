package bcserv02

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func BC02Serv(jsonFile, url string) {
	jsonData, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var bc02ServiceStatuses BC02ServiceStatuses
	err = json.Unmarshal(jsonData, &bc02ServiceStatuses)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data:", err)
		return
	}

	var wg sync.WaitGroup

	for _, bc02ServiceStatus := range bc02ServiceStatuses.Bc02ServiceStatus {
		wg.Add(1)
		go StartMockService02(&wg, url, bc02ServiceStatus.Cadence, bc02ServiceStatus.ServiceStatuse)
	}

	wg.Wait()

}

// ==================================================

// StartMockService02 starts a mockservice from the MockJSON file.  go function
func StartMockService02(wg *sync.WaitGroup, url string, cadence int, serviceStatus ServiceStatus) {
	defer wg.Done()

	fmt.Printf("Starting mock service with cadence: %d, ServiceID: %s\n", cadence, serviceStatus.ServiceID)

	x := 0
	for {
		x++
		response, err := putStatusUpdate(url, serviceStatus)
		if err != nil {

			fmt.Println(x, "   xx Error:", err)
		} else {
			fmt.Println("Response:", response)
		}

		time.Sleep(time.Duration(cadence) * time.Second)

		// TODO:  send status update to MAP and display at console in a more elegant way
	}

}

// ==================================================

// ==================================================

func putStatusUpdate(url string, serviceStatus ServiceStatus) (string, error) {
	requestBody, err := json.Marshal(map[string]ServiceStatus{
		"serviceStatus": serviceStatus,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return "Status updated successfully", nil
	} else if resp.StatusCode == http.StatusNotFound {
		var notFoundResponse map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&notFoundResponse)
		if err != nil {
			return "", err
		}
		message := notFoundResponse["message"].(string)
		services := notFoundResponse["services"].([]interface{})
		return fmt.Sprintf("Error: %s. Services not found: %v", message, services), nil
	} else {
		return "", fmt.Errorf("Unexpected error: %d", resp.StatusCode)
	}
}

//==================================================

//==================================================

//==================================================
