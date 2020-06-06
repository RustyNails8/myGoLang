// Medium https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318
package main

import (
	"testing"
)

// greethello Example
func TestMain(t *testing.T) {

	// Test for empty argument
	emptyGreet := hello("")
	// Output
	// "hello Dude !!"

	if emptyGreet != "hello Dude !!" {
		t.Errorf("hello(\"\") failed, expected %v , got %v", "hello Dude !!", emptyGreet)
	}

	// Test for real argumemt
	result := hello("Sumit")
	// Output
	// "hello Sumit !!"

	if result != "hello Sumit !!" {
		t.Errorf("hello(\"Sumit\") failed, expected %v , got %v", "hello Sumit !!", result)
	}

}
