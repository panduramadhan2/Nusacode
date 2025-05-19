package kurang

import "fmt"

func KurangVariadic(numbers ...int) int {
	if len(numbers) != 5 {
		fmt.Println("Harus memberikan tepat 5 angka.")
		return 0
	}
	result := numbers[0]
	for i := 1; i < 5; i++ {
		result -= numbers[i]
	}
	return result
}
