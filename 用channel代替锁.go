package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var mu sync.Mutex
var ch1 =make(chan int,1)

func add() {

	for i := 0; i < 50000; i++ {
		<-ch1
		x = x + 1
		ch1<-0
		fmt.Println(x)
	}
	wg.Done()
}


func main() {
	wg.Add(2)
	ch1<-9
	go add()
	go add()
    wg.Wait()
	fmt.Println(x)
}
