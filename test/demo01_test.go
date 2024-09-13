package test

import (
	"fmt"
	"testing"
)

func PrintNumbers(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

func TestDemo01(t *testing.T) {
	PrintNumbers(1, 2, 3, 4)
	nums := []int{11, 22, 33, 44}
	PrintNumbers(nums...)
}
