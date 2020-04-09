package books

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}

// random generator
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func channelsDemo() {
	wg := &sync.WaitGroup{}
	// channel with a buffer of 1, this is a buffered channel
	channel := make(chan int, 1)

	wg.Add(2)

	// receiver-only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// used while extracting values from a channel that was generated by a for-loop
		for msg := range ch {
			fmt.Println(msg)
		}

		// this extracts the value and ok, (comma ok syntax) from the channel
		// value will be the value from the channel, while ok will check if the channel is closed
		if value, ok := <-ch; ok {
			fmt.Println(value)
		}

		// will receive a 0 value if the channel is closed
		fmt.Println(<-ch)

		wg.Done()
	}(channel, wg)

	// sender-only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {

		for i := 0; i < 10; i++ {
			// passing i into channel
			ch <- i
		}
		// this should only be on the send-only channel, NOT on receive-only channel
		close(ch)
		// will throw a panic!
		// panic: send on closed channel
		// ch <- 9

		wg.Done()
	}(channel, wg)

	wg.Wait()
}

func main() {
	waitGroup := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}

	cacheCh := make(chan Book)
	dbCh := make(chan Book)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		// number of tasks to wait on
		waitGroup.Add(2)

		go func(id int, waitGroup *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}

			waitGroup.Done()
		}(id, waitGroup, mutex, cacheCh)

		go func(id int, waitGroup *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryDatabase(id, m); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			waitGroup.Done()
		}(id, waitGroup, mutex, dbCh)

		// create 1 goroutine per query
		go func(cacheCh, dbCh <-chan Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(b)
				// this is because the dbCh is always sending, while the cacheCh occasionally sends
				<-dbCh
			case b := <-dbCh:
				fmt.Println("from db")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)
		time.Sleep(150 * time.Millisecond)
	}

	waitGroup.Wait()
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	// lock resource so that this goroutine has access to resource
	m.RLock()
	b, ok := cache[id]

	// unlock resource
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}

	return Book{}, false
}
