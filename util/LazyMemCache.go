package main

import (
	"sync"
	"sync/atomic"
	"time"
)

type LazyMemCache struct {
	data   interface{}
	flag   uint32
	lock   sync.Mutex
	expire time.Duration
	do     LazyMemCacheDo
}

type LazyMemCacheDo func() (interface{}, error)

func NewMemCache(expire time.Duration, do LazyMemCacheDo) *LazyMemCache {
	return &LazyMemCache{expire: expire, do: do}
}

func (cache *LazyMemCache) prepare() (err error) {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	if cache.flag != 0 {
		return nil
	}
	defer func() {
		if err == nil {
			atomic.StoreUint32(&cache.flag, 1)
		}
	}()
	var newData interface{}
	if newData, err = cache.do(); err == nil {
		cache.data = newData
		time.AfterFunc(cache.expire, func() { atomic.StoreUint32(&cache.flag, 0) })
	}
	return
}

func (cache *LazyMemCache) Load() (data interface{}, err error) {
	if atomic.LoadUint32(&cache.flag) == 0 {
		err = cache.prepare()
	}
	return cache.data, err
}
