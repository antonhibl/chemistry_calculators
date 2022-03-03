package main

import (
	"fmt"
	"math"
)

func main() {
	// initialize variable for pH
	var pH_val float64
	// Scan inputs
	fmt.Print("Enter the pH value: ")
	fmt.Scan(&pH_val)
	// calculate the hydronium ion concentration
	var hydronium_concentration = math.Pow(10, -(pH_val))
	// return the outputs
	fmt.Printf("\nThe Hydronium Ion concentration is: %f", hydronium_concentration)
}
