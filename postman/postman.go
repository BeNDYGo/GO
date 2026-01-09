package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Postman(ctx context.Context, wg *sync.WaitGroup, cnl chan<- string, num int, mail string) {
	defer wg.Done()
	for {
		select{
		case <- ctx.Done():
			fmt.Printf("Почтальон %v закончил рабочий день\n", num)
			return
		default:
			fmt.Printf("Почтальон номер %v взял письмо\n", num)
			time.Sleep(1 * time.Second)

			cnl <- mail

			fmt.Printf("Почтальон номер %v донес письмо до почты %v\n", num, mail)
		}
	}
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string{
	mailTransferpoint := make(chan string)
	wg := &sync.WaitGroup{}

	for i:= 1; i <= postmanCount; i++{
		wg.Add(1)
		go Postman(ctx, wg, mailTransferpoint, i, posmanToMail(i))
	}

	go func(){
		wg.Wait()
		close(mailTransferpoint)
	}()
	
	return mailTransferpoint
}

func posmanToMail(postmanNumber int) string{
	ptm := map[int]string{
		1: "Важное письмо номер 1",
		2: "Важное письмо номер 2",
		3: "Важное письмо номер 3",
	}

	mail, ok := ptm[postmanNumber]
	if !ok{
		return "Реклама"
	}
	return mail

}