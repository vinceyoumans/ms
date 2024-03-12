package regserv

type ServiceRequest struct {
	ServiceID          string `json:"serviceID"`
	ServiceName        string `json:"serviceName"`
	ServiceDescription string `json:"serviceDescription"`
	ServiceCategory    string `json:"serviceCategory"`
	ServiceVersion     string `json:"serviceVersion"`
	Active             bool   `json:"active"`
	ServiceInstance    struct {
		Name   string `json:"name"`
		Active bool   `json:"active"`
	} `json:"serviceInstance"`
}

type ServiceResponse struct {
	LocationID      string `json:"locationID"`
	LocationName    string `json:"locationName"`
	ServiceID       string `json:"serviceID"`
	ServiceName     string `json:"serviceName"`
	Active          bool   `json:"active"`
	ServiceInstance struct {
		Name              string `json:"name"`
		Active            bool   `json:"active"`
		ServiceInstanceID string `json:"serviceInstanceID"`
	} `json:"serviceInstance"`
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
