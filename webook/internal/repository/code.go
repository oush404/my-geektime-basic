package repository

import (
	"context"
	"my-geektime-basic/webook/internal/repository/cache"
)

var ErrCodeVerifyTooMany = cache.ErrCodeVerifyTooMany
var ErrCodeSendTooMany = cache.ErrCodeSendTooMany

type CodeRepository interface {
	Set(ctx context.Context, biz, phone, code string) error
	Verify(ctx context.Context, biz, phone, code string) (bool, error)
}

type CachedCodeReposiotry struct {
	cache cache.CodeCache
}

func (c *CachedCodeReposiotry) Set(ctx context.Context, biz, phone, code string) error {
	return c.cache.Set(ctx, biz, phone, code)
}

func (c *CachedCodeReposiotry) Verify(ctx context.Context, biz, phone, code string) (bool, error) {
	return c.cache.Verify(ctx, biz, phone, code)
}

func NewCodeRepository(c cache.CodeCache) CodeRepository {
	return &CachedCodeReposiotry{
		cache: c,
	}
}
