package algebra

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Numbers ...Numbers gets input from the user from the user to build the number.
func Numbers() ([]int, error) {
	// Declare factorsList
	var nbrsList []int

	// Setup new reader object
	reader := bufio.NewReader(os.Stdin)

	// Get input from user
	fmt.Print("Enter input values (seperated by spaces): ")
	text, err := reader.ReadString('\n')

	if err != nil {
		errorMsg := fmt.Sprintf("Error: %v\n", err)
		return nil, errors.New(errorMsg)
	}

	// Remove carriage return and line feed
	text = strings.Replace(text, "\r\n", "", -1)

	// Build a slice of the entered values
	strSlice := strings.Split(text, " ")

	// Convert each value in the slice to an int
	for _, strSliceVal := range strSlice {
		// Convert string to number
		sliceNbr, err := strconv.Atoi(strSliceVal)

		if err != nil {
			errorMsg := fmt.Sprintf("Error: %v\n", err)
			return nil, errors.New(errorMsg)
		}

		// Add number to nbrs list
		nbrsList = append(nbrsList, sliceNbr)
	}

	return nbrsList, nil
}

// BuildFactorsList ...BuildFactorsList takes a number and returns the list of numberic factors.
func BuildFactorsList(nbrVal int) []int {
	// Declare factorsList
	var factorsList []int

	for factorVal := 1; factorVal <= nbrVal; factorVal++ {

		// If factor found, add to factorsList
		if (nbrVal % factorVal) == 0 {
			factorsList = append(factorsList, factorVal)
		}
	}

	return factorsList
}

// BuildFactorsMap ...BuildFactorsMap takes a list of numbers and returns a map of numbers of factors
func BuildFactorsMap(nbrsList []int) (map[int][]int, error) {
	// Declare map
	var factorsMap map[int][]int

	// Create empty map
	factorsMap = make(map[int][]int)

	// Declare factorsList
	var factorsList []int

	// Create factor map
	for _, nbrsVal := range nbrsList {
		// Build factors for each number
		factorsList = BuildFactorsList(nbrsVal)

		// Add factor list to map
		factorsMap[nbrsVal] = factorsList
	}

	if len(factorsMap) == 0 {
		return nil, errors.New("Error: factors map should not be empty")
	}

	return factorsMap, nil
}

// BuildPrimesMap ...BuildPrimesMap takes a list of numbes and rturns a map of numbers of primes
func BuildPrimesMap(nbrsList []int) (map[int][]int, error) {
	// Build factors from the numbers list
	factorsMap, err := BuildFactorsMap(nbrsList)

	if err != nil {
		errorMsg := fmt.Sprintf("Error: %v\n", err)
		return nil, errors.New(errorMsg)
	}

	// Declare map
	var primesMap map[int][]int

	// Build primes map
	if len(factorsMap) > 0 {
		// Create empty map
		primesMap = make(map[int][]int)

		// Create primes map
		for key, factorsList := range factorsMap {
			if len(factorsList) == 2 {
				primesMap[key] = factorsList
			}
		}
	} else {
		errorMsg := fmt.Sprintf("Error: primes use factors to build prime map, factor map should not be empty\n")
		return nil, errors.New(errorMsg)
	}

	return primesMap, nil
}
