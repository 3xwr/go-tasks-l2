package main

import "fmt"

func getCar(carType string) (icar, error) {
	if carType == "porsche" {
		return newPorsche(), nil
	}

	if carType == "lada" {
		return newLada(), nil
	}

	if carType == "ferrari" {
		return newFerrari(), nil
	}
	return nil, fmt.Errorf("Wrong car type passed")
}
