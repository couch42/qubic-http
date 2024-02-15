package handlers

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qubic/qubic-http/external/opensearch"
	"github.com/qubic/qubic-http/foundation/web"
	"net/http"
)

type optionsHandler struct {
	opensearchClient *opensearch.Client
}

func (h *optionsHandler) SetCorsHeader(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status, err := h.opensearchClient.GetStatus(ctx)
	if err != nil {
		return errors.Wrap(err, "getting status from opensearch")
	}

	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	return web.Respond(ctx, w, status, http.StatusOK)
}
