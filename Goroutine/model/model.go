package main

import (
	"fmt"
	"net/http"
)

func chanOwner() <-chan int {
	results := make(chan int, 5)

	go func() {
		defer close(results)
		for i := 0; i <= 5; i++ {
			results <- i
		}
	}()
	return results
}

func main() {
	//results := chanOwner()
	//for res := range results {
	//	fmt.Println(res)
	//}

	// ------for-select-------
	//for { // 循环：保证 Goroutine 不退出，一直活着
	//	select {
	//	// 多路复用：同时监听多个通道，哪个有消息处理哪个
	//	case <-channelA:
	//		// do something
	//	case <-channelB:
	//		// do something
	//	}
	//}

	// demo1
	//done := make(chan struct{})
	//data := []string{"a", "b", "c", "d", "e"}
	//generator := func(done <-chan struct{}, strings []string) <-chan string {
	//	out := make(chan string)
	//	go func() {
	//		defer close(out)
	//		for _, s := range strings {
	//			select {
	//			case <-done:
	//				return
	//			case out <- s:
	//
	//			}
	//		}
	//	}()
	//	return out
	//}
	//
	//stream := generator(done, data)
	//for i := 0; i < 2; i++ {
	//	fmt.Println(<-stream)
	//}
	//close(done)
	//
	//// demo2
	//// 创建一个可以取消的上下文（代替done）
	//ctx, cancel := context.WithCancel(context.Background())
	//go func() {
	//	workStream := make(chan int)
	//
	//	go func() {
	//		for i := 0; ; i++ {
	//			workStream <- i
	//			time.Sleep(time.Millisecond * 100)
	//		}
	//	}()
	//	// for-select
	//	for {
	//		select {
	//		case task := <-workStream:
	//			fmt.Println("正在处理： %d...\n", task)
	//		case <-time.After(time.Second):
	//			fmt.Println("Worker: 心跳检测 - 我还活着")
	//		case <-ctx.Done():
	//			fmt.Println("Worker: 收到停止信号，清理资源，准备下班！")
	//			return
	//		}
	//	}
	//}()
	//time.Sleep(3 * time.Second)
	//// 发出停止信号
	//fmt.Println("Main: 系统关闭，通知 Worker 下线")
	//cancel()
	//// 等一会看 Worker 的遗言
	//time.Sleep(1 * time.Second)

	//--------default分支-------
	//for {
	//	select {
	//	case req := <-requests:
	//		handle(req)
	//	default:
	//		// 当 requests 通道没数据时，select 不会阻塞，而是立刻走这里
	//		// 用于：
	//		// 1. 轮询 (Polling)
	//		// 2. 占满 CPU (Spin Lock) - 小心使用！
	//		// 3. 尝试性发送/接收
	//	}
	//}

	//----------or-goroutine------------
	//辅助函数：创建一个 n 之后关闭的通道
	//sig := func(after time.Duration) <-chan interface{} {
	//	c := make(chan interface{})
	//	go func() {
	//		defer close(c)
	//		time.Sleep(after)
	//	}()
	//	return c
	//}
	//
	//start := time.Now()
	//
	//// 监听 5 个通道：
	//// 分别在 1小时, 5分钟, 1秒, 1小时, 1分钟 后关闭
	//// 显然，那个 "1秒" 的通道会最先触发
	//<-or(
	//	sig(1*time.Hour),
	//	sig(5*time.Minute),
	//	sig(1*time.Second),
	//	sig(1*time.Hour),
	//	sig(1*time.Minute),
	//)
	//
	//fmt.Printf("Or-channel 完成，耗时: %v\n", time.Since(start))

	////------------错误处理-------------
	//done := make(chan struct{})
	//defer close(done)
	//
	//urls := []string{"https://www.google.com", "https://bad.host", "https://www.baidu.com"}
	//for result := range checkStatus(done, urls...) {
	//	if result.Error != nil {
	//		fmt.Println("Error: %v (on %s)\n", result.Error, result.Url)
	//		continue
	//	}
	//	fmt.Println("Response: %v (on %s)\n", result.Response.Status, result.Url)
	//}

	// ----------构建流水线-------------
	done := make(chan struct{})
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4)

	pipeline := add(done, multiply(done, intStream, 2), 1)

	for v := range pipeline {
		fmt.Println(v)
	}
}

// ----------防止协程泄露--------------
// 父子协程联动
func doWork(done <-chan struct{}, strings <-chan string) <-chan interface{} {
	completed := make(chan interface{})
	go func() {
		defer fmt.Println("doWork 安全退出")
		defer close(completed)

		for {
			select {
			case s := <-strings:
				// 正常处理
				fmt.Println(s)
			case <-done: // 【逃生门】
				return
			}
		}
	}()
	return completed
}

// 使用递归实现or-Channel
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})

	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

// ---------错误处理---------
type Result struct {
	Error    error
	Response *http.Response
	Url      string
}

func checkStatus(done <-chan struct{}, urls ...string) <-chan Result {
	results := make(chan Result)
	go func() {
		defer close(results)

		for _, url := range urls {
			var result Result
			result.Url = url
			resp, err := http.Get(url)

			result.Error = err
			result.Response = resp

			select {
			case <-done:
				return
			case results <- result:

			}
		}
	}()
	return results
}

// ----------构建流水线-----------

// 1. 数据源生成器 (Generator)：把切片变成 Channel 流
// 作用：将静态数据转化为流动数据
func generator(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// 2. 处理阶段 A：乘法器
// 作用：从上游拿数据，乘以 2，发给下游
func multiply(done <-chan struct{}, inputStream <-chan int, multiplier int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range inputStream {
			result := i * multiplier
			select {
			case out <- result:
			case <-done:
				return
			}
		}
	}()
	return out
}

// 3. 处理阶段 B：加法器
// 作用：从上游拿数据，加上 1，发给下游
func add(done <-chan struct{}, inputStream <-chan int, additive int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range inputStream {
			result := i + additive
			select {
			case out <- result:
			case <-done:
				return
			}
		}
	}()
	return out
}
