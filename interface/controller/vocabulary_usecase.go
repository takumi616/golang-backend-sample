package controller

import (
	"context"

	"github.com/takumi616/golang-backend-sample/domain"
)

type VocabularyUsecase interface {
	AddVocabulary(ctx context.Context, vocabulary *domain.Vocabulary) (int64, error)

	FetchVocabularyByNo(ctx context.Context, vocabularyNo int64) (*domain.Vocabulary, error)

	FetchVocabularyList(ctx context.Context) ([]*domain.Vocabulary, error)

	UpdateVocabulary(ctx context.Context, vocabularyNo int64, vocabulary *domain.Vocabulary) (int64, error)
}
