/*
 * First social
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type BaseCollection struct {

	Entities []map[string]interface{} `json:"entities,omitempty"`

	HasMore bool `json:"hasMore,omitempty"`
}

// AssertBaseCollectionRequired checks if the required fields are not zero-ed
func AssertBaseCollectionRequired(obj BaseCollection) error {
	return nil
}

// AssertRecurseBaseCollectionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of BaseCollection (e.g. [][]BaseCollection), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseBaseCollectionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aBaseCollection, ok := obj.(BaseCollection)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertBaseCollectionRequired(aBaseCollection)
	})
}