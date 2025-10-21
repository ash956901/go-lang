package main

import (
	"fmt"
	"sync"
)

// good practise is to include mutex
// inside the struct you are using it for
type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func(){
		wg.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock()
	// dont lock the whole logic but only the line 
	// of code which carries modification of the shared
	// resource 
	p.views += 1

}

// Race condition --> multiple process modify same resource
// the modification isnt atomic, what one process modifies can again be modified by another
// use mutexes(lock) to avoid this
func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()

	fmt.Println(myPost.views)
}
