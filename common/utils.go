package commonutils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetConsoleInput() (string, error) {
	// Setup new reader object
	reader := bufio.NewReader(os.Stdin)

	// Get input from user
	fmt.Print("Enter input values (seperated by spaces): ")
	text, err := reader.ReadString('\n')

	if err != nil {
		errorMsg := fmt.Sprintf("%v\n", err)
		return "", errors.New(errorMsg)
	}

	return text, nil
}

func RemoveCRLF(text string) string {
	// Remove carriage return and line feed
	return strings.Replace(text, "\r\n", "", -1)
}

func SplitTxt(text string, delimiter string) []string {
	// Build a slice of the entered values
	return strings.Split(text, delimiter)
}

func BuildNbrList(nbrListStr []string) ([]int, error) {
	var nbrsList []int

	// Convert each value in the slice to an int
	for _, strVal := range nbrListStr {
		// Convert string to number
		nbr, err := strconv.Atoi(strVal)

		if err != nil {
			errorMsg := fmt.Sprintf("%v\n", err)
			return nil, errors.New(errorMsg)
		}

		// Add number to nbrs list
		nbrsList = append(nbrsList, nbr)
	}

	return nbrsList, nil
}
