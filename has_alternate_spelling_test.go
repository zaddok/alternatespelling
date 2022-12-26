package alternatespelling

import (
	"fmt"
	"testing"
)

func TestHasAlternateSpelling(t *testing.T) {

	if HasAlternateSpelling("agonise") != "agonize" {
		fmt.Printf("Spelling translate failed\n")
	}

	if HasAlternateSpelling("agonize") != "agonise" {
		fmt.Printf("Spelling translate failed\n")
	}

	if HasAlternateSpelling("fish") != "" {
		fmt.Printf("Spelling translate failed\n")
	}

}
