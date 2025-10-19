package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func main() {
	// 创建一堆任务
	tasks := make([]Task, 5)
	for i := 0; i < 5; i++ {
		tasks[i] = createTask(i)
	}

	// 执行任务调度
	results := scheduleTask(tasks)

	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("任务 %d 执行失败 %v\n", result.TaskId, result.Error)
		} else {
			fmt.Printf("任务 %d 执行成功,耗时: %v\n", result.TaskId, result.duration)
		}
	}
}

type TaskResult struct {
	TaskId   int
	duration time.Duration
	Error    error
}

// 定义一个任务函数类型
type Task func() error

// 创建任务的方法
func createTask(TaskId int) Task {
	return func() error {
		// 随机耗时 100-500毫秒
		duration := time.Duration(100+rand.Intn(401)) * time.Millisecond
		time.Sleep(duration)

		// 一定的出错概率
		if rand.Float32() < 0.1 {
			return fmt.Errorf("执行 %d 任务失败\n", TaskId)
		}

		//fmt.Printf("执行 %d 任务成功\n", TaskId)
		return nil
	}
}

func scheduleTask(tasks []Task) []TaskResult {
	var wg sync.WaitGroup
	resultsCh := make(chan TaskResult, len(tasks))
	for i, t := range tasks {
		wg.Add(1)
		go func(index int, t Task) {
			defer wg.Done()

			// 记录一下开始时间
			start := time.Now()

			// 实际调用
			err := t()

			duration := time.Since(start)

			resultsCh <- TaskResult{
				TaskId:   index,
				duration: duration,
				Error:    err,
			}
		}(i, t)
	}

	wg.Wait()
	close(resultsCh)

	results := make([]TaskResult, 0, len(tasks))

	for v := range resultsCh {
		results = append(results, v)
	}

	return results
}
