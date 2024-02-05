// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTestKey(t *testing.T) {
	for i := 0; i < 10; i++ {
		assert.NotNil(t, string(GetTestKey(i)))
	}
}

func TestRandomValue(t *testing.T) {
	for i := 0; i < 10; i++ {
		assert.NotNil(t, string(RandomValue(10)))
	}
}
