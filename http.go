package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getLocationAreas(url string) (string, error) {
	// creating GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating GET request: %v", err)
	}

	// make request, get and process response
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error getting location areas: %v", err)
	}
	defer res.Body.Close()

	// decode response
	var data LocationArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return "", fmt.Errorf("error decoding response JSON: %v", err)
	}

	return data.Name, nil
}
