package transformer

import (
	"github.com/takumi616/golang-backend-sample/domain"
	"github.com/takumi616/golang-backend-sample/interface/controller/request"
	"github.com/takumi616/golang-backend-sample/interface/controller/response"
)

// http request -> domain model
func ToDomain(req *request.VocabularyReq) *domain.Vocabulary {
	return &domain.Vocabulary{
		Title:    req.Title,
		Meaning:  req.Meaning,
		Sentence: req.Sentence,
	}
}

// domain model -> http response
func ToResponse(vocabulary *domain.Vocabulary) *response.VocabularyRes {
	return &response.VocabularyRes{
		VocabularyNo: vocabulary.VocabularyNo,
		Title:        vocabulary.Title,
		Meaning:      vocabulary.Meaning,
		Sentence:     vocabulary.Sentence,
	}
}
