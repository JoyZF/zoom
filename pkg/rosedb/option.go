// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rosedb

import (
	"os"

	"github.com/JoyZF/wal"
)

type Options struct {
	// DirPath specifies the directory path where the WAL segment files will be stored.
	DirPath string

	// SegmentSize specifies the maximum size of each segment file in bytes.
	SegmentSize int64

	// BlockCache specifies the size of the block cache in number of bytes.
	// A block cache is used to store recently accessed data blocks, improving read performance.
	// If BlockCache is set to 0, no block cache will be used.
	BlockCache uint32

	// Sync is whether to synchronize writes through os buffer cache and down onto the actual disk.
	// Setting sync is required for durability of a single write operation, but also results in slower writes.
	//
	// If false, and the machine crashes, then some recent writes may be lost.
	// Note that if it is just the process that crashes (machine does not) then no writes will be lost.
	//
	// In other words, Sync being false has the same semantics as a write
	// system call. Sync being true means write followed by fsync.
	Sync bool

	// BytesPerSync specifies the number of bytes to write before calling fsync.
	BytesPerSync uint32

	// WatchQueueSize the cache length of the watch queue.
	// if the size greater than 0, which means enable the watch.
	WatchQueueSize uint64

	// AutoMergeEnable enable the auto merge.
	// auto merge will be triggered when cron expr is satisfied.
	// cron expression follows the standard cron expression.
	// e.g. "0 0 * * *" means merge at 00:00:00 every day.
	// it also supports seconds optionally.
	// when enable the second field, the cron expression will be like this: "0/10 * * * * *" (every 10 seconds).
	// when auto merge is enabled, the db will be closed and reopened after merge done.
	// do not set this shecule too frequently, it will affect the performance.
	// refer to https://en.wikipedia.org/wiki/Cron
	AutoMergeCronExpr string
}

var DefaultOptions = Options{
	DirPath:           "./data",
	SegmentSize:       1 * wal.GB,
	BlockCache:        0,
	Sync:              false,
	BytesPerSync:      0,
	WatchQueueSize:    0,
	AutoMergeCronExpr: "",
}

var DefaultBatchOptions = BatchOptions{
	Sync:     true,
	ReadOnly: false,
}

func tempDBDir() string {
	dir, _ := os.MkdirTemp("", "rosedb-temp")
	return dir
}
