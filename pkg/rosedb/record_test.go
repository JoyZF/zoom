package rosedb

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/bytebufferpool"
	"testing"
)

func Test_encodeLogRecord(t *testing.T) {
	buf := bytebufferpool.ByteBuffer{
		B: []byte{},
	}
	bytes := encodeLogRecord(&LogRecord{
		Key:     []byte("this is key"),
		Value:   []byte("this is value"),
		Type:    LogRecordDeleted,
		BatchId: 1,
		Expire:  0,
	}, []byte("header"), &buf)

	logRecord := decodeLogRecord(bytes)
	fmt.Println(string(logRecord.Key))
	b, _ := json.Marshal(logRecord)
	fmt.Println(string(b))
	var temp LogRecord
	_ = json.Unmarshal(b, &temp)
	fmt.Println(temp)
}
