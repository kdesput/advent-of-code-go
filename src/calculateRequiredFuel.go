package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculateRequiredFuel() int64 {
	file, err := os.Open("../input/input1")

	if err != nil {
		fmt.Println("ERROR!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var requiredFuel int64

	for scanner.Scan() {
		moduleMass, err2 := strconv.ParseInt(scanner.Text(), 10, 64)
		if err2 != nil {
			fmt.Println("ERROR!")
		}

		requiredFuel += requiredFuelForModule(moduleMass)
	}

	return requiredFuel
}

func requiredFuelForMass(mass int64) int64 {
	return mass/3 - 2
}

func requiredFuelForModule(moduleMass int64) int64 {
	additionalRequiredFuel := requiredFuelForMass(moduleMass)
	var sumOfRequiredFuel int64

	for additionalRequiredFuel > 0 {
		sumOfRequiredFuel += additionalRequiredFuel
		additionalRequiredFuel = requiredFuelForMass(additionalRequiredFuel)
	}

	return sumOfRequiredFuel
}
