package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Content struct {
	rw  sync.RWMutex
	val int
}

func (c *Content) Read() int {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.val
}
func (c *Content) Write(v int) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.val = v
}

func main() {
	const (
		readerNum = 100
		writerNum = 3
	)
	content := new(Content)
	var wg sync.WaitGroup
	for i := 0; i < writerNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			content.Write(rand.Intn(10))
		}()
	}
	for i := 0; i < readerNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(content.Read())
		}()
	}

}
