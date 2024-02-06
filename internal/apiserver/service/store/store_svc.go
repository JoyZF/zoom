package store

import (
	"context"
	v1 "github.com/JoyZF/zoom/internal/apiserver/types/v1"
	"github.com/JoyZF/zoom/pkg/store"
	"time"
)

type Store struct {
}

func NewStore() Store {
	return Store{}
}

func (s Store) Get(ctx context.Context, key string) (string, error) {
	val, err := store.GetStore().Get([]byte(key))
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func (s Store) Put(ctx context.Context, req *v1.StorePutReq) error {
	return store.GetStore().Put([]byte(req.Key), []byte(req.Value))
}

func (s Store) Delete(ctx context.Context, req *v1.KeyReq) error {
	return store.GetStore().Delete([]byte(req.Key))
}

func (s Store) TTL(ctx context.Context, req *v1.KeyReq) (int64, error) {
	ttl, err := store.GetStore().TTL([]byte(req.Key))
	if err != nil {
		return 0, err
	}
	t := ttl.Milliseconds() / 1e3
	return t, nil
}

func (s Store) PutWithTTL(ctx context.Context, req *v1.StorePutWithTTLReq) error {
	return store.GetStore().
		PutWithTTL([]byte(req.Key), []byte(req.Value), time.Duration(req.TTL)*time.Second)
}

func (s Store) Sync(ctx context.Context) error {
	return store.GetStore().Sync()
}

func (s Store) Stat(ctx context.Context) any {
	return store.GetStore().Stat()
}

func (s Store) Exist(ctx context.Context, key string) (bool, error) {
	return store.GetStore().Exist([]byte(key))
}

func (s Store) Expire(ctx context.Context, req *v1.ExpireReq) error {
	return store.GetStore().Expire([]byte(req.Key), time.Duration(req.TTL)*time.Second)
}
