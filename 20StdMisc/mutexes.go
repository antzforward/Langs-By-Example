package main

import (
	"fmt"
	"sync"
)

// 容器保存着一个计数器映射；由于我们希望多个 goroutine 并发更新它，
// 因此我们添加了一个Mutex来同步访问。需要注意的是，互斥量不能被复制，
// 因此如果 struct传递互斥量，应该使用指针。
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {

	c.mu.Lock()
	//访问之前锁定互斥锁counters；在函数末尾使用defer 语句将其解锁。
	//defer 定义的类似IDispose在c#代码，近似于finally调用，多以说它是函数末尾调用
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	//请注意，互斥锁的零值可以按原样使用，因此这里不需要初始化。
	c := Container{

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup
	//此函数循环增加一个命名计数器。
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	//同时运行多个 goroutine；注意它们都访问同一个Container，并且其中两个访问同一个计数器。
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	//等待 goroutines 完成
	wg.Wait()
	fmt.Println(c.counters)
}

/**go run .\mutexes.go
map[a:20000 b:10000]*/
