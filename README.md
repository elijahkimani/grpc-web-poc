## GRPC Web Concept

### Introduction

The app was built as a POC of how a web front end can connect to a gRPC server via a gateway proxy

The aim is to have the web front end be able to communicate to the gRPC server using RESTful endpoints

The proxy/intermediary will then handle conversion to and from gRPC calls for the REST calls received

### Approach

The ere are two possible approaches to this

1. Create a gRPC proxy using Go Multiplexer server to expose RESTful API endpoints
2. Create a gRPC client using Node and use express to create RESTful API endpoints

## Prerequisites

1. Install `protoc`
2. Install `buf` - Helps generate APIs from protocol buffers
3. Install `GoLang` - Runs the MUX proxy server
4. Install `Node js`

## Instructions

1. To start the gRPC server, execute the command `npm run start:grpc-server`
   - This will expose the server in the endpoint [http:\\\\localhost:30043](http:\\localhost:30043)
2. To start the go proxy server, execute the command `go run proxy/main.go`
   - This exposes the proxy server for RESTful calls in [http:\\\\localhost:8081](http:\\localhost:8081)
   - Try accessing the endpoint [http:\\\\localhost:8081\\v1\\products](http:\\localhost:8081\v1\products) in your browser or postman. A list of products should be displayed
3. To start the gRPC Node client, execute the command `npm run start:grpc-client`
