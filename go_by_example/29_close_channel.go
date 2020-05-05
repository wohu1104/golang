package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, flag := <-jobs
			if flag {
				fmt.Println("received ", j)
			} else {
				fmt.Println("received  all jobs", j)
				done <- true
				return
			}
		}
	}()

	for i := 0; i <= 3; i++ {
		jobs <- i
		fmt.Println("send jobs", i)
	}

	close(jobs)
	fmt.Println("send all jobs")
	<-done
}
