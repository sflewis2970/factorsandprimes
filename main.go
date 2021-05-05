package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sflewis2970/factorsandprimes/algebra"
	"github.com/sflewis2970/factorsandprimes/mathutils"
)

const (
	// FACTOR ...Find factors of values
	FACTOR = "-factor"
	// PRIME ...Determine if numbers are prime numbers
	PRIME = "-prime"
	// GENERATE ...Generate {N} numbers
	GENERATE = "-generate"
	// MAXVALUE ...Max value of generated number
	MAXVALUE = "-maxvalue"
)

type applOptions struct {
	factor         bool
	prime          bool
	nbrsToGenerate int
	maxValue       int
}

func parseArgs(applArgsList []string) applOptions {
	fmt.Println("Entering parseArgs...")

	// Declare applOptionData object
	var applOptionsData applOptions

	// Parse application arguments
	applArgsListSize := len(applArgsList)
	for idx := 0; idx < applArgsListSize; idx++ {
		switch applArgsList[idx] {
		case FACTOR:
			applOptionsData.factor = true
		case PRIME:
			applOptionsData.prime = true
		case GENERATE:
			// Capture the numbers to generate value
			nbrsToGenerateItem := idx + 1
			if nbrsToGenerateItem < applArgsListSize {
				nbrsToGenerate, err := strconv.Atoi(applArgsList[nbrsToGenerateItem])
				errorMsg := fmt.Sprintf("Error: %v\n", err)
				if err != nil {
					log.Fatal(errorMsg)
				}

				applOptionsData.nbrsToGenerate = nbrsToGenerate
			} else {
				log.Fatal("Error: invalid parameter given.")
			}

			// Capture the numbers to generate value
			maxValueItem := idx + 2
			if maxValueItem < applArgsListSize {
				maxValue, err := strconv.Atoi(applArgsList[maxValueItem])
				errorMsg := fmt.Sprintf("Error: %v\n", err)
				if err != nil {
					log.Fatal(errorMsg)
				}

				applOptionsData.maxValue = maxValue
			} else {
				log.Fatal("Error: invalid parameter given.")
			}
		case MAXVALUE:
			nextItem := idx + 1
			if nextItem < applArgsListSize {
				maxValue, err := strconv.Atoi(applArgsList[nextItem])
				errorMsg := fmt.Sprintf("Error: %v\n", err)
				if err != nil {
					log.Fatal(errorMsg)
				}

				applOptionsData.maxValue = maxValue
				idx++ // Skip the next item
			} else {
				log.Fatal("Error: invalid parameter given.")
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

	fmt.Println("Exiting parseArgs...")

	return applOptionsData
}

func main() {
	// Starting string
	fmt.Println("Starting algebra application")

	// Create a slice that contains all parameters with the exception of application name
	applArgs := os.Args[1:]

	// Declare applOptionsData struct
	var applOptionsData applOptions

	if len(applArgs) == 0 {
		// Display usage message
		fmt.Println("Usage Error")

		// Exit the application
		log.Fatal("Usage: './main [<-factor>] [<-prime>] [<-generate> [{nbrOfGeneratedValues}] [{maxValue}]]'")
	} else {
		// parse args from command-line
		applOptionsData = parseArgs(applArgs)
		fmt.Println("Application options data: ", applOptionsData)
	}

	// Generate numbers requested by user input
	var nbrsList []int
	if applOptionsData.nbrsToGenerate > 0 {
		newList, err := mathutils.GenerateInts(applOptionsData.nbrsToGenerate, applOptionsData.maxValue)

		if err != nil {
			errorMsg := fmt.Sprintf("Error generating numbers, err - %v\n", err)
			log.Fatal(errorMsg)
		}

		// Updatge the numbers list
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
		// Add entries into the map
		factorsMap, err := algebra.BuildFactorsMap(nbrsList)

		if err != nil {
			errorMsg := fmt.Sprintf("Error processing Factors, err - %v\n", err)
			log.Fatal(errorMsg)
		}

		fmt.Println("The following are the factors for the numbers provided.")

		for key, factorsList := range factorsMap {
			fmt.Printf("Number: %v, factors: %v\n", key, factorsList)
		}
	}

	// Build primes map
	if applOptionsData.prime {
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

	// Ending string
	fmt.Println("Exiting Algebra application")
}
