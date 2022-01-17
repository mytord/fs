/*
 * First social
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ProfileCollection struct {

	Entities []Profile `json:"entities,omitempty"`

	HasMore bool `json:"hasMore,omitempty"`
}

// AssertProfileCollectionRequired checks if the required fields are not zero-ed
func AssertProfileCollectionRequired(obj ProfileCollection) error {
	for _, el := range obj.Entities {
		if err := AssertProfileRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseProfileCollectionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ProfileCollection (e.g. [][]ProfileCollection), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseProfileCollectionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aProfileCollection, ok := obj.(ProfileCollection)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertProfileCollectionRequired(aProfileCollection)
	})
}