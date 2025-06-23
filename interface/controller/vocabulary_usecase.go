package controller

import (
	"context"

	"github.com/takumi616/golang-backend-sample/domain"
)

type VocabularyUsecase interface {
	AddVocabulary(ctx context.Context, vocabulary *domain.Vocabulary) (int64, error)
}
