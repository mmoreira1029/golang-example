package client

import (
	"fmt"
	"net/http"
)

func SendRequest(url string) error {

	res, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("Error making HTTP request: %d", err)
	} else {
		fmt.Println("Client: Got response!")
		fmt.Println(fmt.Sprintf("Client: HTTP Status code: %d", res.StatusCode))
	}

	return nil
}
