package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tour/tree"
)

func Test_Concurrency1(t *testing.T) {
	b := false

	go func(v bool) { b = v }(true)

	for !b {
		// NOP
	}

	assert.True(t, b)
}

func Test_Concurrency2(t *testing.T) {
	ch := make(chan int)
	defer close(ch)

	go SumStream([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, ch)

	sum := <-ch

	assert.Equal(t, 45, sum)
}

func Test_Concurrency3(t *testing.T) {
	ch := make(chan int, 100)
	// No `close` API call

	go SumStream([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, ch)

	sum := <-ch

	assert.Equal(t, 45, sum)

	close(ch)
}

func Test_Concurrency4(t *testing.T) {
	ch := make(chan int, 100)
	defer close(ch)

	go SumStream([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, ch)

	sum := <-ch

	assert.Equal(t, 45, sum)
}

func Test_Concurrency5(t *testing.T) {
	ch := make(chan int)
	defer close(ch)

	quit := make(chan int)
	defer close(quit)

	go func() {
		for i := range 10 {
			fmt.Printf("fibo(%d)=%d\n", i, <-ch)
		}
		quit <- 0
	}()

	FibonacciStream(ch, quit)

}

func Test_Concurrency6(t *testing.T) {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(1 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Printf("tick")
		case <-boom:
			fmt.Println(" BOOM!")
			return
		default:
			time.Sleep(50 * time.Millisecond)
			fmt.Printf(".")
		}
	}
}

func Test_Concurrency7(t *testing.T) {
	// Nothing to be implemented here

	assert.True(t, true)
}

func Test_Concurrency8(t *testing.T) {
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	sum := 0
	for range 10 {
		sum += <-ch
	}

	assert.Equal(t, 55, sum)

	assert.True(t, Same(tree.New(1), tree.New(1)))
	assert.False(t, Same(tree.New(1), tree.New(2)))
}

func Test_Concurrency9(t *testing.T) {
	sc := &SafeCounter{sync.Mutex{}, make(map[string]int)}

	n := 100
	key := NewRandomString()
	for range n {
		go func() { sc.Increment(key) }()
	}

	for sc.GetValue(key) < n {
		time.Sleep(10 * time.Millisecond)
	}

	assert.Equal(t, n, sc.GetValue(key))
}

func Test_Concurrency10(t *testing.T) {
	fetcher := &FetcherMock{}

	res := &FetchResults{
		sync.Mutex{},
		make(map[string]*FetchResult),
	}

	Crawl("https://golang.org/", 4, fetcher, res)
}
