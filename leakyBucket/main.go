package main

import (
	"fmt"
	"sync"
	"time"
)

// 每个请求来了，需要把执行的任务逻辑封装成Task， 放入桶， 等待worker取出执行
type Task struct {
	handler func() Result // worker 从桶中取出请求对象后要执行的业务逻辑函数
	resChan chan Result   // 等待worker执行并返回结果的channel
	taskID  int
}

func handler() Result {
	time.Sleep(300 * time.Millisecond)
	return Result{}
}

func NewTask(id int) Task {
	return Task{
		handler: handler,
		resChan: make(chan Result),
		taskID:  id,
	}
}

type Result struct{}

// 漏桶
type LeakyBucket struct {
	BucketSize int       // 桶大小
	NumWorker  int       // 同时从桶中获取任务执行的worker数量
	bucket     chan Task // 存放任务的桶
}

func NewLeakyBucket(bucketSize int, numWorker int) *LeakyBucket {
	return &LeakyBucket{
		BucketSize: bucketSize,
		NumWorker:  numWorker,
		bucket:     make(chan Task, bucketSize),
	}
}

func (b *LeakyBucket) validate(task Task) bool {
	select {
	case b.bucket <- task:
	default:
		fmt.Printf("request[id=%d] is refused\n", task.taskID)
		return false
	}
	<-task.resChan
	fmt.Printf("request[id=%d] is run\n", task.taskID)
	return true
}

func (b *LeakyBucket) Start() {
	go func() {
		for i := 0; i < b.NumWorker; i++ {
			go func() {
				for {
					task := <-b.bucket
					result := task.handler()
					task.resChan <- result
				}
			}()
		}
	}()
}

func main() {
	bucket := NewLeakyBucket(10, 4)
	bucket.Start()

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			task := NewTask(id)
			bucket.validate(task)
		}(i)
	}
	wg.Wait()
}
