### Basics:
- `goroutines`: very lightweight threads
- `coroutine`: a group of goroutines
- every go program has at least a goroutine (the main goroutine)
- `waitGroup`: is a synchronization primitive provided by the `sync` package. It is commonly used for coordinating and waiting for a collection of Goroutines (concurrent threads) to finish their execution before proceeding. The `sync.WaitGroup` type helps ensure that the main Goroutine (or any other controlling Goroutine) waits for all other Goroutines to complete their tasks. In waitGroups, order is not guranteed
```go
package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("This is the second thing to be printed!", &wg)
}
```

___
### Race Conditions
- mutex: mutual exclusion -> allow us deal with race conditions
- lock/unlock
- race condition: multiple goroutines try to access the same data
- channels: means of having goroutines share data
- Channels are Go's Philosophy: have things share memory by communicating, rather than communicating by sharing memory
- `go run -race .`: run a program with race detector enabled
- `go test -race .`: run tests with race detector enabled
- to use mutex in go
- lock/unlock mechanism

  ```go
  func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
    }
```