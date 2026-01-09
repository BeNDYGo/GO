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

	fmt.Printf("Шахтер %v начал работу", num)
	for{
		select{
		case <- ctx.Done():
			fmt.Printf("Шахтер номер %v ушел домой", num)
			return
		default:
			coal := rand.Intn(50)
			time.Sleep(1 * time.Second)
			fmt.Printf("Шахтер %v добыл %v угля", num, coal)
			cnl <- coal
		
		}
	}
}

func StartMine(ctx context.Context, countMiners int) <-chan int{
	coalTrinsferPoint := make(chan int)
	wg := &sync.WaitGroup{}

	for i:= 1; i <= countMiners; i++{
		wg.Add(1)
		go Miner(ctx, wg, coalTrinsferPoint, i)
	}

	go func(){
		wg.Wait()
		close(coalTrinsferPoint)
	}()

	return coalTrinsferPoint
}
