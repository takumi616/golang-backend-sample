package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/takumi616/golang-backend-sample/interface/controller/helper"
	"github.com/takumi616/golang-backend-sample/interface/controller/request"
	"github.com/takumi616/golang-backend-sample/interface/controller/response"
	"github.com/takumi616/golang-backend-sample/interface/controller/transformer"
)

type VocabularyController struct {
	Usecase VocabularyUsecase
}

func NewVocabularyController(usecase VocabularyUsecase) *VocabularyController {
	return &VocabularyController{
		Usecase: usecase,
	}
}

func (c *VocabularyController) AddVocabulary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Read http request body
	var req request.VocabularyReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.ErrorContext(ctx, "failed to read a request body", slog.String("error", err.Error()))
		helper.WriteResponse(
			ctx, w, http.StatusBadRequest,
			response.ErrorRes{Message: "Invalid request format. Failed to parse JSON."},
		)
		return
	}
	defer r.Body.Close()

	// Validation check
	if err := req.Validate(); err != nil {
		slog.ErrorContext(ctx, "invalid request parameters", slog.String("error", err.Error()))
		helper.WriteResponse(
			ctx, w, http.StatusBadRequest,
			response.ErrorRes{Message: "Invalid input parameters. Please check your request."},
		)
		return
	}

	// Transform the request body into the domain model
	vocabulary := transformer.ToDomain(&req)

	// Execute the application layer logic
	vocabularyNo, err := c.Usecase.AddVocabulary(ctx, vocabulary)
	if err != nil {
		helper.WriteResponse(
			ctx, w, http.StatusInternalServerError,
			response.ErrorRes{Message: "Failed to add the vocabulary due to a server error."},
		)
		return
	}

	// Write a returned result to the response body
	helper.WriteResponse(ctx, w, http.StatusCreated, response.VocabularyNoRes{VocabularyNo: vocabularyNo})
}

func (c *VocabularyController) FetchVocabularyByNo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get the vocabularyNo from the request path
	vocabularyNo, err := strconv.Atoi(r.PathValue("vocabularyNo"))
	if err != nil {
		slog.ErrorContext(ctx, "invalid path value", slog.String("error", err.Error()))
		helper.WriteResponse(
			ctx, w, http.StatusBadRequest,
			response.ErrorRes{Message: "Invalid request path value. Please check your http request path."},
		)
		return
	}

	// Execute the application layer logic
	vocabulary, err := c.Usecase.FetchVocabularyByNo(ctx, int64(vocabularyNo))
	if errors.Is(err, sql.ErrNoRows) {
		helper.WriteResponse(
			ctx, w, http.StatusNotFound,
			response.ErrorRes{Message: "Failed to get the vocabulary since specified data may not be registered."},
		)
		return
	}

	if err != nil {
		helper.WriteResponse(
			ctx, w, http.StatusInternalServerError,
			response.ErrorRes{Message: "Failed to get the vocabulary due to a server error."},
		)
		return
	}

	// Write a returned result to the response body
	helper.WriteResponse(ctx, w, http.StatusOK, transformer.ToResponse(vocabulary))
}

func (c *VocabularyController) FetchVocabularyList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Execute the application layer logic
	vocabularyList, err := c.Usecase.FetchVocabularyList(ctx)
	if err != nil {
		helper.WriteResponse(ctx, w, http.StatusInternalServerError, response.ErrorRes{Message: "Failed to get the vocabularies due to a server error."})
		return
	}

	// Transform the domain model into the response struct
	var vocabularyResponse []*response.VocabularyRes
	for _, vocabulary := range vocabularyList {
		vocabularyResponse = append(vocabularyResponse, transformer.ToResponse(vocabulary))
	}

	// Write a returned result to the response body
	helper.WriteResponse(ctx, w, http.StatusOK, vocabularyResponse)
}

func (c *VocabularyController) UpdateVocabulary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get the vocabularyNo from the request path
	vocabularyNo, err := strconv.Atoi(r.PathValue("vocabularyNo"))
	if err != nil {
		slog.ErrorContext(ctx, "invalid path value", slog.String("error", err.Error()))
		helper.WriteResponse(
			ctx, w, http.StatusBadRequest,
			response.ErrorRes{Message: "Invalid request path value. Please check your http request path."},
		)
		return
	}

	// Read http request body
	var req request.VocabularyReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.ErrorContext(ctx, "failed to read a request body", slog.String("error", err.Error()))
		helper.WriteResponse(
			ctx, w, http.StatusBadRequest,
			response.ErrorRes{Message: "Invalid request format. Failed to parse JSON."},
		)
		return
	}
	defer r.Body.Close()

	// Transform the request body into entity
	vocabulary := transformer.ToDomain(&req)

	// Execute the application layer logic
	updated, err := c.Usecase.UpdateVocabulary(ctx, int64(vocabularyNo), vocabulary)
	if errors.Is(err, sql.ErrNoRows) {
		helper.WriteResponse(
			ctx, w, http.StatusNotFound,
			response.ErrorRes{Message: "Failed to update the vocabulary since specified data may not be registered."},
		)
		return
	}

	if err != nil {
		helper.WriteResponse(
			ctx, w, http.StatusInternalServerError,
			response.ErrorRes{Message: "Failed to get the vocabularies due to a server error."},
		)
		return
	}

	// Write a returned result to the response body
	helper.WriteResponse(ctx, w, http.StatusOK, response.VocabularyNoRes{VocabularyNo: updated})
}
