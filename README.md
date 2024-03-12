# mockService


CLI that Simulates AirSide SERVICES


## regserv
This simulates a one time registration of service.
currently the service json is hardcoded but

TODO:  
create json file of services to register



## bcserv01
to be removed



## bcserv02

A Broadcast PUT update service to collector.
Parameters include:
- URL
- cadance: number of seconds between broadcast
- type ServiceStatus struct {
  ServiceID         string `json:"serviceID"`
  ServiceName       string `json:"serviceName"`
  HmsHSI            int    `json:"hmsHSI"`
  ServiceInstanceID string `json:"serviceInstanceID"`
  }

Broadcasts can be modified in the./mockJSON/BCServ02.json file.




TODO:
- update Documentation


TODO:
- create demonstration airside Services endpoint to demonstrate MockServ 



