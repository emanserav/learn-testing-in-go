package external_test

import (
	"fmt"

	"external"
)

func ExampleMultiply() {
	res := external.Multiply(1, 2)
	fmt.Println(res)
	// Output: 2
}
