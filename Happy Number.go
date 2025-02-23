package main

import (
	"fmt"
	"strconv"
	"strings"
)

func converting(n int) int {

	if n == 1 || n == 0 {
		return n
	}

	convertedS := strconv.Itoa(n)

	if len(convertedS) == 1 {
		return n
	}

	listString := strings.Split(convertedS, "")
	res := 0

	for _, number := range listString {
		numberConverted, _ := strconv.Atoi(number)
		res += numberConverted * numberConverted
	}
	return converting(res)
}

func isHappy(n int) bool {
	result := converting(n)

	if result == 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isHappy(7))
}
