// The hugeparam command identifies by-value parameters that are larger than n bytes.
//
// Example:
//	$ ./hugeparams encoding/xml
package main

import "fmt"

func main() {
	x := "hello"
	fmt.Println(x)

	_ = func() {
		x := "shadowed"
		fmt.Println(x)
	}
}
