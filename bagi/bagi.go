package bagi

import "fmt"

func Bagi(a, b int) int {
	if b == 0 {
		fmt.Println("Error: pembagian dengan nol.")
		return 0
	}
	return a / b
}