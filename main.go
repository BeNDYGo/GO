package main

import (
	"Main/miner"
	"Main/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var coal atomic.Int64
	var mails []string
	mtx := sync.Mutex{}

	minerContext, minerClose := context.WithCancel(context.Background())
	postmanContext, postmanClose := context.WithCancel(context.Background())

	go func(){
		time.Sleep(6 * time.Second)
		minerClose()
	}()

	go func(){
		time.Sleep(3 * time.Second)
		postmanClose()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 30)
	mailTransferpoint := postman.PostmanPool(postmanContext, 1)
	
	wg := &sync.WaitGroup{}
	
	wg.Add(1)
	go func(){
		defer wg.Done()

		for v := range coalTransferPoint{
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func(){
		defer wg.Done()

		for v := range mailTransferpoint{
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()
	
	wg.Wait()

	fmt.Println("Добыто угля: ", coal.Load())
	mtx.Lock()
	fmt.Println("Доставлено писем: ", len(mails))
	mtx.Unlock()
}
