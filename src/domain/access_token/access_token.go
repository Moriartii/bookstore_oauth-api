package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/Moriartii/bookstore_oauth-api/src/utils/crypto_utils"
	"github.com/Moriartii/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant_type
	Username string `json:"username"`
	Password string `json:"password"`

	//User for client_cridentials grant_type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {

	case grantTypePassword:
		break

	case grantTypeClientCredentials:
		break

	default:
		return errors.NewBadRequestError("invaild grant_type parameter")
	}

	//TODO: Validate parameters for each grant_type
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

//Web frontend - CLient-Id: 123
// Android APP - CLient-Id: 234 (expired time longer)

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invaild access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invaild user id")
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequestError("invaild client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invaild expiration time")
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	//now := time.Now().UTC()
	//expirationTime := time.Unix(at.Expires, 0)
	//return expirationTime.Before(now)
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
