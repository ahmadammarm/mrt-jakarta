package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Request(client *http.Client, url string) ([]byte, error) {
    response, err := client.Get(url)

    if err != nil {
        return nil, err
    }

    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("error: %s", response.Status)
    }
    body, err := ioutil.ReadAll(response.Body)

    if err != nil {
        return nil, err
    }
    
    return body, nil
}