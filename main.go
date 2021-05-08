package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sflewis2970/factorsandprimes/algebra"
	commonutils "github.com/sflewis2970/factorsandprimes/common"
	"github.com/sflewis2970/factorsandprimes/mathutils"
)

const (
	// FACTOR ...Find factors of values
	FACTOR = "-FACTOR"
	// PRIME ...Determine if numbers are prime numbers
	PRIME = "-PRIME"
	// INPUTDATA ...Use the following values as input
	INPUTDATA = "-INPUTDATA"
	// GENERATE ...Generate {N} numbers
	GENERATE = "-GENERATE"
)

type applOptions struct {
	factor         bool
	prime          bool
	nbrsToGenerate int
	maxValue       int
	nbrsList       []int
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
		case INPUTDATA:
			// Advance to the next string
			idx++

			// Make sure that the idx is still in range
			if idx < applArgsListSize {
				// make sure the next string is NOT the next parameter option
				nextStr := applArgsList[idx]

				if strings.Contains(nextStr, "-") {
					log.Fatal("parameter missing - data values expected\n")
				} else {
					nbrListStr := commonutils.SplitTxt(nextStr, " ")
					nbrsList, err := commonutils.BuildNbrList(nbrListStr)
					if err != nil {
						log.Fatal(err)
					}

					applOptionsData.nbrsList = nbrsList
				}
			} else {
				log.Fatal("parameter missing - data values expected\n")
			}
		case GENERATE:
			// Advance to the next string
			idx++

			// Make sure that the idx is still in range
			if idx < applArgsListSize {
				// make sure the next string is NOT the next parameter option
				nextStr := applArgsList[idx]

				if strings.Contains(nextStr, "-") {
					log.Fatal("parameter missing - data values expected\n")
				} else {
					nbrListStr := commonutils.SplitTxt(nextStr, " ")
					nbrsList, err := commonutils.BuildNbrList(nbrListStr)
					if err != nil {
						log.Fatal(err)
					}

					applOptionsData.nbrsToGenerate = nbrsList[0]
					applOptionsData.maxValue = nbrsList[1]
				}
			} else {
				log.Fatal("parameter missing - data values expected\n")
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

	nbrsListSize := len(applOptionsData.nbrsList)
	if nbrsListSize == 0 && applOptionsData.nbrsToGenerate == 0 && applOptionsData.maxValue == 0 {
		log.Fatal("Usage error: either -inputdata '{data values}' or -generate '{nbrsToGenerate} {maxValue}' is required")
	}

	if nbrsListSize > 0 && (applOptionsData.nbrsToGenerate > 0 && applOptionsData.maxValue > 0) {
		log.Fatal("Usage error: only -inputdata '{data values}' or -generate '{nbrsToGenerate} {maxValue}' is required")
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
		// Exit the application
		log.Fatal("Usage: './main <-factor> <-prime> <<-inputdata> '{numbers separated by space}'> <<-generate> '{nbrOfGeneratedValues} {maxValue}'>'")
	}

	// Generate numbers requested by user input
	nbrsListSize := len(applOptionsData.nbrsList)
	if nbrsListSize == 0 {
		fmt.Printf("nbrs list is empty, generating numbers\n")
		newList, err := mathutils.GenerateInts(applOptionsData.nbrsToGenerate, applOptionsData.maxValue)
		if err != nil {
			errorMsg := fmt.Sprintf("Error generating numbers, err - %v\n", err)
			log.Fatal(errorMsg)
		}

		// Update the numbers list
		applOptionsData.nbrsList = newList
	}

	// Build factors map
	if applOptionsData.factor {
		BuildFactorsMap(applOptionsData.nbrsList)
	}

	// Build primes map
	if applOptionsData.prime {
		BuildPrimesMap(applOptionsData.nbrsList)
	}

	// Ending string
	fmt.Println("Exiting application")
}
