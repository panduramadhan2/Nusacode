package main

import (
	"Tugas-Func/add"
	"Tugas-Func/bagi"
	"Tugas-Func/kurang"
	"fmt"
)

func main() {
	fmt.Println("Add(10,5):", add.Add(10, 5))
	fmt.Println("Bagi(20, 4):", bagi.Bagi(20, 4))
	fmt.Println("KurangVariadic(100, 10, 5, 2, 1):", kurang.KurangVariadic(100, 10, 5, 2, 1))
}
