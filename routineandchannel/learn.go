package routineandchannel

import (
	"fmt"
	"time"
)

func Learn() {

	workerCount := 10

	// 1. Use channel to receive data from go routine
	//    Send data from go routine back to outside
	responseChannel := make(chan string)
	for i := 0; i < workerCount; i++ {
		workerID := fmt.Sprintf("worker-%d", i)
		go worker1(workerID, responseChannel)
	}

	for i := 0; i < workerCount; i++ {
		res := <-responseChannel
		println(res)
	}

	close(responseChannel)
	println("All response returned")

	// 2. Use channel to signal go module to exit
	//    Send data from outside go model into module
	exitChannel := make(chan bool)
	for i := 0; i < workerCount; i++ {
		workerID := fmt.Sprintf("worker-%d", i)
		go worker3(workerID, exitChannel)
	}

	time.Sleep(10 * time.Second)
	for i := 0; i < workerCount; i++ {
		exitChannel <- true
	}
	close(exitChannel)
	time.Sleep(2 * time.Second)
	println("Main is exited")
}

func worker1(workerID string, responseChannel chan string) {
	// Simulate request latency
	time.Sleep(1 * time.Second)
	responseChannel <- (workerID + " Response")
}

func worker3(workerID string, exitChannel chan bool) {
	i := 0
	for true {
		i++

		select {
		case <-exitChannel:
			println(fmt.Sprintf("Worker=%s has exited", workerID))
			return
		default:
			println(fmt.Sprintf("Worker=%s, Counter=%d", workerID, i))
			time.Sleep(1 * time.Second)
		}
	}
}