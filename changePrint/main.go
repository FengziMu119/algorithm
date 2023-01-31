package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	oddChan := make(chan struct{})
	evenChan := make(chan struct{})
	waitGroup.Add(2)
	go func() {
		defer waitGroup.Done()
		for i := 1; i <= 100; i += 2 {
			<-oddChan
			fmt.Println()
			evenChan <- struct{}{}
		}
		<-oddChan
	}()

	go func() {
		defer waitGroup.Done()
		for i := 2; i <= 100; i += 2 {
			<-evenChan
			fmt.Println()
			oddChan <- struct{}{}
		}
	}()
	oddChan <- struct{}{}
	waitGroup.Wait()
}
