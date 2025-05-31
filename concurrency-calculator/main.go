package main

import (
	"fmt"
	"sync"

	"concurrency-calculator/internal/calculator"
)

func main() {
	var wg sync.WaitGroup
	calc := &calculator.Calculator{Result: 1}

	operations := []func(){
		func() { calc.Add(10) },
		func() { calc.Subtract(3) },
		func() { calc.Multiply(5) },
		func() { calc.Add(7) },
		func() { calc.Multiply(2) },
	}

	wg.Add(len(operations))
	for _, op := range operations {
		go func(f func()) {
			defer wg.Done()
			f()
		}(op)
	}

	wg.Wait()
	fmt.Println("Final Result:", calc.Result)
}
