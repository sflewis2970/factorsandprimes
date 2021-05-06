package mathutils

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// enterNumbersToGenerate ...Enter how many numbers will be geberated by the system
func enterNumbersToGenerate() (int, error) {
	// Ask the user to enter the number of numbers to generate value
	// Setup new reader object
	reader := bufio.NewReader(os.Stdin)

	// Get input from user
	fmt.Printf("How many numbers would you like generated? ")
	text, err := reader.ReadString('\n')

	if err != nil {
		errorMsg := fmt.Sprintf("Error: %v\n", err)
		return 0, errors.New(errorMsg)
	}

	// Remove carriage return and line feed
	text = strings.Replace(text, "\r\n", "", -1)

	// Convert string to number
	nbrsToGenerate, err := strconv.Atoi(text)

	if err != nil {
		errorMsg := fmt.Sprintf("Error: %v\n", err)
		return 0, errors.New(errorMsg)
	}

	return nbrsToGenerate, nil
}

// enterMaximumValue ...Enter the maximum value for each number
func enterMaximumValue(nbrsToGenerate int) (int, error) {
	// Setup new reader object
	reader := bufio.NewReader(os.Stdin)

	// Get input from user
	fmt.Printf("The system will generate %v numbers.\n", nbrsToGenerate)
	fmt.Printf("Enter the maximum number for each number: ")
	text, err := reader.ReadString('\n')

	if err != nil {
		errorMsg := fmt.Sprintf("Error: %v\n", err)
		return 0, errors.New(errorMsg)
	}

	// Remove carriage return and line feed
	text = strings.Replace(text, "\r\n", "", -1)

	// Convert string to number
	maxValue, err := strconv.Atoi(text)

	if err != nil {
		errorMsg := fmt.Sprintf("Error: %v\n", err)
		return 0, errors.New(errorMsg)
	}

	return maxValue, nil
}

// GenerateInt ...GenerateInt produces a single number
func generateInt(maxValue int) int {
	// Generate new number from randomizer
	return rand.Intn(maxValue) + 1
}

// FindNumber ...FindNumber takes a list of numbers and a number to search for and returns true if the number is found and false if not found.
func FindNumber(nbrsList []int, nbrToFind int) bool {

	// Create primes map
	for _, nbrsVal := range nbrsList {
		if nbrsVal == nbrToFind {
			return true
		}
	}

	return false
}

// GenerateInts ...Generate numbers builds a list of numbers
func GenerateInts(nbrsToGenerate, maxValue int) ([]int, error) {
	if nbrsToGenerate <= 0 {
		enterNumbersToGenerate()
	}

	// Ask the user to enter the maximum value
	if maxValue <= 0 {
		enterMaximumValue(nbrsToGenerate)
	}

	// Declare nbrsList
	var nbrsList []int

	// Seed random nmber ganerator
	rand.Seed(time.Now().UnixNano())

	// Generate integer
	for idx := 0; idx < nbrsToGenerate; idx++ {
		newNbr := generateInt(maxValue)

		// Make sure number is not in the list
		FindNumber(nbrsList, newNbr)

		// Don't include 1 in the list of available numbers
		if newNbr == 1 {
			newNbr++
		}

		// Append number to slice
		nbrsList = append(nbrsList, newNbr)
	}

	return nbrsList, nil
}
