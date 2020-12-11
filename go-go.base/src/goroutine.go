package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var mutex sync.Mutex

func Hello() {
	fmt.Println("Hello")
}

func main() {
	go Hello()
	// time.Sleep(time.Second) // main休眠，goroutine执行时，main可能已经执行完并退出

	// Channel创建
	ch := make(chan int)

	// 匿名函数 goroutine
	go func() {
		fmt.Println("Hello world")
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch) // 关闭channel
	}()
	// fmt.Println(<-ch) // 阻塞goroutine
	// 未关闭channel，fatal error: all goroutines are asleep - deadlock!
	for v := range ch {
		fmt.Println(v)
	}

	// 单向通道
	// c := make(chan int)
	// 声明一个只能写入数据的通道类型, 并赋值为c
	// var chSendOnly chan<- int = c
	//声明一个只能读取数据的通道类型, 并赋值为c
	// var chRecvOnly <-chan int
	// chRecvOnly <- 1 // invalid operation: chRecvOnly <- 1 (send to receive-only type <-chan int)
	// fmt.Println(<-chSendOnly) // invalid operation: <-chSendOnly (receive from send-only type chan<- int)

	// 带缓存的channel
	ch = make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}

	// cb := make(chan bool)
	// addFun := func(c chan bool, index int) {
	// 	a := 1
	// 	for i := 0; i < 1000000; i++ {
	// 		a += i
	// 	}
	// 	fmt.Println(index, a)

	// 	if 9 == index {
	// 		c <- true
	// 	}
	// }
	// for i := 0; i < 10; i++ {
	// 	go addFun(cb, i)
	// }
	// <-cb
	// fmt.Println("")

	// fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	// WaitGroup
	add2Fun := func(wg *sync.WaitGroup, index int) {
		a := 1
		for i := 0; i < 1000000; i++ {
			a += i
		}
		fmt.Println(index, a)

		wg.Done()
	}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go add2Fun(&wg, i)
	}
	wg.Wait()

	// 互斥锁
	// 缺少锁，最后count不一定等于2000000；多个goroutine可能同时操作非原子性的变量count
	count := 0
	var countLock sync.Mutex
	for i := 0; i < 2; i++ {
		go func() {
			for i := 1000000; i > 0; i-- {
				countLock.Lock()
				count++
				countLock.Unlock()
			}
			fmt.Println(count)
		}()
	}

	// 读写锁
	count = 1
	var countGuard sync.RWMutex
	read := func(m map[string]string) {
		for {
			// countGuard.Lock()
			countGuard.RLock()
			var _ string = m["name"]
			count++
			// countGuard.Unlock()
			countGuard.RUnlock()
		}
	}
	write := func(m map[string]string) {
		for {
			countGuard.Lock()
			m["name"] = "johny"
			count++
			// time.Sleep(time.Millisecond * 3)
			countGuard.Unlock()
		}
	}
	var num int = 3
	var mapA map[string]string = map[string]string{"nema": ""}

	for i := 0; i < num; i++ {
		go read(mapA)
	}
	for i := 0; i < num; i++ {
		go write(mapA)
	}
	// 执行到此处，前面的goroutine被结束
	time.Sleep(time.Second * 3)
	fmt.Printf("最终读写次数：%d\n", count)

	// select
	// Channel timeout
	ct := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case num := <-ct:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ct <- i
		time.Sleep(time.Second)
	}
	<-quit
	fmt.Println("程序结束")
}
