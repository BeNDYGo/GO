package main

import (
	"fmt"
)

type Auto interface{
	Drive(length int)
}

type BMW struct{
		Year  int
		Color string
		Fuel int
}

type Tesla struct{
		Year  int
		Color string
		Pover int
}

func (b *BMW) Drive(length int){
	b.Fuel -= length
	fmt.Println("BMW едет на расстояние", length, "km, осталось топлива:", b.Fuel, "л")
}

func (t *Tesla) Drive(length int){
	t.Pover -= length
	fmt.Println("Tesla едет на расстояние", length, "km, осталось заряда:", t.Pover, "киловатт-час")
}

func main(){
	autos := []Auto{
		&BMW{Fuel: 500},
		&Tesla{Pover: 300},
	}

	lengths := []int{130, 170}

	for i, auto := range autos {
		fmt.Printf("Автомобиль %d: ", i)
		auto.Drive(lengths[i])
	}
}
