package main

import (
	"fmt"
	"math"
)

func main() {
	var hydronium_concentration float64
	var base_concentration float64
	var acid_concentration float64

	fmt.Print("Enter the concentration of Hydronium Ions: ")
	fmt.Scan(&hydronium_concentration)
	fmt.Print("Enter the Proton Donor Concentration: ")
	fmt.Scan(&base_concentration)
	fmt.Print("Enter the Proton Acceptor Concentration: ")
	fmt.Scan(&acid_concentration)

	var numerator = (hydronium_concentration * base_concentration)
	var KA = (numerator / acid_concentration)
	var pKA = -math.Log10(KA)

	fmt.Printf("\nThe KA is %f", KA)
	fmt.Printf("\nThe pKA is %f", pKA)
}
