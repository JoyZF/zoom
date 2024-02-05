// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rosedb

import (
	"sync"
	"time"
)

type WatchActionType = byte

const (
	WatchActionPut WatchActionType = iota
	WatchActionDelete
)

type Event struct {
	Action  WatchActionType
	Key     []byte
	Value   []byte
	BatchId uint64
}

type Watcher struct {
	queue eventQueue
	mu    sync.RWMutex
}

type eventQueue struct {
	Events   []*Event
	Capacity uint64
	Front    uint64 // read point
	Back     uint64 // write point
}

func (q *eventQueue) push(e *Event) {
	q.Events[q.Back] = e
	q.Back = (q.Back + 1) % q.Capacity
}

func (q *eventQueue) pop() *Event {
	e := q.Events[q.Front]
	q.frontTakeAStep()
	return e
}

func (q *eventQueue) frontTakeAStep() {
	q.Front = (q.Front + 1) % q.Capacity
}

func (q *eventQueue) isFull() bool {
	return (q.Back+1)%q.Capacity == q.Front
}

func (q *eventQueue) isEmpty() bool {
	return q.Back == q.Front
}

func NewWatcher(capacity uint64) *Watcher {
	return &Watcher{
		queue: eventQueue{
			Events:   make([]*Event, capacity),
			Capacity: capacity,
		},
	}
}

func (w *Watcher) putEvent(e *Event) {
	w.mu.Lock()
	w.queue.push(e)
	if w.queue.isFull() {
		w.queue.frontTakeAStep()
	}
	w.mu.Unlock()
}

// getEvent if queue is empty, it will return nil.
func (w *Watcher) getEvent() *Event {
	w.mu.RLock()
	defer w.mu.RUnlock()
	if w.queue.isEmpty() {
		return nil
	}
	return w.queue.pop()
}

// sendEvent send events to DB's watch
func (w *Watcher) sendEvent(c chan *Event) {
	for {
		event := w.getEvent()
		if event == nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		c <- event
	}
}
