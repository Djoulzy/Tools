// Copyright (c) 2015-2017 Marin Atanasov Nikolov <dnaeon@gmail.com>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//
//  1. Redistributions of source code must retain the above copyright
//     notice, this list of conditions and the following disclaimer
//     in this position and unchanged.
//  2. Redistributions in binary form must reproduce the above copyright
//     notice, this list of conditions and the following disclaimer in the
//     documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE AUTHOR(S) ``AS IS'' AND ANY EXPRESS OR
// IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
// OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE AUTHOR(S) BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
// THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package cmap

import "sync"

// ConcurrentMap is a map type that can be safely shared between
// goroutines that require read/write access to a map
type CMap struct {
	sync.RWMutex
	items map[string]interface{}
}

// ConcurrentMapItem contains a key/value pair item of a concurrent map
type CMapItem struct {
	Key   string
	Value interface{}
}

// NewConcurrentMap creates a new concurrent map
func NewCMap() *CMap {
	cm := &CMap{
		items: make(map[string]interface{}),
	}
	return cm
}

// Set adds an item to a concurrent map
func (cm *CMap) Set(key string, value interface{}) {
	cm.Lock()
	cm.items[key] = value
	cm.Unlock()
}

// Get retrieves the value for a concurrent map item
func (cm *CMap) Get(key string) (interface{}, bool) {
	cm.Lock()
	value, ok := cm.items[key]
	cm.Unlock()
	return value, ok
}

// Delete supress an entry in the map
func (cm *CMap) Delete(key string) {
	cm.Lock()
	delete(cm.items, key)
	cm.Unlock()
}

// Length return the size of the map
func (cm *CMap) Length() int {
	return len(cm.items)
}

// Iter iterates over the items in a concurrent map
// Each item is sent over a channel, so that
// we can iterate over the map using the builtin range keyword
func (cm *CMap) Iter() <-chan CMapItem {
	c := make(chan CMapItem)
	f := func() {
		cm.Lock()
		for k, v := range cm.items {
			c <- CMapItem{k, v}
		}
		close(c)
		cm.Unlock()
	}
	go f()
	return c
}
