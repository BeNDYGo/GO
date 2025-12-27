package main

import (
	"fmt"
)

func main(){

	type Word struct{
		Text string
		Leght int
	}

	// массив от 65 до 90 с большими буквами
	upcrar := [...]int{}
	c := 0
	for i := 65; i< 91; i++{
		upcrar[c] = i
		c += 1
	}
}

func (word *Word) Normalize(srt string){


}