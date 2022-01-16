package internal

import (
	"database/sql"
	openapi "github.com/mytord/fs/backend/gen/opencliapi"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error, result *openapi.ImplResponse) {
	if _, ok := err.(*openapi.ParsingError); ok {
		// Handle parsing errors
		WriteBadRequestResponse([]openapi.ErrorResponseErrors{
			{Message: err.Error()},
		}, w)
		return
	}

	if requiredErr, ok := err.(*openapi.RequiredError); ok {
		// Handle missing required errors
		WriteBadRequestResponse([]openapi.ErrorResponseErrors{
			{Message: requiredErr.Error()},
		}, w)
		return
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errors []openapi.ErrorResponseErrors

		for _, err := range validationErrors {
			errors = append(errors, openapi.ErrorResponseErrors{
				Message: err.Error(),
			})
		}

		WriteBadRequestResponse(errors, w)
		return
	}

	if err == sql.ErrNoRows {
		// Handle 404 errors
		WriteNotFoundResponse(w)
		return
	}

	// Handle domain errors
	if err == ErrProfileAlreadyExists {
		WriteBadRequestResponse(
			[]openapi.ErrorResponseErrors{
				{Message: err.Error()},
			}, w)
		return
	}

	zap.L().Error("something went wrong", zap.Error(err))

	// Handle all other errors
	openapi.EncodeJSONResponse(openapi.ErrorResponse{
		Errors: []openapi.ErrorResponseErrors{
			{Message: "unexpected error"},
		},
	}, &result.Code, result.Headers, w)
}

func WriteBadRequestResponse(errors []openapi.ErrorResponseErrors, w http.ResponseWriter) {
	openapi.EncodeJSONResponse(openapi.ErrorResponse{Errors: errors}, func(i int) *int { return &i }(http.StatusBadRequest), map[string][]string{}, w)
}

func WriteNotFoundResponse(w http.ResponseWriter) {
	openapi.EncodeJSONResponse(openapi.ErrorResponse{
		Errors: []openapi.ErrorResponseErrors{
			{Message: "not found"},
		},
	}, func(i int) *int { return &i }(http.StatusNotFound), map[string][]string{}, w)
}
