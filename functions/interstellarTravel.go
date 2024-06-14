package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Println("Remaining fuel:", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 7000000
	default:
		fuel = 0
	}
	return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	fmt.Println("Welcome to", planet)
}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else if fuelRemaining > fuelCost {
		cantFly()
	}

	fuelCost = calculateFuel(planet)
	return fuelRemaining
}

func main() {
	// Test your functions!

	// Create `planetChoice` and `fuel`
	var fuel int = 1000000
	planetChoice := "Venus"
	// And then liftoff!

	fuel = flyToPlanet(planetChoice, fuel)

	fuelGauge(fuel)

}
