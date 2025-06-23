package helper

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

func WriteResponse(ctx context.Context, w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		slog.ErrorContext(ctx, "failed to write an response", "err", err)
	}
}
