package main

import (
	"fmt"
	"math"
)

func main() {
	// Initialize variables to be scanned
	// H3O+ concentration
	var hydronium_concentration float64
	// conjugate base concentration
	var base_concentration float64
	// Acid concentration
	var acid_concentration float64

	// Scan in the inputs
	// hydronium concentration
	fmt.Print("Enter the concentration of Hydronium Ions: ")
	fmt.Scan(&hydronium_concentration)
	// Base concentration
	fmt.Print("Enter the Conjugate Base Concentration: ")
	fmt.Scan(&base_concentration)
	// Acid Concentration
	fmt.Print("Enter the Acid Concentration: ")
	fmt.Scan(&acid_concentration)

	// calculate the numerator of equation
	var numerator = (hydronium_concentration * base_concentration)
	// calculate the KA value
	var KA = (numerator / acid_concentration)
	// calculate the pKA value
	var pKA = -math.Log10(KA)

	// return results
	// KA
	fmt.Printf("\nThe KA is %f", KA)
	// pKA
	fmt.Printf("\nThe pKA is %f", pKA)
}
