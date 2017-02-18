package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hardTask(name string) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		time.Sleep(1 + time.Second)
		fmt.Printf("Hard Task %s...\n", name)
	}

	fmt.Printf("Hard Task %s DONE\n", name)
}

func easyTask() {
	time.Sleep(1 + time.Second)
	fmt.Println("========== EASYTASK ==========")
}

func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go hardTask(strconv.Itoa(i))
	}

	easyTask()

	wg.Wait()

	fmt.Println("DONE")
}
