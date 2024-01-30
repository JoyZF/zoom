package rosedb

type Options struct {
	DirPath        string
	SegmentSize    int64
	BlockCache     uint32
	Sync           bool
	BytesPerSync   uint32
	WatchQueueSize int64
}
