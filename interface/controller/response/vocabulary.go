package response

type VocabularyNoRes struct {
	VocabularyNo int64 `json:"vocabulary_no"`
}

type VocabularyRes struct {
	VocabularyNo int64  `json:"vocabulary_no"`
	Title        string `json:"title"`
	Meaning      string `json:"meaning"`
	Sentence     string `json:"sentence"`
}
