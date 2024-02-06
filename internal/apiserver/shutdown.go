package apiserver

import "fmt"

type SyncDB struct {
}

func (s SyncDB) OnShutdown(s2 string) error {
	fmt.Println(s2)
	fmt.Println(12312312213213)
	return nil
}
