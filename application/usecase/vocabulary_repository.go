package usecase

import (
	"context"

	"github.com/takumi616/golang-backend-sample/domain"
)

type VocabularyRepository interface {
	Insert(ctx context.Context, vocabulary *domain.Vocabulary) (int64, error)
}
