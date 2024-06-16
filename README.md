
# **golang-example**

## Introduction

This example Golang service showcases the fundamentals of OpenTelemetry instrumentation by generating **Metrics** and **Traces** data.

## Requirements

1. Install [golang](https://go.dev/doc/install) 
2. ``go get`` the OpenTelemetry libraries:

```
go.opentelemetry.io/otel
go.opentelemetry.io/otel/exporters/stdout/stdoutmetric
go.opentelemetry.io/otel/exporters/stdout/stdouttrace
go.opentelemetry.io/otel/metric
go.opentelemetry.io/otel/propagation
go.opentelemetry.io/otel/sdk/metric
go.opentelemetry.io/otel/sdk/resource
go.opentelemetry.io/otel/sdk/trace
go.opentelemetry.io/otel/semconv/v1.17.0
go.opentelemetry.io/otel/trace
```

## How it works

The project consists of three components: a **server**, a **client**, and an **otelsdk** package.

The **server** package includes an HTTP server with a single ```/hello``` endpoint. This endpoint is instrumented with OpenTelemetry to capture data about request handling. The metric and trace instruments are executed within the HTTP handler function.

The **client** package is responsible for making a GET request to the server's endpoint. It is called from the **main** package once the server is up and running. The client is programmed to automatically send requests every 1 second.

The **otelsdk** package handles the SDK initialization. It declares the controller functions for metrics and traces providers. Additionally, it defines exporters and interval times for each provider. The data is generated every 5 seconds and exported to stdout. The **server** package calls the providers to enable HTTP server instrumentation.

## Run the service

To run this service, you can simply execute the command ``$go run main.go`` from the root directory. 

The expected output should be as follows: 

At every 1 second: 
```
Server: GET
Client: Got response!
Client: HTTP Status code: 200
```

At every 5 seconds (metrics):
```
{
    "Resource": [
        {
            "Key": "service.name",
            "Value": {
                "Type": "STRING",
                "Value": "hello"
            }
        },
        {
            "Key": "service.version",
            "Value": {
                "Type": "STRING",
                "Value": "0.0.0"
            }
        }
    ],
    "ScopeMetrics": [
        {
            "Scope": {
                "Name": "http-metrics",
                "Version": "v0.0.0",
                "SchemaURL": ""
            },
            "Metrics": [
                {
                    "Name": "request.count",
                    "Description": "Number of HTTP requests",
                    "Unit": "{call}",
                    "Data": {
                        "DataPoints": [
                            {
                                "Attributes": [],
                                "StartTime": "2023-11-24T14:25:56.337693-03:00",
                                "Time": "2023-11-24T14:26:01.341815-03:00",
                                "Value": 5
                            }
                        ],
                        "Temporality": "CumulativeTemporality",
                        "IsMonotonic": true
                    }
                },
                {
                    "Name": "request.duration",
                    "Description": "Duration of request execution",
                    "Unit": "sec",
                    "Data": {
                        "DataPoints": [
                            {
                                "Attributes": [],
                                "StartTime": "2023-11-24T14:25:56.337705-03:00",
                                "Time": "2023-11-24T14:26:01.341819-03:00",
                                "Count": 5,
                                "Bounds": [
                                    0,
                                    5,
                                    10,
                                    25,
                                    50,
                                    75,
                                    100,
                                    250,
                                    500,
                                    750,
                                    1000,
                                    2500,
                                    5000,
                                    7500,
                                    10000
                                ],
                                "BucketCounts": [
                                    0,
                                    5,
                                    0,
                                    0,
                                    0,
                                    0,
                                    0,
                                    0,
                                    0,
                                    0,
                                    0,
                                    0,
                                    0,
                                    0,
                                    0,
                                    0
                                ],
                                "Min": {},
                                "Max": {},
                                "Sum": 0.003570917
                            }
                        ],
                        "Temporality": "CumulativeTemporality"
                    }
                }
            ]
        }
    ]
}
```

At every 5 seconds (traces): 

```
{
        "Name": "hello-trace",
        "SpanContext": {
                "TraceID": "802126e27fa5727f590c29c888be300b",
                "SpanID": "d049dd39ac2c1ead",
                "TraceFlags": "01",
                "TraceState": "",
                "Remote": false
        },
        "Parent": {
                "TraceID": "00000000000000000000000000000000",
                "SpanID": "0000000000000000",
                "TraceFlags": "00",
                "TraceState": "",
                "Remote": false
        },
        "SpanKind": 1,
        "StartTime": "2023-11-24T14:25:56.449844-03:00",
        "EndTime": "2023-11-24T14:25:56.453461458-03:00",
        "Attributes": null,
        "Events": null,
        "Links": null,
        "Status": {
                "Code": "Unset",
                "Description": ""
        },
        "DroppedAttributes": 0,
        "DroppedEvents": 0,
        "DroppedLinks": 0,
        "ChildSpanCount": 0,
        "Resource": [
                {
                        "Key": "service.name",
                        "Value": {
                                "Type": "STRING",
                                "Value": "hello"
                        }
                },
                {
                        "Key": "service.version",
                        "Value": {
                                "Type": "STRING",
                                "Value": "0.0.0"
                        }
                }
        ],
        "InstrumentationLibrary": {
                "Name": "hello-trace",
                "Version": "",
                "SchemaURL": ""
        }
}
{
        "Name": "hello-trace",
        "SpanContext": {
                "TraceID": "662b525c130066e9c067ec0623a5e64b",
                "SpanID": "cffc885f69278827",
                "TraceFlags": "01",
                "TraceState": "",
                "Remote": false
        },
        "Parent": {
                "TraceID": "00000000000000000000000000000000",
                "SpanID": "0000000000000000",
                "TraceFlags": "00",
                "TraceState": "",
                "Remote": false
        },
        "SpanKind": 1,
        "StartTime": "2023-11-24T14:25:57.458527-03:00",
        "EndTime": "2023-11-24T14:25:57.459021333-03:00",
        "Attributes": null,
        "Events": null,
        "Links": null,
        "Status": {
                "Code": "Unset",
                "Description": ""
        },
        "DroppedAttributes": 0,
        "DroppedEvents": 0,
        "DroppedLinks": 0,
        "ChildSpanCount": 0,
        "Resource": [
                {
                        "Key": "service.name",
                        "Value": {
                                "Type": "STRING",
                                "Value": "hello"
                        }
                },
                {
                        "Key": "service.version",
                        "Value": {
                                "Type": "STRING",
                                "Value": "0.0.0"
                        }
                }
        ],
        "InstrumentationLibrary": {
                "Name": "hello-trace",
                "Version": "",
                "SchemaURL": ""
        }
}
{
        "Name": "hello-trace",
        "SpanContext": {
                "TraceID": "783b2357683c61cc91fa4aea885a3f7f",
                "SpanID": "ac1bd5fe6f97a452",
                "TraceFlags": "01",
                "TraceState": "",
                "Remote": false
        },
        "Parent": {
                "TraceID": "00000000000000000000000000000000",
                "SpanID": "0000000000000000",
                "TraceFlags": "00",
                "TraceState": "",
                "Remote": false
        },
        "SpanKind": 1,
        "StartTime": "2023-11-24T14:25:58.462788-03:00",
        "EndTime": "2023-11-24T14:25:58.462832792-03:00",
        "Attributes": null,
        "Events": null,
        "Links": null,
        "Status": {
                "Code": "Unset",
                "Description": ""
        },
        "DroppedAttributes": 0,
        "DroppedEvents": 0,
        "DroppedLinks": 0,
        "ChildSpanCount": 0,
        "Resource": [
                {
                        "Key": "service.name",
                        "Value": {
                                "Type": "STRING",
                                "Value": "hello"
                        }
                },
                {
                        "Key": "service.version",
                        "Value": {
                                "Type": "STRING",
                                "Value": "0.0.0"
                        }
                }
        ],
        "InstrumentationLibrary": {
                "Name": "hello-trace",
                "Version": "",
                "SchemaURL": ""
        }
}
{
        "Name": "hello-trace",
        "SpanContext": {
                "TraceID": "290782f501ddbba7444e13c7bc43b96e",
                "SpanID": "3d22709a05cb030f",
                "TraceFlags": "01",
                "TraceState": "",
                "Remote": false
        },
        "Parent": {
                "TraceID": "00000000000000000000000000000000",
                "SpanID": "0000000000000000",
                "TraceFlags": "00",
                "TraceState": "",
                "Remote": false
        },
        "SpanKind": 1,
        "StartTime": "2023-11-24T14:25:59.466028-03:00",
        "EndTime": "2023-11-24T14:25:59.466054416-03:00",
        "Attributes": null,
        "Events": null,
        "Links": null,
        "Status": {
                "Code": "Unset",
                "Description": ""
        },
        "DroppedAttributes": 0,
        "DroppedEvents": 0,
        "DroppedLinks": 0,
        "ChildSpanCount": 0,
        "Resource": [
                {
                        "Key": "service.name",
                        "Value": {
                                "Type": "STRING",
                                "Value": "hello"
                        }
                },
                {
                        "Key": "service.version",
                        "Value": {
                                "Type": "STRING",
                                "Value": "0.0.0"
                        }
                }
        ],
        "InstrumentationLibrary": {
                "Name": "hello-trace",
                "Version": "",
                "SchemaURL": ""
        }
}
{
        "Name": "hello-trace",
        "SpanContext": {
                "TraceID": "ea7cd820f1689b184538c9ba4ff9ac8a",
                "SpanID": "19a97327aac084ba",
                "TraceFlags": "01",
                "TraceState": "",
                "Remote": false
        },
        "Parent": {
                "TraceID": "00000000000000000000000000000000",
                "SpanID": "0000000000000000",
                "TraceFlags": "00",
                "TraceState": "",
                "Remote": false
        },
        "SpanKind": 1,
        "StartTime": "2023-11-24T14:26:00.46842-03:00",
        "EndTime": "2023-11-24T14:26:00.468442333-03:00",
        "Attributes": null,
        "Events": null,
        "Links": null,
        "Status": {
                "Code": "Unset",
                "Description": ""
        },
        "DroppedAttributes": 0,
        "DroppedEvents": 0,
        "DroppedLinks": 0,
        "ChildSpanCount": 0,
        "Resource": [
                {
                        "Key": "service.name",
                        "Value": {
                                "Type": "STRING",
                                "Value": "hello"
                        }
                },
                {
                        "Key": "service.version",
                        "Value": {
                                "Type": "STRING",
                                "Value": "0.0.0"
                        }
                }
        ],
        "InstrumentationLibrary": {
                "Name": "hello-trace",
                "Version": "",
                "SchemaURL": ""
        }
}
```
