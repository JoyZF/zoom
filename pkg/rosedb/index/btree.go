// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package index

import (
	"bytes"
	"sync"

	"github.com/JoyZF/wal"
	"github.com/google/btree"
)

// MemoryBTree is a memory based btree implementation of the Index interface
// It is a wrapper around the google/btree package: github.com/google/btree
type MemoryBTree struct {
	tree *btree.BTree
	lock *sync.RWMutex
}

func (mt *MemoryBTree) Put(key []byte, position *wal.ChunkPosition) *wal.ChunkPosition {
	mt.lock.Lock()
	defer mt.lock.Unlock()

	i := &item{key: key, pos: position}
	oldValue := mt.tree.ReplaceOrInsert(i)
	if oldValue != nil {
		return oldValue.(*item).pos
	}
	return nil
}

func (mt *MemoryBTree) Get(key []byte) *wal.ChunkPosition {
	mt.lock.RLock()
	defer mt.lock.RUnlock()

	val := mt.tree.Get(&item{key: key})
	if val != nil {
		return val.(*item).pos
	}
	return nil
}

func (mt *MemoryBTree) Delete(key []byte) (*wal.ChunkPosition, bool) {
	mt.lock.Lock()
	defer mt.lock.Unlock()

	i := mt.tree.Delete(&item{key: key})
	if i != nil {
		return i.(*item).pos, true
	}
	return nil, false
}

func (mt *MemoryBTree) Size() int {
	return mt.tree.Len()
}

func (mt *MemoryBTree) Ascend(handleFn func(key []byte, position *wal.ChunkPosition) (bool, error)) {
	mt.lock.RLock()
	defer mt.lock.RUnlock()

	mt.tree.Ascend(func(i btree.Item) bool {
		cont, err := handleFn(i.(*item).key, i.(*item).pos)
		if err != nil {
			return false
		}
		return cont
	})
}

func (mt *MemoryBTree) AscendRange(
	startKey, endKey []byte,
	handleFn func(key []byte, position *wal.ChunkPosition) (bool, error),
) {
	mt.lock.RLock()
	defer mt.lock.RUnlock()

	mt.tree.Descend(func(i btree.Item) bool {
		cont, err := handleFn(i.(*item).key, i.(*item).pos)
		if err != nil {
			return false
		}
		return cont
	})
}

func (mt *MemoryBTree) AscendGreaterOrEqual(
	key []byte,
	handleFn func(key []byte, position *wal.ChunkPosition) (bool, error),
) {
	mt.lock.RLock()
	defer mt.lock.RUnlock()

	mt.tree.AscendGreaterOrEqual(&item{key: key}, func(i btree.Item) bool {
		cont, err := handleFn(i.(*item).key, i.(*item).pos)
		if err != nil {
			return false
		}
		return cont
	})
}

func (mt *MemoryBTree) Descend(handleFn func(key []byte, pos *wal.ChunkPosition) (bool, error)) {
	mt.lock.RLock()
	defer mt.lock.RUnlock()

	mt.tree.Descend(func(i btree.Item) bool {
		cont, err := handleFn(i.(*item).key, i.(*item).pos)
		if err != nil {
			return false
		}
		return cont
	})
}

func (mt *MemoryBTree) DescendRange(
	startKey, endKey []byte,
	handleFn func(key []byte, position *wal.ChunkPosition) (bool, error),
) {
	mt.lock.RLock()
	defer mt.lock.RUnlock()

	mt.tree.DescendRange(&item{key: startKey}, &item{key: endKey}, func(i btree.Item) bool {
		cont, err := handleFn(i.(*item).key, i.(*item).pos)
		if err != nil {
			return false
		}
		return cont
	})
}

func (mt *MemoryBTree) DescendLessOrEqual(
	key []byte,
	handleFn func(key []byte, position *wal.ChunkPosition) (bool, error),
) {
	mt.lock.RLock()
	defer mt.lock.RUnlock()

	mt.tree.DescendLessOrEqual(&item{key: key}, func(i btree.Item) bool {
		cont, err := handleFn(i.(*item).key, i.(*item).pos)
		if err != nil {
			return false
		}
		return cont
	})
}

type item struct {
	key []byte
	pos *wal.ChunkPosition
}

func (it *item) Less(bi btree.Item) bool {
	return bytes.Compare(it.key, bi.(*item).key) < 0
}

func newBTree() *MemoryBTree {
	return &MemoryBTree{
		tree: btree.New(32), // 默认深度32
		lock: new(sync.RWMutex),
	}
}
