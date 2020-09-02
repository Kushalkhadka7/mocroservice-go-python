package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	port := flag.Int("port", 8080, "port to run the http server")
	flag.Parse()

	server := NewServer(*port)

	listener, err := server.StartServer()
	if err != nil {
		fmt.Printf("Error in creating Http server%s\n", err)
		os.Exit(1)
	}

	err = server.StartGrpcServer(listener)

	if err != nil {
		fmt.Printf("Error in starting grpc server %s", err)
		os.Exit(1)
	}

	fmt.Println(err)
}
