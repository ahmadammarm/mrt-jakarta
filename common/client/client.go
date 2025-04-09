package client

import (
	"fmt"
	"io"
	"net/http"
)

func ClientRequest(client *http.Client, url string) ([]byte, error) {
    response, err := client.Get(url)

    if err != nil {
        return nil, err
    }

    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("error: %s", response.Status)
    }
    body, err := io.ReadAll(response.Body)

    if err != nil {
        return nil, err
    }

    return body, nil
}