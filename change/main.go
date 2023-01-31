package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {

	fishChan := make(chan struct{}, 1)
	dogChan := make(chan struct{}, 1)
	catChan := make(chan struct{}, 1)

	go printCat(fishChan, catChan)
	go printDog(catChan, dogChan)
	go printFish(dogChan, fishChan)
	fishChan <- struct{}{}

	waitGroup.Add(3)
	waitGroup.Wait()

}

func printCat(fishChan, catChan chan struct{}) {
	defer waitGroup.Done()
	defer close(catChan)
	for i := 0; i < 10; i++ {
		<-fishChan
		fmt.Println("cat....")
		catChan <- struct{}{}
	}
}

func printDog(catChan, dogChan chan struct{}) {
	defer waitGroup.Done()
	defer close(dogChan)
	for i := 0; i < 10; i++ {
		<-catChan
		fmt.Println("dog....")
		dogChan <- struct{}{}
	}
}

func printFish(dogChan, fishChan chan struct{}) {
	defer waitGroup.Done()
	defer close(fishChan)
	for i := 0; i < 10; i++ {
		<-dogChan
		fmt.Println("fish...")
		fishChan <- struct{}{}
	}
}
