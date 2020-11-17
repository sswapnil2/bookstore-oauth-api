package access_token

import (
	"bookstore-ouath-api/src/utils"
	"strings"
)

type Repository interface {
	GetById(id string) (*AccessToken, *utils.RestErr)
	CreateToken(token AccessToken) *utils.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *utils.RestErr)
	CreateToken(token AccessToken) *utils.RestErr
}

type service struct {
	Repo Repository `json:"repo"`
}

func NewService(repo Repository) Service {
	return &service{
		Repo: repo,
	}
}

func (s *service) GetById(id string) (*AccessToken, *utils.RestErr) {

	tokenId := strings.TrimSpace(id)
	if len(tokenId) == 0 {
		return nil, utils.NewBadRequestError("Token cannot be empty")
	}
	accessToken, err := s.Repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) CreateToken(token AccessToken) *utils.RestErr {
	panic("implement me")
}
