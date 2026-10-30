// /home/krylon/go/src/github.com/blicero/krylib/syncmap.go
// -*- mode: go; coding: utf-8; -*-
// Created on 30. 10. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-10-30 11:11:23 krylon>

package krylib

import "sync"

// SyncMap is a map-like container (it does, unsurprisingly, use a map as its
// underlying storage) that has all accesses synchronized through a sync.RWMutex.
type SyncMap[K comparable, V any] struct {
	dict map[K]V
	lock sync.RWMutex
}

// NewSyncMap creates and returns a fresh instance.
func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	var sm = &SyncMap[K, V]{
		dict: make(map[K]V),
	}

	return sm
} // func NewSyncMap[K any, V any]() *SyncMap

// Clear removes all entries from the map.
func (sm *SyncMap[K, V]) Clear() {
	sm.lock.Lock()
	clear(sm.dict)
	sm.lock.Unlock()
} // func (sm *SyncMap[K, V]) Clear()

// Delete remove the entry for the given key. If the key is not present, it is
// a no-op.
func (sm *SyncMap[K, V]) Delete(key K) {
	sm.lock.Lock()
	delete(sm.dict, key)
	sm.lock.Unlock()
} // func (sm *SyncMap[K, V]) Delete(key K)

// Get attempts to load the value for the given key.
func (sm *SyncMap[K, V]) Get(key K) (V, bool) {
	sm.lock.RLock()
	val, ok := sm.dict[key]
	sm.lock.RUnlock()

	return val, ok
} // func (sm *SyncMap[K, V]) Get(key K) (V, bool)

// Put stores the given value under the given key.
func (sm *SyncMap[K, V]) Put(key K, val V) {
	sm.lock.Lock()
	sm.dict[key] = val
	sm.lock.Unlock()
} // func (sm *SyncMap[K, V]) Put(key K, val V)

// Map creates a shallow copy of the underlying map and returns it.
func (sm *SyncMap[K, V]) Map() map[K]V {
	var dict = make(map[K]V)
	sm.lock.RLock()

	for k, v := range sm.dict {
		dict[k] = v
	}

	sm.lock.RUnlock()

	return dict
} // func (sm *SyncMap[K, V])Map() map[K]V
