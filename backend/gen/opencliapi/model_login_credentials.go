/*
 * First social
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type LoginCredentials struct {

	Email string `json:"email"`

	Password string `json:"password"`
}

// AssertLoginCredentialsRequired checks if the required fields are not zero-ed
func AssertLoginCredentialsRequired(obj LoginCredentials) error {
	elements := map[string]interface{}{
		"email": obj.Email,
		"password": obj.Password,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseLoginCredentialsRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of LoginCredentials (e.g. [][]LoginCredentials), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseLoginCredentialsRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aLoginCredentials, ok := obj.(LoginCredentials)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertLoginCredentialsRequired(aLoginCredentials)
	})
}
