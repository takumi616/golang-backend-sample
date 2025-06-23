package repository

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/takumi616/golang-backend-sample/domain"
	"github.com/takumi616/golang-backend-sample/infrastructure/db/repository/transformer"
)

type VocabularyRepository struct {
	Db *sql.DB
}

func NewVocabularyRepository(db *sql.DB) *VocabularyRepository {
	return &VocabularyRepository{
		Db: db,
	}
}

func (r *VocabularyRepository) Insert(ctx context.Context, vocabulary *domain.Vocabulary) (int64, error) {
	// Transform the received domain model into a DB model
	vocabModel := transformer.ToModel(vocabulary)

	// Begin a transaction
	tx, err := r.Db.BeginTx(ctx, nil)
	if err != nil {
		slog.ErrorContext(ctx, "failed to begin a transaction", "error", err)
		return 0, err
	}
	defer tx.Rollback()

	// Execute an insert process
	var vocabularyNo int64
	err = tx.QueryRowContext(
		ctx,
		"INSERT INTO vocabularies(title, meaning, sentence) VALUES($1, $2, $3) ON CONFLICT (title) DO NOTHING RETURNING vocabulary_no",
		vocabModel.Title, vocabModel.Meaning, vocabModel.Sentence,
	).Scan(&vocabularyNo)

	// sql.ErrNoRows is returned when an insert proccess is skipped with ON CONFLICT DO NOTHING
	if err == sql.ErrNoRows {
		slog.ErrorContext(ctx, "duplicate vocabulary detected", "error", err, "title", vocabModel.Title)
		return 0, err
	}

	if err != nil {
		slog.ErrorContext(ctx, "failed to insert a vocabulary", "error", err)
		return 0, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		slog.ErrorContext(ctx, "failed to commit the transaction", "error", err)
		return 0, err
	}

	slog.InfoContext(ctx, "new vocabulary was inserted successfully", "vocabulary_no", vocabularyNo)
	return vocabularyNo, nil
}
