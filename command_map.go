package main

import (
	"fmt"
)

func commandMap(c *config) error {
	var results []string
	for i := 0; i < 20; i++ {
		locationArea, err := getLocationAreas(c.Next)
		if err != nil {
			return err
		}
		c.Previous = c.Next
		c.Next = getNextURL(c.Next, +1)
		results = append(results, locationArea)
	}

	// print results
	for i := 0; i < 20; i++ {
		fmt.Println(results[i])
	}

	return nil
}

func commandMapBack(c *config) error {
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	var results []string
	for i := 0; i < 20; i++ {
		locationArea, err := getLocationAreas(c.Previous)
		if err != nil {
			return err
		}
		c.Next = c.Previous
		c.Previous = getNextURL(c.Previous, -1)
		results = append(results, locationArea)
	}

	// reverse the results
	for i := 19; i >= 0; i-- {
		fmt.Println(results[i])
	}

	return nil
}
