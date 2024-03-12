package bcserv02

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

type BC02ServiceStatuses struct {
	Bc02ServiceStatus []struct {
		Cadence        int           `json:"cadence"`
		ServiceStatuse ServiceStatus `json:"serviceStatuse"`
	} `json:"bc02ServiceStatus"`
}
