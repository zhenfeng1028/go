// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import (
	"fmt"
	"time"
)

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	start := time.Now()
	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have
	// finished. An alternative way to wait for multiple
	// goroutines is to use a [WaitGroup](waitgroups).
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	end := time.Now()
	fmt.Println("time elapsed:", end.Sub(start))
}

// Our running program shows the 5 jobs being executed by
// various workers. The program only takes about 2 seconds
// despite doing about 5 seconds of total work because
// there are 3 workers operating concurrently.
/*
	worker 1 started  job 2
	worker 2 started  job 3
	worker 3 started  job 1
	worker 3 finished job 1
	worker 3 started  job 4
	worker 2 finished job 3
	worker 2 started  job 5
	worker 1 finished job 2
	worker 2 finished job 5
	worker 3 finished job 4
	time elapsed: 2.001251144s
*/
