package store

import (
	"github.com/rs/zerolog"
	"maps"
	"net/http"

	"codesignal/internal/pkg/utils"
)

func NodeRaftStatsHttpHandler(server Consensus, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		response := server.Stats()
		utils.WriteHttpOkResponse(req.Context(), maps.Clone(response), log, w)
	}
}
