package main

import (
	"bytes"
	"fmt"
	"sync"
)

func sayHello() {
	fmt.Println("hello world")
}

func main() {
	//	go sayHello()
	//	time.Sleep(1 * time.Second)
	//
	//	go func() {
	//		fmt.Println("hello, goroutine")
	//	}()
	//	time.Sleep(1 * time.Second)
	//
	//	sayhello := func() {
	//		fmt.Println("hello, goroutine2")
	//	}
	//	go sayhello()
	//	time.Sleep(1 * time.Second)
	//	// go的并发采用了fork-join模型
	//	// 以上程序都没有创建join点
	//	// 为了创建join点，我们需要引入sync包
	//
	//	// 创建一个计数器
	//var wg sync.WaitGroup
	//	sayhello2 := func() {
	//		defer wg.Done()
	//		fmt.Println("hello, goroutine3")
	//	}
	//	// fork了一个协程所以计数器加一
	//	wg.Add(1)
	//	go sayhello2()
	//	wg.Wait()
	//
	//	salutation := "hello"
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		salutation = "go"
	//	}()
	//	wg.Wait()
	//	fmt.Println(salutation) // 这里打印了go，所以说明创建的地址空间是相同的
	//
	//	for _, salutation2 := range []string{"hello", "greetings", "good day"} {
	//		wg.Add(1)
	//		go func() {
	//			defer wg.Done()
	//			fmt.Println(salutation2) // 这里我们引用字符串类型的切片创建salutation的值
	//		}()
	//	}
	//	wg.Wait()
	//
	//	// C 并没有回收 被丢弃的goroutine
	//	go func() {
	//
	//	}()
	//	// 这里的goroutine将一直存在知道进程退出

	//memConsumed := func() uint64 {
	//	runtime.GC()
	//	var s runtime.MemStats
	//	runtime.ReadMemStats(&s)
	//	return s.Sys
	//}
	//var c <-chan interface{}
	//noop := func() {
	//	wg.Done()
	//	<-c
	//}
	//const numGoroutines = 1e6
	//wg.Add(numGoroutines)
	//before := memConsumed()
	//for i := numGoroutines; i > 0; i-- {
	//	go noop()
	//}
	//wg.Wait()
	//after := memConsumed()
	//fmt.Printf(" %.3fkb", float64(after-before)/numGoroutines/1000)
	//// 输出结果是 9.496kb
	//
	//// mutex
	//var count int
	//var lock sync.Mutex
	//increment := func() {
	//	lock.Lock()
	//	defer lock.Unlock()
	//	count++
	//	fmt.Println("count: ", count)
	//}
	//decrement := func() {
	//	lock.Lock()
	//	defer lock.Unlock()
	//	count--
	//	fmt.Println("count: ", count)
	//}
	//var arithmetic sync.WaitGroup
	//for i := 0; i < 10; i++ {
	//	arithmetic.Add(1)
	//	go func() {
	//		defer arithmetic.Done()
	//		increment()
	//	}()
	//}
	//for i := 0; i < 10; i++ {
	//	arithmetic.Add(1)
	//	go func() {
	//		defer arithmetic.Done()
	//		decrement()
	//	}()
	//}
	//arithmetic.Wait()
	//fmt.Printf(" %v goroutines finished", numGoroutines)

	//// 生产者-消费者模型
	//producer := func(wg *sync.WaitGroup, l sync.Locker) {
	//	defer wg.Done()
	//	for i := 0; i < 5; i++ {
	//		l.Lock()
	//		l.Unlock()
	//		time.Sleep(time.Millisecond)
	//	}
	//}
	//observer := func(wg *sync.WaitGroup, l sync.Locker) {
	//	defer wg.Done()
	//	l.Lock()
	//	defer l.Unlock()
	//}
	//test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
	//	var wg sync.WaitGroup
	//	wg.Add(count + 1)
	//	beginTestTime := time.Now()
	//	go producer(&wg, mutex)
	//	for i := count; i > 0; i-- {
	//		go observer(&wg, rwMutex)
	//	}
	//	wg.Wait()
	//	return time.Since(beginTestTime)
	//}
	//
	//tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	//defer tw.Flush()
	//var m sync.RWMutex
	//fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")
	//
	//for i := 0; i < 20; i++ {
	//	count := int(math.Pow(2, float64(i)))
	//
	//	fmt.Fprintf(
	//		tw,
	//		"%d\t%v\t%v\n",
	//		count,
	//		test(count, &m, m.RLocker()),
	//		test(count, &m, &m),
	//	)
	//}

	//// cond:一个协程的集合点，等待或发布一个event
	//var mu sync.Mutex
	//cond := sync.NewCond(&mu)
	//queue := make([]int, 0)
	//const maxCapacity = 2
	//
	//rmqueue := func(id int) {
	//	mu.Lock()
	//	for len(queue) == 0 {
	//		fmt.Println("[消费者 %d] 队列空了，我因为等待数据而挂起...\n", id)
	//		cond.Wait()
	//	}
	//	item := queue[0]
	//	queue = queue[1:]
	//	fmt.Printf("[消费者 %d] 消费了数据: %d\n", id, item)
	//
	//	mu.Unlock()
	//}
	//
	//addToQueue := func(item int) {
	//	mu.Lock()
	//	fmt.Printf("[生产者] 添加数据: %d\n", item)
	//	queue = append(queue, item)
	//	mu.Unlock()
	//
	//	// Signal 可以在 Unlock 之前或之后调用，通常建议在 Unlock 之后
	//	fmt.Printf("[生产者] 发出通知...\n")
	//	cond.Signal() // 唤醒一个正在 Wait 的消费者
	//}
	//
	//go rmqueue(1)
	//go rmqueue(2)
	//
	//time.Sleep(1 * time.Second)
	//
	//addToQueue(100)
	//time.Sleep(1 * time.Second)
	//addToQueue(200)
	//
	//time.Sleep(1 * time.Second)

	// once

	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		conn := GetDB()
	//		// 只是为了证明大家都拿到了同一个结果
	//		_ = conn
	//	}()
	//}
	//
	//wg.Wait()
	//fmt.Println("所有请求完成")

	//// Pool
	//buf := bufferPool.Get().(*bytes.Buffer)
	//buf.WriteString("hello user a")
	//fmt.Println("a正在使用", buf.String())
	//buf.Reset()
	//bufferPool.Put(buf)
	//fmt.Println("--- 请求 B 来了 ---")
	//
	//bufB := bufferPool.Get().(*bytes.Buffer)
	//
	//bufB.WriteString("Hello User B")
	//fmt.Println("B 正在使用:", bufB.String())
	//
	//bufB.Reset()
	//bufferPool.Put(bufB)
}

//var (
//	once sync.Once
//	db   string
//	wg   sync.WaitGroup
//)
//
//func setupDB() {
//	fmt.Println("setup db")
//	db = "MySQL Connection"
//}
//func GetDB() string {
//	// 使用 Do 方法包裹初始化逻辑
//	// 无论调用多少次 GetDB，setupDB 只会在第一次被执行
//	once.Do(setupDB)
//	return db
//}

var bufferPool = sync.Pool{
	// 定义 New 函数：当池子空的时候，怎么造一个新的？
	New: func() interface{} {
		fmt.Println("池子空了，申请新的内存...")
		return new(bytes.Buffer)
	},
}
