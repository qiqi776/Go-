package main

func main() {
	//ch := make(chan int)
	//go func() {
	//	fmt.Println("正在发送数据")
	//	ch <- 10
	//	fmt.Println("发送完毕")
	//}()
	//time.Sleep(time.Second)
	//v := <-ch
	//fmt.Println("v is", v)

	//ch := make(chan int, 2)
	//
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		ch <- i
	//	}
	//	close(ch)
	//}()
	//
	//for val := range ch {
	//	fmt.Println(val)
	//}
	//fmt.Println("over")

	// select
	//start := time.Now()
	//c := make(chan interface{})
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	close(c)
	//}()
	//fmt.Println("over")
	//select {
	//case <-c:
	//	fmt.Println("Unblocked %v later.\n", time.Since(start))
	//}

	//c1 := make(chan interface{})
	//close(c1)
	//c2 := make(chan interface{})
	//close(c2)
	//var count1, count2 int
	//for i := 0; i < 100; i++ {
	//	select {
	//	case <-c1:
	//		count1++
	//	case <-c2:
	//		count2++
	//	}
	//}
	//fmt.Println("count1: %d\n, count2: %d\n", count1, count2)
}
