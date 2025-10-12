package main

import (
	"fmt"
	"sync"
)

func task(i int,w *sync.WaitGroup) {
	//defer is used so that after the function is executed then only it decerements the counter 
	defer w.Done()
	fmt.Println("Doing task", i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		//works like synchronisation tool
		//increments pointer
		wg.Add(1)
		go task(i,&wg) 
	}
	//waits till all the iterations are done
	wg.Wait()
}
