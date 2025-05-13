package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Queue struct {
	buffer []rune
	mu     sync.Mutex
	cond   sync.Cond
}

func NewWordQueue() *Queue {
	wq := &Queue{
		buffer: make([]rune, 0),
	}
	wq.cond = *sync.NewCond(&wq.mu)
	return wq
}

func (q *Queue) Enqueu(char rune) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.buffer = append(q.buffer, char)

	if char == ' ' && len(q.buffer) > 1 {
		q.cond.Broadcast()
	}
}

func (q *Queue) Dequeue() string {
	q.mu.Lock()
	defer q.mu.Unlock()

	for {
		// Look for space in the buffer
		spaceIndex := -1
		for i, char := range q.buffer {
			if char == ' ' {
				spaceIndex = i
				break
			}
		}

		// If we found a space and there's content before it
		if spaceIndex >= 0 {
			// Extract the word (excluding the space)
			word := string(q.buffer[:spaceIndex])

			// Remove the word and the space from the buffer
			q.buffer = q.buffer[spaceIndex+1:]

			return word
		}

		// No complete word available yet, wait for notification
		q.cond.Wait()
	}
}

func (q *Queue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.buffer)
}

func Publish(q *Queue, numChar int, wg *sync.WaitGroup) {
	defer wg.Done()

	letter := []rune("abcdefghijklmnopqrstuvwxyz     ") // Extra spaces for more words

	for i := range numChar {
		char := letter[rand.Intn(len(letter))]

		q.Enqueu(char)
		fmt.Printf("Publisher %d added '%c' \n", i, char)

		time.Sleep(time.Duration(rand.Intn(50)+10) * time.Millisecond)
	}
}

func Consumer(q *Queue, id int, numWords int, wg *sync.WaitGroup) {
	defer wg.Done()

	for range numWords {
		word := q.Dequeue()

		fmt.Printf("Consumer %d consumeed %s\n", id, word)

		time.Sleep(time.Duration(rand.Intn(100)+50) * time.Millisecond)
	}
}


func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	wordQueue := NewWordQueue()

	var wg sync.WaitGroup

	numPublisher := 3

	for range numPublisher {
		wg.Add(1)
		go Publish(wordQueue, 20, &wg)
	}

	numConsumer := 2
	for i := range numConsumer {
		wg.Add(1)
		go Consumer(wordQueue, i, 5, &wg)
	}

	wg.Wait()

	fmt.Println("Demo completed")

}