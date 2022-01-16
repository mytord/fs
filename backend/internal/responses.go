package internal

import (
	openapi "github.com/mytord/fs/backend/gen/opencliapi"
	"net/http"
	"strconv"
)

func ErrorResponse(err error) (openapi.ImplResponse, error) {
	return ErrorResponseWithStatusCode(err, http.StatusInternalServerError)
}

func ErrorResponseWithStatusCode(err error, code int) (openapi.ImplResponse, error) {
	return openapi.Response(code, nil), err
}

func SuccessResponse(body interface{}) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusOK, body), nil
}

func AuthorizedSuccessResponse(userId int, body interface{}) (openapi.ImplResponse, error) {
	return openapi.ResponseWithHeaders(http.StatusOK, map[string][]string{
		"X-Set-User-Id": []string{strconv.Itoa(userId)},
	}, body), nil
}
