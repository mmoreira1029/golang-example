package main

import (
	"fmt"
	"golang-service-otel-example/client"
	"golang-service-otel-example/server"
	"log"
	"time"
)

var port = "3333"
var serviceName = "hello"
var serviceVersion = "0.0.0"

func main() {

	srv, err := server.NewServer(serviceName, serviceVersion, port)
	if err != nil {
		log.Fatalf("could not create server: %v", err)
	}

	go func() {
		err := srv.StartServer()
		if err != nil {
			fmt.Println(err)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	url := fmt.Sprintf("http://localhost:%s/hello", port)

	for {
		err := client.SendRequest(url)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
