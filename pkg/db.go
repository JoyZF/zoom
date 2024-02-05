// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pkg

import (
	"sync"
	"time"

	"github.com/JoyZF/errors"
)

type DBer interface {
	Get(key []byte) ([]byte, error)
	Put(key, value []byte) error
	PutWithTTL(key []byte, value []byte, ttl time.Duration) error
	Delete(key []byte) error
	TTL(key []byte) (time.Duration, error)
}

var (
	ErrStoreNotFound = errors.New("store not found")
)

var (
	db     DBer
	dbOnce sync.Once
)

func DB(driver string) (err error) {
	dbOnce.Do(func() {
		switch driver {
		case "ROSEDB":
			db, err = NewRoseDB()
		default:
			err = ErrStoreNotFound
		}
		return
	})
	return
}

// GetStore return a DBer
func GetStore() DBer {
	return db
}
