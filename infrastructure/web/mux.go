package web

import (
	"net/http"

	"github.com/takumi616/golang-backend-sample/interface/controller"
)

type ServeMux struct {
	VocabularyController *controller.VocabularyController
}

func NewServeMux(vocabularyController *controller.VocabularyController) *ServeMux {
	return &ServeMux{
		VocabularyController: vocabularyController,
	}
}

func (s *ServeMux) RegisterHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/vocabularies", s.VocabularyController.AddVocabulary)
	mux.HandleFunc("GET /api/vocabularies/{vocabularyNo}", s.VocabularyController.FetchVocabularyByNo)
	mux.HandleFunc("GET /api/vocabularies", s.VocabularyController.FetchVocabularyList)
	mux.HandleFunc("PUT /api/vocabularies/{vocabularyNo}", s.VocabularyController.UpdateVocabulary)
	mux.HandleFunc("DELETE /api/vocabularies/{vocabularyNo}", s.VocabularyController.DeleteVocabulary)

	return mux
}
