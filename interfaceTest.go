package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 99865
	fmt.Println(findNthDigit(n))
}

func findNthDigit(n int) int {
	dis := 1
	count := 9
	start := 1
	for n > count {
		n -= count
		dis++
		start *= 10
		count = dis * start * 9
		fmt.Println(count)
	}
	num := start + (n-1)/dis
	fmt.Println(num)
	return int(strconv.Itoa(num)[(n-1)%dis] - '0')
}