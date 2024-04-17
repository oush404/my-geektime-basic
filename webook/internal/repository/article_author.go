package repository

import (
	"context"
	"my-geektime-basic/webook/internal/domain"
)

type AritcleAuthorRepository interface {
	Create(ctx context.Context, art domain.Article) (int64, error)
	Update(ctx context.Context, art domain.Article) error
}
