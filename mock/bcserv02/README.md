# bcserv

Mock Broadcast service





```json
       "/v1/status": {
            "put": {
                "tags": [
                    "Status Operations"
                ],
                "summary": "Update status of multiple services",
                "requestBody": {
                    "required": true,
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "properties": {
                                    "serviceStatuses": {
                                        "type": "array",
                                        "items": {
                                            "type": "object",
                                            "required": [
                                                "serviceID",
                                                "hmsHSI"
                                            ],
                                            "properties": {
                                                "serviceID": {
                                                    "type": "string"
                                                },
                                                "serviceName": {
                                                    "type": "string"
                                                },
                                                "hmsHSI": {
                                                    "type": "integer",
                                                    "enum": [
                                                        "0",
                                                        "1",
                                                        "2",
                                                        "3",
                                                        "4"
                                                    ]
                                                },
                                                "serviceInstanceID": {
                                                    "type": "string"
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                },
                "responses": {
                    "200": {
                        "description": "Status updated successfully"
                    },
                    "404": {
                        "description": "Services not found",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "type": "object",
                                    "required": [
                                        "message",
                                        "services"
                                    ],
                                    "properties": {
                                        "message": {
                                            "description": "A human readable error message",
                                            "type": "string",
                                            "example": "Few services were not found"
                                        },
                                        "services": {
                                            "description": "List of services that were not found",
                                            "type": "array",
                                            "items": {
                                                "type": "object",
                                                "properties": {
                                                    "serviceID": {
                                                        "type": "string"
                                                    },
                                                    "serviceName": {
                                                        "type": "string"
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Unexpected error",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/Error"
                                }
                            }
                        }
                    }
                }
            },


```




//=====================

"cadence": 5,
"serviceStatuses": [
{

"serviceID": "exampleID",
"serviceName": "exampleService",
"hmsHSI": 1,
"serviceInstanceID": "instance123"
}
// Add more services as needed
]






## Example BC service
- cadance = Num of seconds to rebroadcast ServiceStatus
```json
{
  "bc02ServiceStatus": [
    {
      "cadence": 5,
      "serviceStatuse": 
        {
          "serviceID": "exampleID",
          "serviceName": "exampleService",
          "hmsHSI": 1,
          "serviceInstanceID": "instance123"
        }
      
    }
  ]
}


```
