package onecall

import (
	"sync"
	"testing"
	"time"
)

func TestCallsShouldIncreaseCounterOnce(t *testing.T) {
	increaseCounter := func(counter *int) {
		time.Sleep(time.Second)
		*counter++
	}

	counter := 0
	g := NewGroup()
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			_, _ = g.Do("key", func() (any, error) {
				increaseCounter(&counter)
				return nil, nil
			})
			wg.Done()
		}()
	}

	wg.Wait()

	if counter != 1 {
		t.Errorf("counter should be 1, but it is %d", counter)
	}
}
