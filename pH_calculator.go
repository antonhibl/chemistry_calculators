package main

import (
	"fmt"
	"math"
)

func main() {
	var hydronium_concentration float64

	fmt.Print("Enter the Hydronium Ion concentration: ")
	fmt.Scan(&hydronium_concentration)
	var pH_value = -math.Log10(hydronium_concentration)
	fmt.Printf("\nThe pH value is: %f", pH_value)
}
