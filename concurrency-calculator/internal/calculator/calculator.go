package calculator

import (
	"sync"
)

type Calculator struct {
	Result int
	mu     sync.Mutex
}

func (c *Calculator) Add(value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Result += value
}

func (c *Calculator) Subtract(value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Result -= value
}

func (c *Calculator) Multiply(value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Result *= value
}
