package main

import (
	"fmt"
	"math"
)

func main() {
	var pH_val float64
	fmt.Print("Enter the pH value: ")
	fmt.Scan(&pH_val)
	var hydronium_concentration = math.Pow(10, -(pH_val))
	fmt.Printf("\nThe Hydronium Ion concentration is: %f", hydronium_concentration)
}
