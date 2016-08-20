package memo

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	m := New(func(key string, cancel <-chan struct{}) (interface{}, error) {
		<-cancel
		return nil, fmt.Errorf("canceled")
	})
	defer m.Close()

	cancel := make(chan struct{})
	results := make(chan result)
	go func() {
		v, err := m.Get("foo", cancel)
		results <- result{v, err}
	}()

	time.Sleep(100 * time.Millisecond)
	close(cancel)

	result := <-results
	if result.value != nil || result.err == nil {
		t.Errorf("result.value=%v, result.err=%v", result.value, result.err)
	}
}
