package function

import (
	"super-llm/pkg/mistake"
	"time"
)

// TimerFunc 定时执行函数，返回停止函数
func TimerFunc(fun func(), interval time.Duration) func() {
	ticker := time.NewTicker(interval)
	stopChan := make(chan struct{}, 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				fun()
			case <-stopChan:
				return
			}
		}
	}()
	return func() {
		stopChan <- struct{}{}
		ticker.Stop()
	}
}

func Retry(fn func() bool, retrys int) bool {
	for range retrys + 1 {
		if fn() {
			return true
		}
	}
	return false
}

func Must[T any](fn func() (T, error)) T {
	res, err := fn()
	if err != nil {
		mistake.Unwrap(err)
	}
	return res
}
