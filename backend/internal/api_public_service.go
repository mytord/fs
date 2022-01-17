/*
 * First social
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package internal

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	openapi "github.com/mytord/fs/backend/gen/opencliapi"
	"github.com/mytord/fs/backend/internal/entities"
	"github.com/mytord/fs/backend/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var (
	ErrProfileAlreadyExists = errors.New("profile with the same email already exists")
)

// PublicApiService is a service that implements the logic for the PublicApiServicer
// This service should implement the business logic for every endpoint for the PublicApi API.
// Include any external packages or services that will be required by this service.
type PublicApiService struct {
	profileRep *repositories.ProfileRepository
	validate   *validator.Validate
}

// NewPublicApiService creates a default api service
func NewPublicApiService(profileRep *repositories.ProfileRepository) openapi.PublicApiServicer {
	return &PublicApiService{
		profileRep: profileRep,
		validate:   validator.New(),
	}
}

// CreateProfile - Register new profile
func (s *PublicApiService) CreateProfile(ctx context.Context, profile openapi.Profile) (openapi.ImplResponse, error) {
	profileEntity := &entities.Profile{
		Email:     profile.Email,
		Password:  profile.Password,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		City:      profile.City,
		Age:       int(profile.Age),
		Interests: profile.Interests,
	}

	err := s.validate.Struct(profileEntity)

	if err != nil {
		return ErrorResponse(err)
	}

	alreadyExists, err := s.profileRep.ExistsByEmail(profileEntity.Email)

	if err != nil {
		return ErrorResponse(err)
	}

	if alreadyExists {
		return ErrorResponse(ErrProfileAlreadyExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(profileEntity.Password), 8)

	if err != nil {
		return ErrorResponse(err)
	}

	profileEntity.Password = string(hashedPassword)

	err = s.profileRep.Add(profileEntity)

	if err != nil {
		return ErrorResponse(err)
	}

	return AuthorizedSuccessResponse(profileEntity.Id, nil)
}

// Login - Logs user into the system
func (s *PublicApiService) Login(ctx context.Context, loginCredentials openapi.LoginCredentials) (openapi.ImplResponse, error) {
	profileEntity, err := s.profileRep.FindByEmail(loginCredentials.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return ErrorResponseWithStatusCode(nil, http.StatusUnauthorized)
		}

		return ErrorResponse(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(profileEntity.Password), []byte(loginCredentials.Password))

	if err != nil {
		return ErrorResponseWithStatusCode(nil, http.StatusUnauthorized)
	}

	return AuthorizedSuccessResponse(profileEntity.Id, nil)
}
