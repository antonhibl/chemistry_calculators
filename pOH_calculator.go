package main

import (
	"fmt"
	"math"
)

func main() {
	// initialize variable for hydroxide ion concentration
	var hydroxide_concentration float64
	// take user inputs
	fmt.Print("Enter the Hydroxide Ion concentation: ")
	fmt.Scan(&hydroxide_concentration)
	// calculate the pOH
	var pOH = -math.Log10(hydroxide_concentration)
	// return output pOH
	fmt.Printf("\nThe Hydroxide Ion concentration is: %f", pOH)
}
