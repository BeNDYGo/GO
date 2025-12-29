package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	fmt.Print("Введите слова: ")
	input := bufio.NewScanner(os.Stdin)

	input.Scan()

	result := strings.Fields(input.Text())
	for i, word := range result {
		fmt.Println(i, "-", word)
	}
}