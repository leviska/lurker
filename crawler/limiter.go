package crawler

import (
	"net/url"
	"sync"
)

type RateLimiter struct {
	rates map[string]chan struct{}
	lock  sync.RWMutex

	concurrentCount uint
}

func NewRateLimiter(concurrentCount uint) *RateLimiter {
	return &RateLimiter{
		rates: map[string]chan struct{}{},

		concurrentCount: concurrentCount,
	}
}

func (r *RateLimiter) tryGetLocked(host *url.URL) bool {
	if ch, has := r.rates[host.Host]; has {
		ch <- struct{}{}
		return true
	}
	return false
}

func (r *RateLimiter) tryGet(host *url.URL) bool {
	r.lock.RLock()
	defer r.lock.RUnlock()

	return r.tryGetLocked(host)
}

func (r *RateLimiter) createAndGet(host *url.URL) {
	r.lock.Lock()
	defer r.lock.Unlock()

	// one more check
	if r.tryGetLocked(host) {
		return
	}
	r.rates[host.Host] = make(chan struct{}, r.concurrentCount)
	if !r.tryGetLocked(host) {
		panic("couldn't get rate limiter for unknown reason")
	}
}

func (r *RateLimiter) Get(host *url.URL) {
	if r.tryGet(host) {
		return
	}
	r.createAndGet(host)
}

func (r *RateLimiter) Put(host *url.URL) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	<-r.rates[host.Host]
}
