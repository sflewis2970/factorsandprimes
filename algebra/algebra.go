package algebra

import (
	"errors"
	"fmt"

	commonutils "github.com/sflewis2970/factorsandprimes/common"
)

// Numbers ...Numbers gets input from the user from the user to build the number.
func Numbers() ([]int, error) {
	text, err := commonutils.GetConsoleInput()
	if err != nil {
		errorMsg := fmt.Sprintf("Error: %v\n", err)
		return nil, errors.New(errorMsg)
	}

	text = commonutils.RemoveCRLF(text)

	nbrListStr := commonutils.SplitTxt(text, " ")

	nbrsList, err := commonutils.BuildNbrList(nbrListStr)
	if err != nil {
		errorMsg := fmt.Sprintf("Error: %v\n", err)
		return nil, errors.New(errorMsg)
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
	// Create empty map
	factorsMap := make(map[int][]int)

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
		return nil, errors.New("error - factors map should not be empty")
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
		return nil, errors.New("error - primes use factors to build prime map, factor map should not be empty")
	}

	return primesMap, nil
}
