// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package setutils

import (
	"github.com/TeaOSLab/EdgeNode/internal/zero"
	"sync"
)

type FixedSet struct {
	maxSize int
	locker  sync.RWMutex

	m    map[interface{}]zero.Zero
	keys []interface{}
}

func NewFixedSet(maxSize int) *FixedSet {
	if maxSize <= 0 {
		maxSize = 1024
	}
	return &FixedSet{
		maxSize: maxSize,
		m:       map[interface{}]zero.Zero{},
	}
}

func (this *FixedSet) Push(item interface{}) {
	this.locker.Lock()
	_, ok := this.m[item]
	if !ok {
		// 是否已满
		if len(this.keys) == this.maxSize {
			var firstKey = this.keys[0]
			this.keys = this.keys[1:]
			delete(this.m, firstKey)
		}

		this.m[item] = zero.New()
		this.keys = append(this.keys, item)
	}
	this.locker.Unlock()
}

func (this *FixedSet) Has(item interface{}) bool {
	this.locker.RLock()
	defer this.locker.RUnlock()

	_, ok := this.m[item]
	return ok
}

func (this *FixedSet) Size() int {
	this.locker.RLock()
	defer this.locker.RUnlock()
	return len(this.keys)
}

func (this *FixedSet) Reset() {
	this.locker.Lock()
	this.m = map[interface{}]zero.Zero{}
	this.keys = nil
	this.locker.Unlock()
}
