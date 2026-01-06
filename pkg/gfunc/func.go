package gfunc

import (
	"time"

	"k8s.io/klog/v2"
)

// TimerFunc 定时执行函数，返回停止函数
func TimerFunc(fun func(), interval time.Duration) func() {
	ticker := time.NewTicker(interval)
	stopChan := make(chan struct{}, 1)
	go func() {
		klog.Infof("start timer func, interval: %v", interval)
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
		klog.Infof("stop timer func")
	}
}

// RetryFunc 重试函数，一个执拗的功能函数，只想把交给它的事情做好
func RetryFunc(fun func() bool, tryCount int) bool {
	for i := 0; i < tryCount; i++ {
		if fun() {
			return true
		}
	}

	return false
}

func Must[T any](fn func() (T, error)) T {
	res, err := fn()
	if err != nil {
		klog.Fatal(err)
	}
	return res
}
