package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/takumi616/golang-backend-sample/domain"
	"github.com/takumi616/golang-backend-sample/infrastructure/db/repository/model"
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
		slog.ErrorContext(ctx, "failed to begin a transaction", slog.String("error", err.Error()))
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
	if errors.Is(err, sql.ErrNoRows) {
		slog.WarnContext(
			ctx, "duplicate vocabulary detected", slog.Int64("vocabularyNo", vocabularyNo), slog.String("error", err.Error()),
		)
		return 0, err
	}

	if err != nil {
		slog.ErrorContext(ctx, "failed to insert a vocabulary", slog.String("error", err.Error()))
		return 0, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		slog.ErrorContext(ctx, "failed to commit the transaction", slog.String("error", err.Error()))
		return 0, err
	}

	slog.InfoContext(ctx, "new vocabulary was inserted successfully", slog.Int64("vocabularyNo", vocabularyNo))
	return vocabularyNo, nil
}

func (r *VocabularyRepository) SelectByVocabularyNo(ctx context.Context, vocabularyNo int64) (*domain.Vocabulary, error) {
	// Execute a select process
	var row model.VocabularyOutput
	err := r.Db.QueryRowContext(
		ctx,
		"SELECT vocabulary_no, title, meaning, sentence FROM vocabularies WHERE vocabulary_no = $1",
		vocabularyNo,
	).Scan(&row.VocabularyNo, &row.Title, &row.Meaning, &row.Sentence)

	// Not found by specified vocabularyNo
	if errors.Is(err, sql.ErrNoRows) {
		slog.WarnContext(
			ctx, "no vocabulary found", slog.Int64("vocabularyNo", vocabularyNo), slog.String("error", err.Error()),
		)
		return nil, err
	}

	if err != nil {
		slog.ErrorContext(
			ctx, "failed to query vocabulary", slog.Int64("vocabularyNo", vocabularyNo), slog.String("error", err.Error()),
		)
		return nil, err
	}

	// Transform the selected row into domain model
	vocabulary := transformer.ToDomain(&row)

	slog.InfoContext(ctx, "vocabulary fetched successfully", slog.Int64("vocabularyNo", vocabularyNo))
	return vocabulary, nil
}

func (r *VocabularyRepository) SelectAll(ctx context.Context) ([]*domain.Vocabulary, error) {
	// Execute a select process
	rows, err := r.Db.QueryContext(
		ctx, "SELECT vocabulary_no, title, meaning, sentence FROM vocabularies ORDER BY vocabulary_no ASC",
	)
	if err != nil {
		slog.ErrorContext(ctx, "failed to query all vocabularies", slog.String("error", err.Error()))
		return nil, err
	}
	defer rows.Close()

	// Copy the selected columns into the domain model
	var vocabularyList []*domain.Vocabulary
	for rows.Next() {
		var vocabulary model.VocabularyOutput
		if err := rows.Scan(&vocabulary.VocabularyNo, &vocabulary.Title, &vocabulary.Meaning, &vocabulary.Sentence); err != nil {
			slog.ErrorContext(ctx, "failed to scan vocabulary row", slog.String("error", err.Error()))
			return nil, err
		}
		vocabularyList = append(vocabularyList, transformer.ToDomain(&vocabulary))
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		slog.ErrorContext(ctx, "row iteration error", slog.String("error", err.Error()))
		return nil, err
	}

	slog.InfoContext(ctx, "all vocabularies were fetched successfully", slog.Int("count", len(vocabularyList)))
	return vocabularyList, nil
}
