# Tendo Patient Feedback Example API

## Purpose

To provide a set of simple APIs to get out sample healthcare data for use in Tendo practice exercises 

## How to Use

1. [Install Go](https://go.dev/doc/install)
2. Start the server: `go run main.go`
3. View API docs at [http://localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html)
4. Query an endpoint with an API Client like [Postman](https://www.postman.com/)
   1. Content-Type = "application/json"
   2. Example URL: `localhost:8000/patients`
   3. Example URL: `localhost:8000/doctors/9bf9e532-93bd-11eb-a8b3-0242ac130003`

## How to Configure

The port and input file can be changed via environment variables.
* To change the port: `export PORT=1234` (Default is 8000)
* To change the input file: `export FILENAME=<filename>.json` (Default is patient-feedback-raw-data.json)

The input data may be modified, but it must only use the same fields as found in 
`patient-feedback-raw-data.json`. Do not introduce new fields or the server may fail to initialize.