// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rosedb

import (
	"encoding/binary"

	"github.com/JoyZF/wal"
	"github.com/valyala/bytebufferpool"
)

// LogRecordType is the type of the log record.
type LogRecordType = byte

const (
	// LogRecordNormal is the normal log record type.
	LogRecordNormal LogRecordType = iota
	// LogRecordDeleted is the deleted log record type.
	LogRecordDeleted
	// LogRecordBatchFinished is the batch finished log record type.
	LogRecordBatchFinished
)

// type batchId keySize valueSize expire
//
//	1  +  10  +   5   +   5   +    10  = 31
const maxLogRecordHeaderSize = binary.MaxVarintLen32*2 + binary.MaxVarintLen64*2 + 1

// LogRecord is the log record of the key/value pair.
// It contains the key, the value, the record type and the batch id
// It will be encoded to byte slice and written to the wal.
type LogRecord struct {
	Key     []byte
	Value   []byte
	Type    LogRecordType
	BatchId uint64
	Expire  int64
}

// IsExpired checks whether the log record is expired.
func (lr *LogRecord) IsExpired(now int64) bool {
	return lr.Expire > 0 && lr.Expire <= now
}

// IndexRecord is the index record of the key.
// It contains the key, the record type and the position of the record in the wal.
// Only used in start up to rebuild the index.
type IndexRecord struct {
	key        []byte
	recordType LogRecordType
	position   *wal.ChunkPosition
}

// +-------------+-------------+-------------+--------------+---------------+---------+--------------+
// |    type     |  batch id   |   key size  |   value size |     expire    |  key    |      value   |
// +-------------+-------------+-------------+--------------+---------------+--------+--------------+
//
//	1 byte	      varint            varint       varint         varint        varint      varint
func encodeLogRecord(logRecord *LogRecord, header []byte, buf *bytebufferpool.ByteBuffer) []byte {
	header[0] = logRecord.Type
	var index = 1

	// batch id
	index += binary.PutUvarint(header[index:], logRecord.BatchId)
	// key size
	index += binary.PutVarint(header[index:], int64(len(logRecord.Key)))
	// value size
	index += binary.PutVarint(header[index:], int64(len(logRecord.Value)))
	// expire
	index += binary.PutVarint(header[index:], logRecord.Expire)

	// copy header
	_, _ = buf.Write(header[:index])
	// copy key
	_, _ = buf.Write(logRecord.Key)
	// copy value
	_, _ = buf.Write(logRecord.Value)

	return buf.Bytes()
}

// decodeLogRecord decodes the log record from the given byte slice.
func decodeLogRecord(buf []byte) *LogRecord {
	recordType := buf[0]

	var index uint32 = 1
	// batch id
	batchId, n := binary.Uvarint(buf[index:])
	index += uint32(n)

	// key size
	keySize, n := binary.Varint(buf[index:])
	index += uint32(n)

	// value size
	valueSize, n := binary.Varint(buf[index:])
	index += uint32(n)

	// expire
	expire, n := binary.Varint(buf[index:])
	index += uint32(n)

	// copy key
	key := make([]byte, keySize)
	copy(key[:], buf[index:index+uint32(keySize)])
	index += uint32(keySize)

	// copy value
	value := make([]byte, valueSize)
	copy(value[:], buf[index:index+uint32(valueSize)])

	return &LogRecord{Key: key, Value: value, Expire: expire,
		BatchId: batchId, Type: recordType}
}
