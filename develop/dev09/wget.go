package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func validateWget(wget string) error {
	if wget != "wget" {
		return errors.New("invalid command")
	}
	return nil
}

func getWebSite(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("couldn't connect to the server, bad response: %s", response.Status)
	}
	return response, nil
}

func writeToFile(response *http.Response, filename string) (int64, error) {
	file, err := os.Create(filename + ".txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	size, err := io.Copy(file, response.Body)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	return size, nil
}
