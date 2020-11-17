package http

import (
	"bookstore-ouath-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(ctx *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(ctx *gin.Context) {
	//ctx.JSON(http.StatusNotImplemented, "Method not implemented")

	accessToken, err := handler.service.GetById(ctx.Param("access_token_id"))
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, accessToken)
	return
}
