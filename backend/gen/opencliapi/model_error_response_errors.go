/*
 * First social
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ErrorResponseErrors struct {

	Message string `json:"message,omitempty"`

	Params map[string]interface{} `json:"params,omitempty"`
}

// AssertErrorResponseErrorsRequired checks if the required fields are not zero-ed
func AssertErrorResponseErrorsRequired(obj ErrorResponseErrors) error {
	return nil
}

// AssertRecurseErrorResponseErrorsRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ErrorResponseErrors (e.g. [][]ErrorResponseErrors), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseErrorResponseErrorsRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aErrorResponseErrors, ok := obj.(ErrorResponseErrors)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertErrorResponseErrorsRequired(aErrorResponseErrors)
	})
}
