package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var bank int = 1000

func pay(w http.ResponseWriter, r *http.Request){
	body, err := io.ReadAll(r.Body)
	if err != nil{
		fmt.Println(err)
		return
	}

	bodystr := string(body)
	sum, _ := strconv.Atoi(bodystr)
	bank -= sum
	fmt.Println("Совершена покупка на", sum, "$")
	fmt.Println("Остаток:", bank, "$")
}
func main(){
	http.HandleFunc("/pay", pay)

	err := http.ListenAndServe(":9091", nil)
	if err != nil{
		fmt.Println(err)
	}
}