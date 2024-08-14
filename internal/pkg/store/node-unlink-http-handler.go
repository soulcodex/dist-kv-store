package store

import (
	"github.com/rs/zerolog"
	"net/http"

	"codesignal/internal/pkg/utils"
)

type NodeUnlinkRequest struct {
	Id string `json:"id"`
}

func NodeRaftUnlinkHttpHandler(server Consensus, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var unlinkRequest NodeUnlinkRequest
		err := utils.JsonBodyToStruct[NodeUnlinkRequest](req, &unlinkRequest)
		if err != nil {
			utils.WriteHttpBadRequestResponse(req.Context(), map[string]interface{}{}, log, w)
			return
		}

		unlinkErr := server.Unlink(unlinkRequest.Id)

		switch unlinkErr.(type) {
		case nil:
			response := map[string]interface{}{"message": "node unlinked successfully"}
			utils.WriteHttpOkResponse(req.Context(), response, log, w)
			break
		default:
			utils.WriteHttpInternalServerError(req.Context(), unlinkErr, log, w)
		}
	}
}
