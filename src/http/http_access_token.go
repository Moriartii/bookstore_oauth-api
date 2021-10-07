package http

import (
	atDomain "github.com/Moriartii/bookstore_oauth-api/src/domain/access_token"
	"github.com/Moriartii/bookstore_oauth-api/src/services/access_token"
	"github.com/Moriartii/bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := (c.Param("access_token_id")) // strings.TrimSpace)c.Param
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invaild json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
