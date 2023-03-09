## GRPC Web Concept

### Introduction

The app was built as a POC of how a web front end can connect to a gRPC server via a gateway proxy

The aim is to have the web front end be able to communicate to the gRPC server using RESTful endpoints

The proxy/intermediary will then handle conversion to and from gRPC calls for the REST calls received

### Approach

The ere are two possible approaches to this

1. Create a gRPC proxy using Go Multiplexer server to expose RESTful API endpoints
2. Create a gRPC client using Node and use express to create RESTful API endpoints
