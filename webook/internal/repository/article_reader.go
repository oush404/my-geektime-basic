package repository

import (
	"context"
	"my-geektime-basic/webook/internal/domain"
)

type ArticleReaderRepository interface {
	Save(ctx context.Context, art domain.Article) error
}
