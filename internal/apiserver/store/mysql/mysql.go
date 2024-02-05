// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"context"
	"fmt"
	"sync"

	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/db"
	"gorm.io/gorm"

	"github.com/JoyZF/zoom/internal/pkg/code"
	"github.com/JoyZF/zoom/internal/pkg/options"
)

type datastore struct {
	db *gorm.DB

	// can include two database instance if needed
	// docker *grom.DB
	// db *gorm.DB
}

func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}

	return db.Close()
}

var (
	DbIns *gorm.DB
	once  sync.Once
)

// GetMySQLClient create mysql client with the given config.
func GetMySQLClient(opts *options.MySQLOptions) (*gorm.DB, error) {
	if opts == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory")
	}

	var err error
	once.Do(func() {
		options := &db.Options{
			Host:                  opts.Host,
			Username:              opts.Username,
			Password:              opts.Password,
			Database:              opts.Database,
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			LogLevel:              opts.LogLevel,
		}
		DbIns, err = db.New(options)

		// uncomment the following line if you need auto migration the given models
		// not suggested in production environment.
		// migrateDatabase(dbIns)

	})

	if err != nil {
		return nil, code.ErrorWithCode(context.Background(), code.SystemError)
	}

	return DbIns, nil
}
