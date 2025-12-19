package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Employee struct {
	id int
}

type Item1 struct {
	id int
}

type Item2 struct {
	id int
}

type Item3 struct {
	id int
}

type Item interface {
	// Process 這是一個耗時操作
	Process()
}

func (i *Item1) Process() {
	time.Sleep(100 * time.Millisecond)
}

func (i *Item2) Process() {
	time.Sleep(200 * time.Millisecond)
}

func (i *Item3) Process() {
	time.Sleep(300 * time.Millisecond)
}

const (
	worker = 5
	count  = 10
)

func main() {
	rand.Seed(time.Now().UnixNano())

	all := make([]Item, 0, count*3)
	id := 1

	for i := 0; i < count; i++ {
		all = append(all, &Item1{id: id})
		id++
	}
	for i := 0; i < count; i++ {
		all = append(all, &Item2{id: id})
		id++
	}
	for i := 0; i < count; i++ {
		all = append(all, &Item3{id: id})
		id++
	}

	// 打亂順序
	rand.Shuffle(len(all), func(i, j int) {
		all[i], all[j] = all[j], all[i]
	})

	itemCh := make(chan Item, 10)
	done := make([]int, worker)

	var wg sync.WaitGroup
	wg.Add(worker)

	begin := time.Now()

	for i := 0; i < worker; i++ {
		staff := &Employee{id: i}

		go func(staff *Employee) {
			defer wg.Done()

			for it := range itemCh {
				start := time.Now()
				fmt.Printf("[%s] 員工%d 開始處理%T\n",
					start.Format("15:04:05.000"), staff.id, it)

				it.Process()

				end := time.Now()
				fmt.Printf("[%s] 員工%d 完成處理%T，花費 %s\n",
					end.Format("15:04:05.000"), staff.id, it, end.Sub(start))

				done[staff.id]++
			}
		}(staff)
	}

	go func() {
		for _, it := range all {
			itemCh <- it
		}
		close(itemCh)
	}()
	wg.Wait()

	fmt.Printf("共花費: %s 完成 %d 個項目\n", time.Since(begin), len(all))
	for i := 0; i < worker; i++ {
		fmt.Printf("員工%d做了%d個\n", i, done[i])
	}
}
