package lock

import "sync"

type Locker struct {
	lock sync.Mutex
}

func (l *Locker) Do(f func()) {
	l.lock.Lock()
	defer l.lock.Unlock()
	f()
}

type RWLocker struct {
	lock sync.RWMutex
}

func (l *RWLocker) RDo(f func()) {
	l.lock.RLock()
	defer l.lock.RUnlock()
	f()
}

func (l *RWLocker) WDo(f func()) {
	l.lock.Lock()
	defer l.lock.Unlock()
	f()
}
