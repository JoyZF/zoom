// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

import (
	"sync"
	"time"

	"github.com/JoyZF/zoom/pkg/rosedb"
)

type RoseDB struct {
	DB *rosedb.DB
}

var (
	roseDB *RoseDB
	once   sync.Once
)

func NewRoseDB() (*RoseDB, error) {
	var err error
	once.Do(func() {
		var db *rosedb.DB
		db, err = rosedb.Open(rosedb.DefaultOptions)
		if err != nil {
			return
		}
		roseDB = &RoseDB{DB: db}
	})

	if err != nil {
		return nil, err
	}
	return roseDB, nil
}

func (r *RoseDB) Get(key []byte) ([]byte, error) {
	return r.DB.Get(key)
}

func (r *RoseDB) Put(key, value []byte) error {
	return r.DB.Put(key, value)
}

func (r *RoseDB) PutWithTTL(key, value []byte, ttl time.Duration) error {
	return r.DB.PutWithTTL(key, value, ttl)
}

func (r *RoseDB) Delete(key []byte) error {
	return r.DB.Delete(key)
}

func (r *RoseDB) TTL(key []byte) (time.Duration, error) {
	return r.DB.TTL(key)
}

func (r *RoseDB) Sync() error {
	return r.DB.Sync()
}

func (r *RoseDB) Stat() any {
	return r.DB.Stat()
}

func (r *RoseDB) Exist(key []byte) (bool, error) {
	return r.DB.Exist(key)
}

func (r *RoseDB) Expire(key []byte, ttl time.Duration) error {
	return r.DB.Expire(key, ttl)
}
