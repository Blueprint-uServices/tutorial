module github.com/blueprint-uservices/tutorial/examples/helloworld/wiring

go 1.20

require github.com/blueprint-uservices/tutorial/examples/helloworld/workflow v0.0.0
require github.com/blueprint-uservices/tutorial/plugins v0.0.0

require (
	github.com/blueprint-uservices/blueprint/blueprint v0.0.0-20240124230554-8949221e29cc
	github.com/blueprint-uservices/blueprint/plugins v0.0.0-20240124230554-8949221e29cc
)

require (
	github.com/blueprint-uservices/blueprint/runtime v0.0.0-20240120085724-a66c24cd32b1 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/otiai10/copy v1.11.0 // indirect
	go.mongodb.org/mongo-driver v1.12.1 // indirect
	go.opentelemetry.io/otel v1.21.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v0.44.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.20.0 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	go.opentelemetry.io/otel/sdk v1.21.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.21.0 // indirect
	golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/tools v0.15.0 // indirect
)

replace github.com/blueprint-uservices/tutorial/examples/helloworld/workflow => ../workflow
replace github.com/blueprint-uservices/tutorial/plugins => ../../../plugins