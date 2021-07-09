package main

import "fmt"

// gcd
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// http server


func main() {
	fmt.Println(gcd(20, 8))
}