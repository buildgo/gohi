package util

import "fmt"

func Describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
