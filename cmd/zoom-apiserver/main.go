// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"math/rand"
	"time"

	"github.com/JoyZF/zoom/internal/apiserver"
)

//	@title			zoom-api-server API
//	@version		1.0
//	@description	a kv store by web api
//	@license.name	Apache 2.0
//	@contact.name	joyssss94@gmail.com
//	@contact.url	https://github.com/JoyZF/zoom
//	@host			localhost:8080
//	@BasePath		/v1
func main() {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	apiserver.NewApp("zoom-api-server").Run()
}
