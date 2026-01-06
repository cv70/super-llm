package gtime

import (
	"runtime"
	"time"

	"k8s.io/klog/v2"
)

func LogTimeCost(thresholdMs int64, extra ...string) func() {
	st := time.Now()
	return func() {
		dur := time.Since(st)
		if dur.Milliseconds() >= thresholdMs {
			pc, _, _, ok := runtime.Caller(1)
			var callerName string
			if ok {
				f := runtime.FuncForPC(pc)
				if f != nil {
					callerName = f.Name()
				}
			}
			klog.Infof("[cost: %s][%s] %s", dur, callerName, extra)
		}
	}
}

func LogTimeCostPerJob(thresholdMs, jobCnt int64, extra ...string) func() {
	st := time.Now()
	return func() {
		dur := time.Since(st)
		if dur.Milliseconds() >= thresholdMs {
			pc, _, _, ok := runtime.Caller(1)
			var callerName string
			if ok {
				f := runtime.FuncForPC(pc)
				if f != nil {
					callerName = f.Name()
				}
			}

			klog.Infof("[cost: %s][per job cost: %s][%s] %s", dur, dur/time.Duration(jobCnt), callerName, extra)
		}
	}
}

func LogTimeCostAny(thresholdMs int64, extra ...any) func() {
	st := time.Now()
	return func() {
		dur := time.Since(st)
		if dur.Milliseconds() >= thresholdMs {
			pc, _, _, ok := runtime.Caller(1)
			var callerName string
			if ok {
				f := runtime.FuncForPC(pc)
				if f != nil {
					callerName = f.Name()
				}
			}
			for i, e := range extra {
				if f, ok := e.(func() string); ok {
					extra[i] = f()
				}
			}
			klog.Infof("[cost: %s][%s] %s", dur, callerName, extra)
		}
	}
}

func WarnLogTimeCostAny(thresholdMs int64, extra ...any) func() {
	st := time.Now()
	return func() {
		dur := time.Since(st)
		if dur.Milliseconds() >= thresholdMs {
			pc, _, _, ok := runtime.Caller(1)
			var callerName string
			if ok {
				f := runtime.FuncForPC(pc)
				if f != nil {
					callerName = f.Name()
				}
			}
			for i, e := range extra {
				if f, ok := e.(func() string); ok {
					extra[i] = f()
				}
			}
			klog.Warningf("[cost: %s][%s] %s", dur, callerName, extra)
		}
	}
}
