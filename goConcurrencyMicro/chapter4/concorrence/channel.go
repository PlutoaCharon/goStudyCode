package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func printInput(ch chan string) {

	for val := range ch {
		if val == "EOF" {
			break
		}
		fmt.Printf("Input is %s\n", val)
	}
}

func consume(ch chan int) {

	time.Sleep(time.Second * 100)
	<-ch
}

func main() {
	ch := make(chan string)
	go printInput(ch)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		ch <- val
		if val == "EOF" {
			fmt.Println("End the game")
			break
		}
	}

	defer close(ch)

	//// 创建一个长度为 2 的 channel
	//ch := make(chan int, 2)
	//go consume(ch)
	//
	//ch <- 0
	//ch <- 1
	//// 发送数据不被阻塞
	//fmt.Println("I am free!")
	//ch <- 2
	//fmt.Println("I can not go there within 100s!")
	//
	//time.Sleep(time.Second)

	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//
	//go send(ch1, 0)
	//go send(ch2, 10)
	//
	//// 主 goroutine 休眠 1s，保证调度成功
	//time.Sleep(time.Second)
	//
	//for {
	//	select {
	//	case val := <- ch1: // 从 ch1 读取数据
	//		fmt.Printf("get value %d from ch1\n", val)
	//	case val := <- ch2 : // 从 ch2 读取数据
	//		fmt.Printf("get value %d from ch2\n", val)
	//	case <-time.After(2 * time.Second): // 超时设置
	//		fmt.Println("Time out")
	//		return
	//	}
	//}
}

func send(ch chan int, begin int) {

	for i := begin; i < begin+10; i++ {
		ch <- i

	}

}
