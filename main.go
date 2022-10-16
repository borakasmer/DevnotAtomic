package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type WareHouse struct {
	Name  string
	Count int
	Price float64
}

var (
	wh = WareHouse{"PS5", 5, 16000}
	wg = sync.WaitGroup{}
	mt = sync.Mutex{}
	at = atomic.Pointer[WareHouse]{}
)

func SellOneAtomic(name string, count int, price float64) {
	defer wg.Done()
	at.Swap(&WareHouse{name, count, price})
}
func SellOneMutex(wh *WareHouse) {
	defer wg.Done()
	mt.Lock()
	wh.Count--
	wh.Price += 1000
	mt.Unlock()
}
func main() {
	wg.Add(3)
	at.Store(&wh)
	go SellOneAtomic("PS5", 4, 17000)
	go SellOneAtomic("PS5", 4, 17000)
	go SellOneAtomic("PS5", 4, 17000)
	/*
		go SellOneMutex(&wh)
		go SellOneMutex(&wh)
		go SellOneMutex(&wh)
	*/
	wg.Wait()
	fmt.Println(at.Load())
	//fmt.Println(wh)
}
