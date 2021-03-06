package main

import (
	"auth/server"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	port := flag.Int("port", 6001, "port to run the http server")
	flag.Parse()

	server := server.New(*port)

	listener, err := server.StartHTTPServer()
	if err != nil {
		log.Printf("Error in creating Http server: %s\n", err)
		os.Exit(1)
	}

	if err = server.StartGrpcServer(listener); err != nil {
		fmt.Printf("Error in starting grpc server %s", err)
		os.Exit(1)
	}

}
