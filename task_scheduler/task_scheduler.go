package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var currentTaskID = 0

// task structure
type task struct {
	id            int
	isSerial      bool
	function      func()
	queueTime     time.Time
	startTime     time.Time
	finishTime    time.Time
	sleepDuration int
}

var taskQueue []*task

func pushTask(f func(), isSerial bool) {
	t := &task{currentTaskID, isSerial, f, time.Now(), time.Now(), time.Now(), 0}
	currentTaskID++
	taskQueue = append(taskQueue, t)
}

// scheduler function
func schedule() {
	var serialQueue []*task
	var wg sync.WaitGroup
	for _, f := range taskQueue {
		if f.isSerial {
			// add serial tasks to a seperate queue to process later
			serialQueue = append(serialQueue, f)
		} else {
			// launch all concurrent tasks with goroutins
			wg.Add(1)
			go func() {
				defer wg.Done()
				f.startTime = time.Now()
				f.function()
				f.finishTime = time.Now()
			}()
		}
	}
	// run each serial tasks one by one
	for _, f := range serialQueue {
		f.startTime = time.Now()
		f.function()
		f.finishTime = time.Now()
	}
	wg.Wait() // wait until all concurrent tasks to finish
}

func main() {

	// function to be run as tasks
	testFunc := func() {
		rand.Seed(time.Now().UTC().UnixNano())
		r := rand.Intn(5)
		fmt.Printf("this function sleep %d seconds\n", r)
		time.Sleep(time.Duration(r) * time.Second)
	}

	// prepare all tasks, true=serial, false=concurrent
	pushTask(testFunc, true)
	pushTask(testFunc, false)
	pushTask(testFunc, false)
	pushTask(testFunc, false)
	pushTask(testFunc, true)
	pushTask(testFunc, true)
	pushTask(testFunc, true)
	pushTask(testFunc, false)
	pushTask(testFunc, true)
	pushTask(testFunc, false)

	// run scheduler
	schedule()

	// print
	var serialChar byte
	for _, f := range taskQueue {

		if f.isSerial {
			serialChar = 'S'
		} else {
			serialChar = 'C'
		}
		fmt.Printf("Task ID=%d, %c, enqueue %s, start %s, end %s\n", f.id, serialChar,
			f.queueTime.Format("2006-01-02 15:04:05"),
			f.startTime.Format("2006-01-02 15:04:05"),
			f.finishTime.Format("2006-01-02 15:04:05"))
	}
}
