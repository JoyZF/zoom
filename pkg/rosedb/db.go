package rosedb

import (
	"github.com/JoyZF/wal"
	"github.com/JoyZF/zoom/pkg/rosedb/flock"
	"github.com/JoyZF/zoom/pkg/rosedb/index"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

const (
	dataFileNameSuffix = ".SEG"
)

type DB struct {
	dataFiles        *wal.WAL //  data files are a sets of segment files in WAL.
	hintFile         *wal.WAL //  hint file is used to store the key and the position for fast startup.
	index            index.Indexer
	options          Options
	fileLock         *flock.Flock
	mu               sync.RWMutex
	closed           bool
	mergeRunning     uint32 // indicate if the database is merging
	batchPool        sync.Pool
	recordPool       sync.Pool
	encodeHeader     []byte
	watchCh          chan *Event // user consume channel for watch events
	watcher          *Watcher
	expiredCursorKey []byte     // the location to which DeleteExpiredKeys executes.
	cronScheduler    *cron.Cron // cron scheduler for auto merge task
}

type Stat struct {
	KeysNum  int64
	DiskSize int64
}

func Open(options Options) (*DB, error) {
	return nil, nil
}

// openWalFiles return a *wal.WAL, error
func (db *DB) openWalFiles() (*wal.WAL, error) {
	return wal.Open(wal.Options{
		DirPath:        db.options.DirPath,
		SegmentSize:    db.options.SegmentSize,
		SegmentFileExt: dataFileNameSuffix,
		BlockCache:     db.options.BlockCache,
		Sync:           db.options.Sync,
		BytesPerSync:   db.options.BytesPerSync,
	})
}

func (db *DB) loadIndex() error {
	// TODO load index from hintFile
	// TODO load index from dataFiles
	return nil
}

func (db *DB) Close() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.closed {
		return nil
	}

	// close file
	if err := db.closeFiles(); err != nil {
		return err
	}

	// TODO free FLOCK
	// release file lock
	if err := db.fileLock.Unlock(); err != nil {
		return err
	}
	// close watchCh
	if db.options.WatchQueueSize > 0 {
		close(db.watchCh)
	}
	//  close cron scheduler
	if db.cronScheduler != nil {
		db.cronScheduler.Stop()
	}
	db.closed = true
	return nil
}

// closeFiles close all data files and hint file
func (db *DB) closeFiles() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if err := db.dataFiles.Close(); err != nil {
		return err
	}

	if db.hintFile != nil {
		if err := db.hintFile.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Sync all data files to the underlying storage.
func (db *DB) Sync() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	return db.dataFiles.Sync()
}

// Stat returns the statistics of the database.
func (db *DB) Stat() *Stat {
	db.mu.Lock()
	defer db.mu.Unlock()

	return &Stat{}
}

func (db *DB) Put(key []byte, value []byte) error {
	return nil
}

func (db *DB) PutWithTTL(key []byte, value []byte, ttl time.Duration) error {
	return nil
}

func (db *DB) Get(key []byte) ([]byte, error) {
	return nil, nil
}

func (db *DB) Delete(key []byte) error {
	return nil
}

func (db *DB) Exist(key []byte) (bool, error) {
	return false, nil
}

func (db *DB) Expire(key []byte, ttl time.Duration) error {
	return nil
}

func (db *DB) TTL(key []byte) (time.Duration, error) {
	return 0, nil
}

func (db *DB) Persist(key []byte) error {
	return nil
}

func (db *DB) Watch() (<-chan *Event, error) {
	if db.options.WatchQueueSize <= 0 {
		return nil, ErrWatchDisabled
	}
	return db.watchCh, nil
}
