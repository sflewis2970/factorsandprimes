package mathutils_test

import (
	"fmt"
	"testing"

	"github.com/sflewis2970/factorsandprimes/mathutils"
)

func TestFindNumber(t *testing.T) {
	t.Log("Start TestFindNumber")

	// Declare new number list
	nbrList := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// Find number
	findNbr := 60

	// Search for generated number
	gotFound := mathutils.FindNumber(nbrList, findNbr)

	logMsg := ""
	if !gotFound {
		logMsg = fmt.Sprintf("Could not find number %d in: %v\n", findNbr, nbrList)
		t.Errorf(logMsg)
	} else {
		logMsg = fmt.Sprintf("Found %d in: %v\n", findNbr, nbrList)
		t.Log(logMsg)
	}

	// Find number (outside)
	findNbr = 150

	// Search for generated number
	gotFound = mathutils.FindNumber(nbrList, findNbr)

	if gotFound {
		logMsg = fmt.Sprintf("Found number %d in: %v\n", findNbr, nbrList)
		t.Errorf(logMsg)
	} else {
		logMsg = fmt.Sprintf("Could not find number %d in: %v\n", findNbr, nbrList)
		t.Log(logMsg)
	}

	t.Log("Stop TestFindNumber")
}

func TestGenerateInts(t *testing.T) {
	t.Log("Start TestGenerateInts")

	// Set numbers to generate
	nbrsToGenerate := 10

	// Set maximum value
	maxValue := 50

	// Generate numbers and place the result in a list
	got, err := mathutils.GenerateInts(nbrsToGenerate, maxValue)
	if err != nil {
		logMsg := fmt.Sprintf("Error: %v\n", err)
		t.Errorf(logMsg)
	}

	if len(got) == 0 {
		t.Errorf("Error: list cannot be empty")
	}

	// Print generated numbers list
	logMsg := fmt.Sprintf("got: %v\n", got)
	t.Log(logMsg)

	t.Log("Stop TestGenerateInts")
}
