package utils

import (
	"net/http"

	"context"
	"github.com/rs/zerolog"
)

func WriteHttpBadRequestResponse(ctx context.Context, content map[string]interface{}, log zerolog.Logger, w http.ResponseWriter) {
	response, err := MarshalFromMap(content)
	if err != nil {
		log.Error().Err(err).Ctx(ctx).AnErr("Error marshalling json response", err)
		return
	}

	log.Warn().Err(err).Ctx(ctx).AnErr("Bad request received", err)

	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write(response)
}

func WriteHttpOkResponse(ctx context.Context, content map[string]interface{}, log zerolog.Logger, w http.ResponseWriter) {
	response, err := MarshalFromMap(content)
	if err != nil {
		log.Error().Err(err).Ctx(ctx).AnErr("Error marshalling json response", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func WriteHttpConflictResponse(ctx context.Context, content map[string]interface{}, log zerolog.Logger, w http.ResponseWriter) {
	response, err := MarshalFromMap(content)
	if err != nil {
		log.Error().Err(err).Ctx(ctx).AnErr("Error marshalling json response", err)
		return
	}

	w.WriteHeader(http.StatusConflict)
	_, _ = w.Write(response)
}

func WriteHttpNotFoundError(ctx context.Context, content map[string]interface{}, err error, log zerolog.Logger, w http.ResponseWriter) {
	response, marshalErr := MarshalFromMap(content)
	if marshalErr != nil {
		log.Error().Err(marshalErr).Ctx(ctx).AnErr("Error marshalling json response", err)
		return
	}

	log.Warn().Err(err).Ctx(ctx).AnErr("Resource not found", err)

	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write(response)
}

func WriteHttpInternalServerError(ctx context.Context, err error, log zerolog.Logger, w http.ResponseWriter) {
	response, marshalErr := MarshalFromMap(map[string]interface{}{
		"message": "An unexpected error ocurred",
	})

	if marshalErr != nil {
		log.Error().Err(marshalErr).Ctx(ctx).AnErr("Error marshalling json response", err)
		return
	}

	log.Warn().Err(err).Ctx(ctx).AnErr("Application errored unexpectedly", err)

	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(response)
}
