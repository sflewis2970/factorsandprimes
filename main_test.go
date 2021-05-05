package main_test

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	t.Log("Starting test")

	name := "test"

	msg := fmt.Sprintf("Testing %v\n", name)
	t.Log(msg)

	t.Log("Ending test")
}
