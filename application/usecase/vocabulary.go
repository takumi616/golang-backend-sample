package usecase

import (
	"context"

	"github.com/takumi616/golang-backend-sample/domain"
)

type VocabularyUsecase struct {
	Repository VocabularyRepository
}

func NewVocabularyUsecase(repository VocabularyRepository) *VocabularyUsecase {
	return &VocabularyUsecase{
		Repository: repository,
	}
}

func (u *VocabularyUsecase) AddVocabulary(ctx context.Context, vocabulary *domain.Vocabulary) (int64, error) {
	return u.Repository.Insert(ctx, vocabulary)
}

func (u *VocabularyUsecase) FetchVocabularyByNo(ctx context.Context, vocabularyNo int64) (*domain.Vocabulary, error) {
	return u.Repository.SelectByVocabularyNo(ctx, vocabularyNo)
}

func (u *VocabularyUsecase) FetchVocabularyList(ctx context.Context) ([]*domain.Vocabulary, error) {
	return u.Repository.SelectAll(ctx)
}
