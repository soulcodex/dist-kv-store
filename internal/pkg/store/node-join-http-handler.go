package store

import (
	"github.com/rs/zerolog"
	"net/http"

	"codesignal/internal/pkg/utils"
)

type NodeJoinRequest struct {
	Id      string `json:"id"`
	Address string `json:"address"`
}

func NodeRaftJoinHttpHandler(server Consensus, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var joinRequest NodeJoinRequest
		err := utils.JsonBodyToStruct[NodeJoinRequest](req, &joinRequest)
		if err != nil {
			utils.WriteHttpBadRequestResponse(req.Context(), map[string]interface{}{}, log, w)
			return
		}

		joinErr := server.Join(joinRequest.Id, joinRequest.Address)

		switch joinErr.(type) {
		case nil:
			response := map[string]interface{}{"message": "node joined successfully"}
			utils.WriteHttpOkResponse(req.Context(), response, log, w)
			break
		default:
			utils.WriteHttpInternalServerError(req.Context(), joinErr, log, w)
		}
	}
}
