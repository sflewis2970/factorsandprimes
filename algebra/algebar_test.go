package algebra_test

import (
	"fmt"
	"testing"

	"github.com/sflewis2970/factorsandprimes/algebra"
	"github.com/sflewis2970/factorsandprimes/mathutils"
)

func TestBuildFactorsList(t *testing.T) {
	t.Log("Start TestBuildFactorsList")

	// Set maximum value
	nbrsToGenerate := 10

	// Set maximum value
	maxValue := 50

	// Generate
	gotNbrList, err := mathutils.GenerateInts(nbrsToGenerate, maxValue)

	if err != nil {
		logMsg := fmt.Sprintf("Error - %v\n", err)
		t.Errorf(logMsg)
	} else {
		// Check to make sure that the number of items created is correct
		gotNbrListSize := len(gotNbrList)
		if gotNbrListSize != nbrsToGenerate {
			logMsg := fmt.Sprintf("Error - invalid gotNbrList size %v\n", gotNbrListSize)
			t.Errorf(logMsg)
		}
	}

	// Deslare number list
	var gotFactorList []int

	// Build factors for each number
	for _, gotNbr := range gotNbrList {
		// Check to make sure that each number is a valid number
		if gotNbr < 0 || gotNbr > maxValue {
			logMsg := fmt.Sprintf("Error - gotNbr out of range: %v\n", gotNbr)
			t.Errorf(logMsg)
		}

		gotFactorList = algebra.BuildFactorsList(gotNbr)

		// Each number generated should at least 2 factors
		if len(gotFactorList) == 0 {
			logMsg := fmt.Sprintf("Error: gotNbr %d has no factors\n", gotNbr)
			t.Errorf(logMsg)
		}

		logMsg := fmt.Sprintf("%d has %d factors\n", gotNbr, gotFactorList)
		t.Log(logMsg)
	}

	t.Log("Stop TestBuildFactorsList")
}

func TestBuildFactorsMap(t *testing.T) {
	t.Log("Start TestBuildFactorsMap")

	// Set numbers to generate
	nbrsToGenerate := 10

	// Set maximum value
	maxValue := 50

	// Generate
	newNbrList, err := mathutils.GenerateInts(nbrsToGenerate, maxValue)

	if err != nil {
		logMsg := fmt.Sprintf("Error - %v\n", err)
		t.Errorf(logMsg)
	}

	// Declare list
	var gotList []int

	for newNbr := range newNbrList {
		gotList = algebra.BuildFactorsList(newNbr)
	}

	if len(gotList) == 0 {
		t.Errorf("Error list cannot be empty")
	}

	// Output the factors generated
	logMsg := fmt.Sprintf("factors list: %v\n", gotList)
	t.Log(logMsg)

	t.Log("Stop BuildFactorsMap")
}

func TestBuilPrimesMap(t *testing.T) {
	t.Log("Start BuildPrimesMap")

	// Set maximum value
	nbrsToGenerate := 10

	// Set maximum value
	maxValue := 100

	// Generate
	nbrList, err := mathutils.GenerateInts(nbrsToGenerate, maxValue)

	if err != nil {
		logMsg := fmt.Sprintf("Error = %v\n", err)
		t.Errorf(logMsg)
	}

	// Generate
	gotFactorMap, err := algebra.BuildFactorsMap(nbrList)

	if err != nil {
		logMsg := fmt.Sprintf("Error = %v\n", err)
		t.Errorf(logMsg)
	}

	if len(gotFactorMap) == 0 {
		t.Errorf("Error map cannot be empty")
	}

	t.Log("Stop BuildPrimesMap")
}
