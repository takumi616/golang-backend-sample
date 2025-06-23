package request

import (
	"errors"
	"fmt"
	"strings"
)

type VocabularyReq struct {
	Title    string `json:"title"`
	Meaning  string `json:"meaning"`
	Sentence string `json:"sentence"`
}

func (r *VocabularyReq) Validate() error {
	r.Title = strings.TrimSpace(r.Title)
	r.Meaning = strings.TrimSpace(r.Meaning)
	r.Sentence = strings.TrimSpace(r.Sentence)

	if r.Title == "" {
		return errors.New("title is required")
	}
	if len(r.Title) > 20 {
		return fmt.Errorf("title must be 20 characters or fewer")
	}

	if r.Meaning == "" {
		return errors.New("meaning is required")
	}

	if r.Sentence == "" {
		return errors.New("sentence is required")
	}

	return nil
}
