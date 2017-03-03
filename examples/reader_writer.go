package main

import (
	"fmt"
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLNMOPQRSTUVWXYZ"

type Reader struct {
	id int
	in chan string
}

type Writer struct {
	id  int
	out chan string
}

type Queue struct {
	readers []*Reader
	writers []*Writer
}

func main() {
	rand.Seed(time.Now().UnixNano())

	q := initQueue(20, 50)
	q.start()

	time.Sleep(time.Duration(60) * time.Second)
}

func initQueue(r, w int) *Queue {
	readers := make([]*Reader, 0)
	writers := make([]*Writer, 0)

	for i := 0; i < r; i++ {
		readers = append(readers, &Reader{id: i, in: make(chan string)})
	}

	for i := 0; i < w; i++ {
		writers = append(writers, &Writer{id: i, out: make(chan string)})
	}

	return &Queue{
		readers: readers,
		writers: writers,
	}
}

func (q *Queue) start() {
	for _, w := range q.writers {
		go w.do()
	}
	for _, r := range q.readers {
		go r.get()
	}
	q.fanOut(q.fanIn())
}

func (q *Queue) fanIn() <-chan string {
	c := make(chan string)
	for _, w := range q.writers {
		go func(in <-chan string) {
			for {
				c <- <-in
			}
		}(w.out)
	}
	return c
}

func (q *Queue) fanOut(in <-chan string) {
	for {
		s := <-in
		for _, r := range q.readers {
			r.push(s)
		}
	}
}

func (w *Writer) do() {
	for {
		s := randomString(10)
		time.Sleep(time.Duration(rand.Intn(5000)+5000) * time.Millisecond) // OMIT
		fmt.Printf("#%d Write: %s\n", w.id, s)                             // OMIT
		w.out <- s
	}
}

func (r *Reader) push(s string) {
	go func() { r.in <- s }()
}

func (r *Reader) get() {
	for {
		time.Sleep(time.Duration(rand.Intn(500)+1000) * time.Millisecond) // OMIT
		s := <-r.in
		fmt.Printf("#%d Read: %s\n", r.id, s) // OMIT
	}
}

func randomString(length int) string {
	b := make([]byte, 0)

	for i := 0; i < length; i++ {
		b = append(b, letters[rand.Intn(52)])
	}

	return string(b)
}
