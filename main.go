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

func GoDrive(auto Auto, length int) {
	auto.Drive(length)
}

func main(){
	bmw := &BMW{Fuel: 500}
	tesla := &Tesla{Pover: 300}

	GoDrive(bmw, 130)
	GoDrive(tesla, 170)
}
