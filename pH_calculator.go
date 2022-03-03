package main

import (
	"fmt"
	"math"
)

func main() {
	// Initialize variables
	var hydronium_concentration float64
	// scan input
	fmt.Print("Enter the Hydronium Ion concentration: ")
	fmt.Scan(&hydronium_concentration)
	// calculate the pH
	var pH_value = -math.Log10(hydronium_concentration)
	// return the results
	fmt.Printf("\nThe pH value is: %f", pH_value)
}
