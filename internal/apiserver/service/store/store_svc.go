package store

import (
	"context"
	v1 "github.com/JoyZF/zoom/internal/apiserver/types/v1"
	"github.com/JoyZF/zoom/pkg"
	"time"
)

type Store struct {
}

func NewStore() Store {
	return Store{}
}

func (s Store) Get(ctx context.Context, key string) (string, error) {
	store := pkg.GetStore()
	val, err := store.Get([]byte(key))
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func (s Store) Put(ctx context.Context, req *v1.StorePutReq) error {
	store := pkg.GetStore()
	return store.Put([]byte(req.Key), []byte(req.Value))
}

func (s Store) Delete(ctx context.Context, req *v1.StoreGetReq) error {
	store := pkg.GetStore()
	return store.Delete([]byte(req.Key))
}

func (s Store) TTL(ctx context.Context, req *v1.StoreGetReq) (int64, error) {
	store := pkg.GetStore()
	ttl, err := store.TTL([]byte(req.Key))
	if err != nil {
		return 0, err
	}
	t := ttl.Milliseconds() / 1e3
	return t, nil
}

func (s Store) PutWithTTL(ctx context.Context, req *v1.StorePutWithTTLReq) error {
	store := pkg.GetStore()
	return store.PutWithTTL([]byte(req.Key), []byte(req.Value), time.Duration(req.TTL)*time.Second)
}
