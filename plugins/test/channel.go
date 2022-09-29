package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// r, e := http.Get("http://localhost:9985/public/notify/testsetstst")
	// fmt.Println(r, e)
	// return
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			r, e := http.Get(fmt.Sprintf("http://localhost:9985/public/notify/%d", i))
			fmt.Println(r, e)
			wg.Done()
		}(i)
		fmt.Println(i)
	}
	wg.Wait()

}
