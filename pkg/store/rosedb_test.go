// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	_, err := DB()
	if err != nil {
		panic(err)
	}
}

func TestRoseDB_Delete(t *testing.T) {
	db, _ := DB()
	err := db.Delete([]byte("test_del"))
	assert.Nil(t, err)
}

func TestRoseDB_Get(t *testing.T) {
	db, _ := DB()
	err := db.Put([]byte("test"), []byte("test"))
	assert.Nil(t, err)
	bytes, err := db.Get([]byte("test"))
	assert.Nil(t, err)
	assert.True(t, string(bytes) != "")
}

func TestRoseDB_Put(t *testing.T) {
	db, _ := DB()
	err := db.Put([]byte("test"), []byte("test"))
	assert.Nil(t, err)
}

func TestRoseDB_PutWithTTL(t *testing.T) {
	db, _ := DB()
	err := db.PutWithTTL([]byte("test"), []byte("test"), 60*time.Second)
	assert.Nil(t, err)
}

func TestRoseDB_TTL(t *testing.T) {
	db, _ := DB()
	err := db.PutWithTTL([]byte("test"), []byte("test"), 60*time.Second)
	assert.Nil(t, err)
	ttl, err := db.TTL([]byte("test"))
	assert.True(t, ttl.Seconds() > 0)
}
