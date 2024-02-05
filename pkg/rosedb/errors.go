// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rosedb

import "errors"

var (
	ErrKeyIsEmpty      = errors.New("the key is empty")
	ErrKeyNotFound     = errors.New("key not found in database")
	ErrDatabaseIsUsing = errors.New("the database directory is used by another process")
	ErrReadOnlyBatch   = errors.New("the batch is read only")
	ErrBatchCommitted  = errors.New("the batch is committed")
	ErrBatchRolledBack = errors.New("the batch is rolled back")
	ErrDBClosed        = errors.New("the database is closed")
	ErrMergeRunning    = errors.New("the merge operation is running")
	ErrWatchDisabled   = errors.New("the watch is disabled")
)
