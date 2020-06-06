//Package formatmore shows how DCPvlan Sheet can be formatted
package main

import (
	"fmt"
)

func main() {

	fmt.Printf("%8s | %12s | %16s | %6s | %10s\n", "Cluster", "LPAR", "IP", "Slice", "Service")
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("%8s | %12s | %16s | %6s | %10s\n", "legdgn01", "dcpaix3456", "10.141.65.12", "175", "AGH_58_CI")
	fmt.Printf("%8s | %12s | %16s | %6s | %10s\n", "legdgn02", "dcpaix4567", "10.141.66.121", "80", "DGH_21_DB")
	fmt.Printf("%8s | %12s | %16s | %6s | %10s\n", "legdgi01", "dcpaix45678", "10.141.67.14", "30", "PGH_48_HD")
	fmt.Printf("%8s | %12s | %16s | %6s | %10s\n", "legpgn01", "dcpaix23546", "10.141.66.377", "187", "VGH_56_JI")
	fmt.Printf("%8s | %12s | %16s | %6s | %10s\n", "legpgi01", "dcpaix64576", "10.141.60.25", "475", "JGH_68_SY")

}
