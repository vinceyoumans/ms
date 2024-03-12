package bcserv01

type ServiceStatus struct {
	ServiceID         string `json:"serviceID"`
	ServiceName       string `json:"serviceName"`
	HmsHSI            int    `json:"hmsHSI"`
	ServiceInstanceID string `json:"serviceInstanceID"`
}

type UpdateStatusRequest struct {
	ServiceStatuses []ServiceStatus `json:"serviceStatuses"`
}

//===================================

type BC01ServiceStatuses struct {
	Bc01ServiceStatus []struct {
		Cadence        int           `json:"cadence"`
		ServiceStatuse ServiceStatus `json:"serviceStatuse"`
	} `json:"bc01ServiceStatus"`
}
