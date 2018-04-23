package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//限流器，简单利用令牌桶算法实现，不知道是否有错误
//每秒会有r个令牌放入桶中
//一个桶中最多可以存放b个令牌，当令牌放入桶中时，若桶已满，则令牌被直接丢弃
//若请求的令牌数大于桶中剩余的，该请求要么被缓存，要么被丢弃
// 令牌桶
type Bucket struct {
	cap       int
	available int
	interval  time.Duration
	mu        *sync.Mutex
}

func NewBucket(cap int, t time.Duration) *Bucket {
	if cap <= 0 {
		return nil
	}
	b := &Bucket{
		cap:       cap,
		available: cap,
		interval:  t,
		mu:        new(sync.Mutex),
	}
	go b.CreateToken()
	return b
}

//返回获取的令牌数和是否获取成功
func (t *Bucket) TakeToken(count int) (int, int, bool) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.available <= 0 {
		return 0, 0, false
	}
	if count > t.available {
		return t.available, 0, false
	}
	t.available = t.available - count
	return t.available + count, count, true
}

func (t *Bucket) CreateToken() {
	go func() {
		for range time.Tick(t.interval) {
			go func() {
				t.mu.Lock()
				defer t.mu.Unlock()
				if t.available < t.cap {
					t.available++
				}
			}()
		}
	}()
}

func main() {
	bucket := NewBucket(100, time.Microsecond*100)
	s1 := rand.NewSource(42) //用指定值创建一个随机数种子
	r1 := rand.New(s1)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(s int) {
			defer wg.Done()
			r := r1.Intn(10)
			remain, count, result := bucket.TakeToken(r)
			for !result {
				fmt.Println("Request Reject", s, r)
				return
			}
			fmt.Println("Request Done", s, remain, r, count, result)

		}(i)
	}
	wg.Wait()
}
