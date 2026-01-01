package main

import (
	"errors"
	"fmt"
)
type User struct{
	Name string
	Ballance int
}

func Pay (user *User, usd int) error{
	if user.Ballance - usd >= 0 {
		user.Ballance -= usd
		return nil
	}
	return errors.New("недостаточно средстав")
}

func main(){
	user1 := User{
		Name: "Олег",
		Ballance: 1000,
	}

	pay := Pay(&user1, 150)

	if pay == nil {
		fmt.Printf("Баланс %s: %b\n", user1.Name, user1.Ballance)
	}else {
		fmt.Println("недостаточно средств")
	}

}
