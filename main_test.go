package main

import "testing"

func BenchmarkSellOneMutex(b *testing.B) {
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go SellOneMutex(&wh)
	}
	wg.Wait()
}
func BenchmarkSellOneAtomic(b *testing.B) {
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go SellOneAtomic("PS5", 100000-i, float64(16000+i))
	}
	wg.Wait()
}
