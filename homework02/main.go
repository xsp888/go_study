package main

import (
	"fmt"
	"homework02/models"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	coroutine()
	coroutineTask()
	areaPerimeter()
	printEmployee()
	channel()
	bufferChannel()
	lock()
	lockAtomic()
}

// 编写一个程序，使用 go 关键字启动两个协程，
// 一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
func coroutine() {

	go func() {
		for i := 1; i < 10; i++ {
			if i%2 == 1 {
				fmt.Println(i)
			}

		}
	}()

	go func() {
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}

		}
	}()
	time.Sleep(100 * time.Millisecond)

}

// 设计一个任务调度器，接收一组任务（可以用函数表示），
// 并使用协程并发执行这些任务，同时统计每个任务的执行时间
func coroutineTask() {
	var wg sync.WaitGroup

	for i := 100; i < 200; i += 10 {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			// 开始时间
			start := time.Now()
			// 执行业务
			timeTask(i)
			// 结束时间
			duration := time.Since(start)
			fmt.Printf("%d任务消耗的时间：%v\n", i, duration)
		}(i)

	}
	wg.Wait()

}

func timeTask(times int) {
	// 模拟业务操作
	time.Sleep(time.Duration(times) * time.Millisecond)
}

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法
func areaPerimeter() {

	// 矩形
	rectangle := models.Rectangle{Length: 10, Width: 3}
	// 圆
	circle := models.Circle{Radius: 4}

	shapes := []models.Shape{rectangle, circle}
	for _, shape := range shapes {
		shape.Area()
		shape.Perimeter()
	}

}

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息
func printEmployee() {
	employee := models.Employee{Person: models.Person{Name: "张三", Age: 18}, EmployeeID: 111111111}
	employee.PrintInfo()

}

// 编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，
// 另一个协程从通道中接收这些整数并打印出来
func channel() {
	// 通道
	ch := make(chan int)
	// 生产者
	go producer(ch)
	// 消费者
	go consumer(ch)
	time.Sleep(2 * time.Second)

}

func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("消费者接收: %d\n", num)
	}

}

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，
// 消费者协程从通道中接收这些整数并打印
func bufferChannel() {
	// 通道
	ch := make(chan int, 10)
	go bufferProducer(ch)
	go bufferConsumer(ch)
	time.Sleep(2 * time.Second)

}

func bufferProducer(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func bufferConsumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("缓冲区消费者接收: %d\n", num)
	}

}

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，
// 最后输出计数器的值
func lock() {
	num := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				num++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("sync.Mutex计数: %d\n", num)

}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，
// 最后输出计数器的值
func lockAtomic() {
	var num int64 = 0
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&num, 1)
				// num++
				if j%100 == 0 {
					time.Sleep(2 * time.Microsecond)
				}
			}
		}()
	}
	wg.Wait()
	fmt.Printf("atomic计数: %d\n", num)

}
