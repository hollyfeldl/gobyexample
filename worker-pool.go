// https://gobyexample.com/worker-pools
// some play to make the work variable by id to see that workers will do 
// different work.

package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second * time.Duration(id))
		results <- j * 2
	}

}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)
	queue := 12

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= queue; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= queue; a++ {
		<-results
	}
}