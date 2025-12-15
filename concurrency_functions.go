package main

import (
	"fmt"
	"sync"
	"time"
)

// ========== GOROUTINE DAN CHANNEL FUNCTIONS ==========

// Fungsi worker sederhana
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Millisecond * 500) // Simulasi kerja
		results <- job * 2
	}
}

// Fungsi untuk producer
func producer(jobs chan<- int, numJobs int) {
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)
}

// Fungsi ping-pong dengan channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// ========== FUNGSI DENGAN WAITGROUP ==========

// Fungsi yang menggunakan WaitGroup
func taskWithWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Menandai task selesai
	
	fmt.Printf("Task %d starting\n", id)
	time.Sleep(time.Millisecond * time.Duration(id*100))
	fmt.Printf("Task %d completed\n", id)
}

// ========== FUNGSI DENGAN MUTEX ==========

// Counter dengan mutex untuk thread safety
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// Fungsi untuk increment counter concurrently
func incrementCounter(counter *SafeCounter, times int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		counter.Increment()
		time.Sleep(time.Millisecond)
	}
}

// ========== SELECT DENGAN MULTIPLE CHANNELS ==========

// Fungsi dengan select statement
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Fibonacci generator quit")
			return
		}
	}
}

// Fungsi dengan timeout
func timeoutExample(ch chan string) {
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: No message received within 2 seconds")
	}
}

// ========== BUFFERED CHANNEL ==========

// Fungsi untuk demo buffered channel
func bufferedChannelDemo() {
	// Buffered channel dengan kapasitas 3
	ch := make(chan int, 3)
	
	// Mengirim tanpa blocking karena ada buffer
	ch <- 1
	ch <- 2
	ch <- 3
	
	fmt.Println("Sent 3 values to buffered channel")
	
	// Menerima nilai
	fmt.Printf("Received: %d\n", <-ch)
	fmt.Printf("Received: %d\n", <-ch)
	fmt.Printf("Received: %d\n", <-ch)
}

// ========== FAN-IN FAN-OUT PATTERN ==========

// Fungsi generator untuk fan-out
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

// Fungsi square untuk processing
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
			time.Sleep(time.Millisecond * 100)
		}
		close(out)
	}()
	return out
}

// Fan-in function untuk menggabungkan multiple channels
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}
	
	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}

// Contoh penggunaan concurrency functions
func DemoConcurrencyFunctions() {
	fmt.Println("=== CONCURRENCY FUNCTIONS ===")
	
	// Worker pool pattern
	fmt.Println("1. Worker Pool Pattern:")
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	
	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	
	// Send jobs
	go producer(jobs, 5)
	
	// Collect results
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
	
	// Ping-pong example
	fmt.Println("\n2. Ping-Pong with Channels:")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	
	ping(pings, "Hello")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	
	// WaitGroup example
	fmt.Println("\n3. WaitGroup Example:")
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go taskWithWaitGroup(i, &wg)
	}
	
	wg.Wait()
	fmt.Println("All tasks completed")
	
	// Mutex example
	fmt.Println("\n4. Mutex Example:")
	counter := &SafeCounter{}
	var wg2 sync.WaitGroup
	
	// Start 5 goroutines yang masing-masing increment 10 kali
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go incrementCounter(counter, 10, &wg2)
	}
	
	wg2.Wait()
	fmt.Printf("Final counter value: %d\n", counter.Value())
	
	// Buffered channel
	fmt.Println("\n5. Buffered Channel:")
	bufferedChannelDemo()
	
	// Select with timeout
	fmt.Println("\n6. Select with Timeout:")
	ch := make(chan string)
	
	// Test timeout
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Hello after 3 seconds"
	}()
	
	timeoutExample(ch)
	
	// Fan-in Fan-out
	fmt.Println("\n7. Fan-in Fan-out Pattern:")
	input := generator(1, 2, 3, 4, 5)
	
	// Fan-out: distribute work to multiple workers
	worker1 := square(input)
	worker2 := square(input)
	
	// Fan-in: merge results
	output := fanIn(worker1, worker2)
	
	fmt.Print("Squared results: ")
	for result := range output {
		fmt.Printf("%d ", result)
	}
	fmt.Println()
	
	// Fibonacci with select
	fmt.Println("\n8. Fibonacci with Select:")
	c := make(chan int)
	quit := make(chan int)
	
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d ", <-c)
		}
		quit <- 0
	}()
	
	fibonacci2(c, quit)
	fmt.Println()
	
	fmt.Println()
}