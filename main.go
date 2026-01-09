package main

import (
	"fmt"
	"sync"
)

var c int = 0
var mtx sync.RWMutex

func f(wg *sync.WaitGroup){
	defer wg.Done()
	for i := 1 ;i <= 1000; i++{
		mtx.Lock()
		c++
		mtx.Unlock()

	}
}

func main(){
	wg := &sync.WaitGroup{}

	wg.Add(3)
	go f(wg)
	go f(wg)
	go f(wg)

	wg.Wait()

	fmt.Printf("Значение i: %v ", c)
}
