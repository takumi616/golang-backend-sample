package main

import (
	"context"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"github.com/takumi616/golang-backend-sample/application/usecase"
	"github.com/takumi616/golang-backend-sample/config"
	"github.com/takumi616/golang-backend-sample/infrastructure/db"
	"github.com/takumi616/golang-backend-sample/infrastructure/db/repository"
	"github.com/takumi616/golang-backend-sample/infrastructure/web"
	"github.com/takumi616/golang-backend-sample/interface/controller"
)

func run(ctx context.Context) error {
	// Initialize a structure logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Get the environment variables
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		return err
	}

	// Open the DB
	db, err := db.Open(ctx, cfg)
	if err != nil {
		return err
	}

	//Set up dependencies between layers
	vocabularyRepository := repository.NewVocabularyRepository(db)
	vocabularyUsecase := usecase.NewVocabularyUsecase(vocabularyRepository)
	vocabularyController := controller.NewVocabularyController(vocabularyUsecase)

	// Register the handlers
	serveMux := web.NewServeMux(vocabularyController)
	mux := serveMux.RegisterHandler()

	// Run the http server
	server := web.NewServer(cfg.Port, mux)
	if err = server.Run(ctx); err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		slog.ErrorContext(ctx, "Golang application could not start", "err", err)
	}
}
