/*
 * First social
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// PrivateApiController binds http requests to an api service and writes the service results to the http response
type PrivateApiController struct {
	service      PrivateApiServicer
	errorHandler ErrorHandler
}

// PrivateApiOption for how the controller is set up.
type PrivateApiOption func(*PrivateApiController)

// WithPrivateApiErrorHandler inject ErrorHandler into controller
func WithPrivateApiErrorHandler(h ErrorHandler) PrivateApiOption {
	return func(c *PrivateApiController) {
		c.errorHandler = h
	}
}

// NewPrivateApiController creates a default api controller
func NewPrivateApiController(s PrivateApiServicer, opts ...PrivateApiOption) Router {
	controller := &PrivateApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the PrivateApiController
func (c *PrivateApiController) Routes() Routes {
	return Routes{
		{
			"GetCurrentProfile",
			strings.ToUpper("Get"),
			"/api/profile",
			c.GetCurrentProfile,
		},
		{
			"GetProfile",
			strings.ToUpper("Get"),
			"/api/profiles/{id}",
			c.GetProfile,
		},
		{
			"ListProfiles",
			strings.ToUpper("Get"),
			"/api/profiles",
			c.ListProfiles,
		},
	}
}

// GetCurrentProfile - Get current profile
func (c *PrivateApiController) GetCurrentProfile(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetCurrentProfile(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetProfile - Get profile by id
func (c *PrivateApiController) GetProfile(w http.ResponseWriter, r *http.Request) {
	idParam, err := parseInt64Parameter(chi.URLParam(r, "id"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetProfile(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// ListProfiles - List profiles
func (c *PrivateApiController) ListProfiles(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limitParam, err := parseInt32Parameter(query.Get("limit"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	offsetParam, err := parseInt32Parameter(query.Get("offset"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	filterFirstNameParam := query.Get("filterFirstName")
	filterLastNameParam := query.Get("filterLastName")
	result, err := c.service.ListProfiles(r.Context(), limitParam, offsetParam, filterFirstNameParam, filterLastNameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}
