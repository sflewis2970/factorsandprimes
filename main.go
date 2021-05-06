package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sflewis2970/factorsandprimes/algebra"
	"github.com/sflewis2970/factorsandprimes/mathutils"
)

const (
	// FACTOR ...Find factors of values
	FACTOR = "-FACTOR"
	// PRIME ...Determine if numbers are prime numbers
	PRIME = "-PRIME"
	// GENERATE ...Generate {N} numbers
	// MAXVALUE ...Max value of generated number
	GENERATE = "-GENERATE"
)

type applOptions struct {
	factor         bool
	prime          bool
	nbrsToGenerate int
	maxValue       int
}

func parseArgs(applArgsList []string) applOptions {
	// Declare applOptionData object
	var applOptionsData applOptions

	// Parse application arguments
	applArgsListSize := len(applArgsList)

	for idx := 0; idx < applArgsListSize; idx++ {
		switch strings.ToUpper(applArgsList[idx]) {
		case FACTOR:
			applOptionsData.factor = true
		case PRIME:
			applOptionsData.prime = true
		case GENERATE:
			// Calculate number of parameters remaining
			nbrsOfParamsRemaining := (applArgsListSize - 1) - idx

			// Parse the remainging parameters
			if nbrsOfParamsRemaining > 0 {
				// advance the loop pointer
				idx++

				if idx < applArgsListSize {
					nbrsToGenerate, err := strconv.Atoi(applArgsList[idx])
					if err != nil {
						errorMsg := fmt.Sprintf("Error: %v\n", err)
						log.Fatal(errorMsg)
					}

					applOptionsData.nbrsToGenerate = nbrsToGenerate
				} else {
					log.Fatal("Error parsing nbrsToGenerate")
				}

				nbrsOfParamsRemaining--
				if nbrsOfParamsRemaining > 0 {
					// advance the loop pointer
					idx++

					if idx < applArgsListSize {
						maxValue, err := strconv.Atoi(applArgsList[idx])
						if err != nil {
							errorMsg := fmt.Sprintf("Error: %v\n", err)
							log.Fatal(errorMsg)
						}

						applOptionsData.maxValue = maxValue
					} else {
						log.Fatal("Error parsing maxValue")
					}
				} else {
					log.Fatal("Error maxValue not found")
				}
			} else {
				log.Fatal("Error: missing generate parameter values")
			}

		default:
			errorMsg := fmt.Sprintf("Error: Option not supported: %s.\n", applArgsList[idx])
			log.Fatal(errorMsg)
		}
	}

	// Verify that either -factor or -prime flag has been applied
	if !applOptionsData.factor && !applOptionsData.prime {
		log.Fatal("Usage error: either -factor or -prime flag is required")
	}

	return applOptionsData
}

func BuildFactorsMap(nbrsList []int) {
	// Add entries into the map
	factorsMap, err := algebra.BuildFactorsMap(nbrsList)

	if err != nil {
		errorMsg := fmt.Sprintf("Error processing Factors, err - %v\n", err)
		log.Fatal(errorMsg)
	}

	for key, factorsList := range factorsMap {
		fmt.Printf("Number: %v, factors: %v\n", key, factorsList)
	}
}

func BuildPrimesMap(nbrsList []int) {
	fmt.Println("Building prime numbers")

	// Add entries into the map
	primesMap, err := algebra.BuildPrimesMap(nbrsList)

	if err != nil {
		errorMsg := fmt.Sprintf("Error processing Primes: err - %v\n", err)
		log.Fatal(errorMsg)
	}

	if len(primesMap) != 0 {
		fmt.Println("The following are prime numbers.")

		for key := range primesMap {
			fmt.Printf("Prime number: %v.\n", key)
		}
	} else {
		fmt.Printf("No prime numbers found.\n")
	}
}

func main() {
	// Starting string
	fmt.Println("Starting application")

	// Strip off the application and store the parameters in the new slice
	applArgs := os.Args[1:]

	// Declare applOptionsData struct
	var applOptionsData applOptions

	if len(applArgs) > 0 {
		// parse args from command-line
		applOptionsData = parseArgs(applArgs)
	} else {
		// Display usage message
		fmt.Println("Usage Error")

		// Exit the application
		log.Fatal("Usage: './main [<-factor>] [<-prime>] [<-generate> [{nbrOfGeneratedValues}] [{maxValue}]]'")
	}

	// Generate numbers requested by user input
	var nbrsList []int
	if applOptionsData.nbrsToGenerate > 0 {
		newList, err := mathutils.GenerateInts(applOptionsData.nbrsToGenerate, applOptionsData.maxValue)

		if err != nil {
			errorMsg := fmt.Sprintf("Error generating numbers, err - %v\n", err)
			log.Fatal(errorMsg)
		}

		// Update the numbers list
		nbrsList = newList
	} else {
		// Check to see if the nbrsList is empty,
		// If the nbrsList is empty, have the user enter the numbers that will be used
		newList, err := algebra.Numbers()

		if err != nil {
			errorMsg := fmt.Sprintf("Error building factors, err - %v\n", err)
			log.Fatal(errorMsg)
		}

		// Updatge the numbers list
		nbrsList = newList
	}

	// Build factors map
	if applOptionsData.factor {
		BuildFactorsMap(nbrsList)
	}

	// Build primes map
	if applOptionsData.prime {
		BuildPrimesMap(nbrsList)
	}

	// Ending string
	fmt.Println("Exiting application")
}
