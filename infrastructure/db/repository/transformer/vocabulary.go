package transformer

import (
	"github.com/takumi616/golang-backend-sample/domain"
	"github.com/takumi616/golang-backend-sample/infrastructure/db/repository/model"
)

// Domain model -> DB model
func ToModel(vocabulary *domain.Vocabulary) *model.VocabularyInput {
	return &model.VocabularyInput{
		Title:    vocabulary.Title,
		Meaning:  vocabulary.Meaning,
		Sentence: vocabulary.Sentence,
	}
}
