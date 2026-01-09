package miner

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Miner(ctx context.Context, wg *sync.WaitGroup, cnl chan<- int, num int){
	defer wg.Done()

	fmt.Printf("Шахтер %v начал работу\n", num)
	for{
		select{
		case <- ctx.Done():
			fmt.Printf("Шахтер номер %v ушел домой\n", num)
			return
		default:
			coal := rand.Intn(50)
			time.Sleep(1 * time.Second)
			fmt.Printf("Шахтер %v добыл %v угля\n", num, coal)
			cnl <- coal
		
		}
	}
}

func MinerPool(ctx context.Context, countMiners int) <-chan int{
	coalTransferPoint := make(chan int)
	wg := &sync.WaitGroup{}

	for i:= 1; i <= countMiners; i++{
		wg.Add(1)
		go Miner(ctx, wg, coalTransferPoint, i)
	}

	go func(){
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
