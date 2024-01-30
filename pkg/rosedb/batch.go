package rosedb

import (
	"github.com/bwmarrin/snowflake"
	"github.com/valyala/bytebufferpool"
	"sync"
	"time"
)

type Batch struct {
	db            *DB
	pendingWrites []*LogRecord
	options       BatchOptions
	mu            sync.Mutex
	committed     bool
	rollback      bool
	batchId       *snowflake.Node
	buffers       []*bytebufferpool.ByteBuffer
}

type BatchOptions struct {
	Sync     bool
	ReadOnly bool
}

func (b *Batch) init(rdOnly bool, sync bool, db *DB) *Batch {
	return nil
}

func (b *Batch) reset() {

}

func (b *Batch) lock() {

}

func (b *Batch) unlock() {

}
func (b *Batch) Put(key []byte, value []byte) error {
	return nil
}
func (b *Batch) PutWithTTL(key []byte, value []byte, ttl time.Duration) error {
	return nil
}

func (b *Batch) Get(key []byte) ([]byte, error) {
	return nil, nil
}

func (b *Batch) Delete(key []byte) error {
	return nil
}
func (b *Batch) Exist(key []byte) (bool, error) {
	return false, nil
}
func (b *Batch) Expire(key []byte, ttl time.Duration) error {
	return nil
}

func (b *Batch) TTL(key []byte) (time.Duration, error) {
	return 0, nil
}

func (b *Batch) Persist(key []byte) error {
	return nil
}

func (b *Batch) Commit() error {
	return nil
}

func (b *Batch) Rollback() error {
	return nil
}
