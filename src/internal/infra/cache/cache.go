package CacheService

import (
	"context"
	"errors"
	"time"
)

type Cache interface {
	Clear(ctx context.Context) error
	Delete(ctx context.Context, key string) error
	Get(ctx context.Context, key string) (any, error)
	Exists(ctx context.Context, key string) (bool, error)
	Refresh(ctx context.Context, key string, ttl time.Duration) error
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
}

var (
	ErrCacheMiss   = errors.New("Cache miss")
	ErrCacheClosed = errors.New("Cache is closed")
)
