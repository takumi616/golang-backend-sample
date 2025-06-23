package model

type VocabularyInput struct {
	Title    string
	Meaning  string
	Sentence string
}

type VocabularyOutput struct {
	VocabularyNo int64
	Title        string
	Meaning      string
	Sentence     string
}
