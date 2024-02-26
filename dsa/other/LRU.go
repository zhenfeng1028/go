package main

import "container/list"

type LRU interface {
	Get(key interface{}) (value interface{}, ok bool)
	Set(key, value interface{})
	Rem(key interface{})
	Contains(key interface{}) bool
	Len() int
}

type EvictCallBack func(key, value interface{})

type SimpleLRU struct {
	// 缓存容量
	capacity int

	// 散列表
	items map[interface{}]*list.Element

	// 驱逐队列 双向列表
	evictList *list.List

	// 驱逐时的回调
	onEvict EvictCallBack
}

type KVP struct {
	key   interface{}
	value interface{}
}

func NewSimpleLRU(capacity int, onEvict EvictCallBack) *SimpleLRU {
	return &SimpleLRU{
		capacity:  capacity,
		items:     make(map[interface{}]*list.Element, 8),
		evictList: list.New(),
		onEvict:   onEvict,
	}
}

func (s *SimpleLRU) Get(key interface{}) (value interface{}, ok bool) {
	// 快速查找，散列表中是否存在
	if e, ok := s.items[key]; ok {
		// 移到队首
		s.evictList.MoveToFront(e)
		kvp, _ := e.Value.(*KVP)
		return kvp.value, true
	}
	return nil, false
}

func (s *SimpleLRU) Set(key, value interface{}) {
	// 快速查找，散列表中是否存在
	if e, ok := s.items[key]; ok {
		kvp, _ := e.Value.(*KVP)
		kvp.value = value
		// 移到队首
		s.evictList.MoveToFront(e)
		return
	}

	// 插入新节点
	kvp := &KVP{
		key:   key,
		value: value,
	}

	e := s.evictList.PushFront(kvp)
	s.items[key] = e

	// 新增缓存，判断容量是否已满
	if s.evictList.Len() > s.capacity {
		// 驱逐最老节点
		s.remOldest()
	}
}

func (s *SimpleLRU) Rem(key interface{}) {
	if e, ok := s.items[key]; ok {
		s.removeElement(e)
	}
}

func (s *SimpleLRU) Contains(key interface{}) bool {
	_, ok := s.items[key]
	return ok
}

func (s *SimpleLRU) Len() int {
	return s.evictList.Len()
}

func (s *SimpleLRU) remOldest() {
	// 取出最老节点
	e := s.evictList.Back()
	if e != nil {
		s.removeElement(e)
	}
}

func (s *SimpleLRU) removeElement(e *list.Element) {
	s.evictList.Remove(e)

	kvp, _ := e.Value.(*KVP)
	delete(s.items, kvp.key)

	// 回调
	if s.onEvict != nil {
		s.onEvict(kvp.key, kvp.value)
	}
}
