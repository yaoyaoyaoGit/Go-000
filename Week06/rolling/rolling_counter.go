package rolling

import (
	"fmt"
	"time"
)

type (
	Counter struct {
		window *window
	}

	window struct {
		buckets        []float64
		size           int
		bucketSizeInMs time.Duration
		lastAddAt      time.Time
		offset         int
	}
	reducer func(values []float64) float64
)

func NewCounter(numOfBuckets int, bucketSizeInMs int64) Counter {
	return Counter{newWindow(numOfBuckets, time.Duration(bucketSizeInMs)*time.Millisecond)}
}

func (c *Counter) Add(value float64) {
	c.window.add(value)
	return
}

func (c *Counter) Sum() float64 {
	f := func(values []float64) float64 {
		s := 0.0
		for _, v := range values {
			s += v
		}
		return s
	}
	return c.window.reduce(f)
}

func (c *Counter) Avg() float64 {
	f := func(values []float64) float64 {
		s := 0.0
		for _, v := range values {
			s += v
		}
		return s / float64(len(values))
	}
	return c.window.reduce(f)
}

func newWindow(size int, bucketSizeInMs time.Duration) *window {
	b := make([]float64, size)
	return &window{
		buckets:        b,
		size:           size,
		bucketSizeInMs: bucketSizeInMs,
		lastAddAt:      time.Now(),
		offset:         0,
	}
}

func (w *window) add(value float64) {
	ts := w.getTimespan()
	if ts == 0 {
		w.buckets[w.offset] += value
		return
	}
	w.lastAddAt = w.lastAddAt.Add(time.Duration(ts) * w.bucketSizeInMs)
	if ts > w.size {
		ts = w.size
	}
	s := w.offset + 1
	e, e1 := w.offset+ts, 0
	ns := e
	if e > w.size {
		e1 = e - w.size
		e = w.size
		ns = e1
	}
	for i := s; i < e; i++ {
		w.buckets[i] = 0
	}
	for i := 0; i <= e1; i++ {
		w.buckets[i] = 0
	}
	w.buckets[ns] = value
	w.offset = ns
	return
}

func (w *window) getTimespan() int {
	return int(time.Since(w.lastAddAt) / w.bucketSizeInMs)
}

func (w *window) reduce(f reducer) float64 {
	ts := w.getTimespan()
	count := w.size - ts
	if count <= 0 {
		return 0.0
	}
	values := make([]float64, count)
	for i := range values {
		fmt.Println((w.offset+ts+i)%w.size, w.buckets[(w.offset+ts+i)%w.size])
		values[i] = w.buckets[(w.offset+ts+i)%w.size]
	}
	return f(values)
}
