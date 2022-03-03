package main

import "fmt"

func deltaG(enthalpy float64, temperature float64, entropy float64) (free_energy float64) {
	// calculate change in disorder (delta S * T)
	var temperature_calc = (temperature * entropy)
	// calculate the gibbs free energy, delta G
	free_energy = (enthalpy - temperature_calc)

	// if (delta S * T) > delta H
	if temperature_calc > enthalpy {
		// the reaction is entropy driven
		fmt.Printf("\nThis reaction is Entropy Driven\n")
	} else if temperature_calc < enthalpy { // if (delta S * T) < delta H
		// the reaction is enthalpy driven
		fmt.Printf("\nThis reaction is Enthalpy Driven\n")
	}

	// return the gibbs free energy
	return free_energy
}

func main() {
	// initialize variables for user inputs
	// enthalpy
	var delta_enthalpy float64
	// entropy
	var delta_entropy float64
	// temperature in kelvin
	var kelvin_temp float64

	// take user inputs
	fmt.Print("Enter the change in enthalpy: ")
	fmt.Scan(&delta_enthalpy)
	fmt.Print("Enter the change in entropy: ")
	fmt.Scan(&delta_entropy)
	fmt.Print("Enter the temperature of the system in kelvin: ")
	fmt.Scan(&kelvin_temp)

	// calculate the gibbs free energy
	var gibbs_energy = deltaG(delta_enthalpy, kelvin_temp, delta_entropy)

	// return the amount of free energy in the system available to be used
	fmt.Printf("\nThe amount of free energy in the system available for use is: %f KJ", gibbs_energy)

	// determine whether the reaction is spontaneous, nonspontaneous, or in equillibrium
	// if delta G is positive
	if gibbs_energy > 0 {
		// non-spontaneous
		fmt.Printf("\nThis reaction is Non-spontaneous")
	} else if gibbs_energy < 0 { // if delta G is negative
		// spontaneous
		fmt.Printf("\nThis reaction is Spontaneous")
	} else { // otherwise delta G = 0
		// equllibrium
		fmt.Printf("This reaction is in Equillibrium")
	}
}
