// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirSize(t *testing.T) {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	dirSize, err := DirSize(dir)
	fmt.Println(dirSize)
	assert.Nil(t, err)
	assert.True(t, dirSize > 0)
}
