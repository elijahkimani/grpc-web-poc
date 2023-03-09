package main

import (
	"context"
	"log"
	"net"
	"net/http"

	// 'go get...' this below remote packages
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"

	// package is imported with the prefix set as module name
	// and relative path of package as suffix
	gen "example/elijah/grpc-web-poc/proxy/gen/go"
)

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./gen/openapiv2/products.swagger.json")
}

func main() {
	// creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	mux := runtime.NewServeMux()

	// setting up a dial up for gRPC service by specifying endpoint/target url
	err := gen.RegisterProductsServiceHandlerFromEndpoint(context.Background(), mux, "localhost:30043", []grpc.DialOption{grpc.WithInsecure()})

	if err != nil {
		log.Fatal(err)
	}

	httpMux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./proxy/swagger"))
	// take note of the trailing slash '/' to indicate a subtree/wildcard path match
	httpMux.Handle("/docs/", http.StripPrefix("/docs", fileServer))
	httpMux.Handle("/v1/", mux)

	// see https://github.com/rs/cors/blob/master/README.md
	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(httpMux)

	// Creating a normal HTTP server
	server := http.Server{
		Handler: handler,
	}

	// creating a listener for server
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	// start the HTTP server
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
