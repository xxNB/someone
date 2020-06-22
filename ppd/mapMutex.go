package main 

import (
	"math/rand"
	"sync"
)

const N = 10

func main(){
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i:=0;i<N;i++{
		go func(){
			defer wg.Done()
			mu.Lock()
			m[rand.Int()] = rand.Int()
			mu.Unlock()
		}()
	}
	wg.Wait()
	println(len(m))
}