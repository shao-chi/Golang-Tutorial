Chapter 4. Concurrency

* [Goroutines](#goroutines)
* [Channels](#channels)
    * [Buffered Channels](#buffered-channels)
    * [Range and Close](#range-and-close)
* [Select](#select)
    * [Default Selection](#default-selection)
* [sync.Mutex](#syncmutex)

> Exercise:
> 1. [Equivalent Binary Trees](/4_Concurrency/equivalent_binary_trees.go)
> 2. [Web Crawler](/4_Concurrency/web_crawler.go)

### Goroutines

* A goroutine is a lightweight thread managed by the Go runtime.

```go
go f(x, y, z)
f(x, y, z)
```

* Goroutines run in the same address space, so access to shared memory must be synchronized.
* The `sync` package provides useful primitives, although you won't need them much in Go as there are other primitives.

### Channels

* Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.
* The data flows in the direction of the arrow.

```go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.
```

* Like maps and slices, channels must be created before use.

```go
ch := make(chan int)
```

* By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

```go
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```
```plain text
-5 17 12

17 -5 12
```

#### Buffered Channels

* Channels can be buffered.
* Provide the buffer length as the second argument to `make` to initialize a buffered channel.

```go
ch := make(chan int, 100)
```

* Sends to a buffered channel block only when the buffer is full.
* Receives block when the buffer is empty.

```go
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3  // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)  // fatal error: all goroutines are asleep - deadlock!
}
```

#### Range and Close

* A sender can `close` a channel to indicate that no more values will be sent.
* Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression.

```go
v, ok := <-ch
```

* `ok` is `false` if there are no more values to receive and the channel is closed.
* The loop `for i := range c` receives values from the channel repeatedly until it is closed.

* **Note**
    * Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
    * Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a `range` loop.

### Select

* The `select` statement lets a goroutine wait on multiple communication operations.
* A `select` blocks until one of its cases can run, then it executes that case. It chooses one at **random** if multiple are ready.

```go
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

#### Default Selection

* The `default` case in a `select` is run if no other case is ready.
* Use a `default` case to try a send or receive without blocking.

```go
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```

```go
func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```

### sync.Mutex

* But what if we don't need communication? What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?
* This concept is called *mutual exclusion*, and the conventional name for the data structure that provides it is *mutex*.
* We can define a block of code to be executed in mutual exclusion by surrounding it with a call to `Lock` and `Unlock` as shown on the `Inc` method.
* We can also use `defer` to ensure the mutex will be unlocked as in the `Value` method.

```go
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
```