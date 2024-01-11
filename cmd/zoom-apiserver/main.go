package main

import (
	"github.com/JoyZF/zoom/internal/apiserver"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	apiserver.NewApp("zoom-apiserver").Run()
}
