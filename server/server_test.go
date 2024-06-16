package server

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

var port = "3333"
var serviceName = "hello"
var serviceVersion = "0.0.0"
var url = fmt.Sprintf("http://localhost:%s/hello", port)

func TestServerResponse(t *testing.T) {

	srv, err := NewServer(serviceName, serviceVersion, port)
	if err != nil {
		fmt.Println(fmt.Sprintf("could not create server: %v", err))
	}

	go func() {
		err := srv.StartServer()
		if err != nil {
			fmt.Println(err)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	response, err := http.Get(url)

	if err != nil {
		t.Fatalf("HTTP Request failed with error: %d", err)
	} else if response.StatusCode != http.StatusOK {
		t.Fatalf("HTTP Request failed with status: %v", response.Status)
	} else {
		fmt.Println(fmt.Sprintf("HTTP Status code: %d", response.StatusCode))
	}

}
