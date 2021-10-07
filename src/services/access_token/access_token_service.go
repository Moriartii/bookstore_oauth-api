package access_token

import (
	"github.com/Moriartii/bookstore_oauth-api/src/domain/access_token"
	"github.com/Moriartii/bookstore_oauth-api/src/repository/db"
	"github.com/Moriartii/bookstore_oauth-api/src/repository/rest"
	"github.com/Moriartii/bookstore_oauth-api/src/utils/errors"
	"strings"
)

// type Repository interface {
// 	GetById(string) (*AccessToken, *errors.RestErr)
// 	Create(AccessToken) *errors.RestErr
// 	UpdateExpirationTime(AccessToken) *errors.RestErr
// }

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo        db.DbRepository
}

func NewService(usersRepo rest.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invaild or empty access token id")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)

	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	user, err := s.restUsersRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
	// return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
