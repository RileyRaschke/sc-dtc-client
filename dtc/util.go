package dtc

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
