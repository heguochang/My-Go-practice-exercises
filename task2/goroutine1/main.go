package main

import (
	"fmt"
	"sync"
)

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/
func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				fmt.Printf("我是奇数协程 %d \n", i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				fmt.Printf("我是偶数协程 %d \n", i)
			}
		}
	}()

	wg.Wait()
}
