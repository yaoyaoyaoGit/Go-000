package rolling

import (
	"testing"
	"time"
)

func TestRollingCounter(t *testing.T) {
	testCases := []struct {
		sleepTime      int
		expectedOffset int
		expectedValue  float64
	}{
		{sleepTime: 288, expectedOffset: 0, expectedValue: 1.0},
		{sleepTime: 400, expectedOffset: 2, expectedValue: 1.0},
		{sleepTime: 3100, expectedOffset: 2, expectedValue: 1.0},
		{sleepTime: 6400, expectedOffset: 2, expectedValue: 1.0},
	}

	counter := NewCounter(10, 300)

	for _, tc := range testCases {
		time.Sleep(time.Duration(tc.sleepTime) * time.Millisecond)
		counter.Add(1.0)
		if counter.window.buckets[tc.expectedOffset] != tc.expectedValue {
			t.Errorf("fail test case %v, bucket value: %v, buckets: %v", tc, counter.window.buckets[tc.expectedOffset], counter.window.buckets)
		}
	}
}

func TestRollingCountSum(t *testing.T) {
	counter := NewCounter(10, 300)
	time.Sleep(100 * time.Millisecond)
	counter.Add(1.0)
	counter.Add(1.0)
	time.Sleep(400 * time.Millisecond)
	counter.Add(1.0)
	res := counter.Sum()
	if res != 3.0 {
		t.Errorf("expected to be 3, but got %v", res)
	}
	time.Sleep(3000 * time.Millisecond)
	res = counter.Sum()
	if res != 1.0 {
		t.Errorf("expected to be 1, but got %v", res)
	}
}
