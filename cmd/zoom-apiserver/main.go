// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"math/rand"
	"time"

	"github.com/JoyZF/zoom/internal/apiserver"
)

func main() {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	apiserver.NewApp("zoom-api-server").Run()
}
