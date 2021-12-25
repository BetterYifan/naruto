package main

import (
	"fmt"
	"sync"
)

/**
用来控制满足一定条件的生产消费
*/
type EatJob struct {
	mu *sync.Mutex
	c  *sync.Cond

	Eater []int
}

// 当吃饭人数大于2人时，赶出多余的人
//func main() {
//	job := &EatJob{
//		mu:    new(sync.Mutex),
//		Eater: make([]int, 0),
//	}
//	add := func() {
//		fmt.Println("添加eater, 此时有： ", len(job.Eater), job.Eater)
//	}
//	reduce := func() {
//		fmt.Println("减少eater, 此时有： ", len(job.Eater), job.Eater)
//	}
//	job.c = sync.NewCond(job.mu)
//
//	// 消费
//	consume := func(t time.Duration) {
//		time.Sleep(t)
//		job.c.L.Lock()
//		job.Eater = job.Eater[1:]
//		reduce()
//		job.c.L.Unlock()
//		job.c.Signal()
//	}
//
//	// 生产. 添加eater
//	for i := 0; i < 10; i++ {
//		job.c.L.Lock()
//		for len(job.Eater) == 2 {
//			job.c.Wait()
//		}
//		job.Eater = append(job.Eater, i)
//		add()
//		go consume(time.Second)
//		job.c.L.Unlock()
//	}
//}

func main() {
	cond := sync.NewCond(new(sync.Mutex))
	condition := 0

	// 消费者
	go func() {
		for i := 0; i < 100; i++ {
			// 消费者开始消费时，锁住
			cond.L.Lock()
			// 如果没有可消费的值，则等待
			for condition == 0 {
				cond.Wait()
			}
			// 消费
			condition--
			fmt.Printf("Consumer: %d\n", condition)

			// 唤醒一个生产者
			cond.Signal()
			// 解锁
			cond.L.Unlock()
		}
	}()

	// 生产者
	for i := 0; i < 100; i++ {
		// 生产者开始生产
		cond.L.Lock()

		// 当生产太多时，等待消费者消费
		for condition == 10 {
			cond.Wait()
		}
		// 生产
		condition++
		fmt.Printf("Producer: %d\n", condition)

		// 通知消费者可以开始消费了
		cond.Signal()
		// 解锁
		cond.L.Unlock()
	}
}
