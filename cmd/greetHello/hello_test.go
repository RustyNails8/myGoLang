// Medium https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318
package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	// Test for empty argument
	emptyResult := hello("")
	// Output
	// "hello Dude !!"

	if emptyResult != "hello Dude !!" {
		t.Errorf("hello(\"\") failed, expected %v , got %v", "hello Dude !!", emptyResult)
	}

	// Test for real argumemt
	result := hello("Sumit")
	// Output
	// "hello Sumit !!"

	if result != "hello Sumit !!" {
		t.Errorf("hello(\"Sumit\") failed, expected %v , got %v", "hello Sumit !!", result)
	}

}
