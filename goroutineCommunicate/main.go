package main

import (
	"fmt"
	"time"
)

func Worker(workerIndex int, jobs chan int, result chan int) {

	for jobIndex := range jobs {
		fmt.Println("Worker", workerIndex, " has started job", jobIndex)
		fmt.Println("Worker is doing job....")
		time.Sleep(1 * time.Second)
		fmt.Println("Worker", workerIndex, " has finished job", jobIndex)
		result <- jobIndex * 2
	}

}
func API(numJobs int) int {
	jobs := make(chan int, numJobs)
	defer close(jobs)

	result := make(chan int, numJobs)
	defer close(result)
	workNums := 10
	//处理工作
	for workIndex := 0; workIndex < workNums; workIndex++ {
		go Worker(workIndex, jobs, result)
	}
	//工作进入
	for jobIdx := 0; jobIdx < numJobs; jobIdx++ {
		jobs <- jobIdx
	}
	//读取结果
	sum := 0
	for jobIdx := 0; jobIdx < numJobs; jobIdx++ {
		select {
		case temp := <-result:
			fmt.Println(temp, "is pushed in result channel ")
			sum = sum + temp

		}

	}

	return sum
}
func main() {
	fmt.Println("API excute result is ", API(5))
	fmt.Println("API excute result is ", API(10))
	fmt.Println("API excute result is ", API(15))

}
