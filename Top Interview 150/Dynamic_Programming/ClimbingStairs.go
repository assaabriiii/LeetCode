package main

import "fmt"


func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 0 {
        return 0
    }
    n1, n2 := 1,2
    for i:=2 ; i<n ; i++ {
        n1, n2 = n2, n1+n2
    }
	return n2
}

func main() {
    fmt.Println(climbStairs(4))
}