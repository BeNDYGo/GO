package main

import (
	"fmt"
	"sync"
)

var c int = 0

func f(wg *sync.WaitGroup){
	defer wg.Done()
	for i := 1 ;i <= 1000; i++{
		c++
	}
}

func main(){
	wg := &sync.WaitGroup{}

	wg.Add(6)

	go f(wg)
	go f(wg)
	go f(wg)
	go f(wg)
	go f(wg)
	go f(wg)

	wg.Wait()

	fmt.Printf("Значение i: %v ", c)
}
