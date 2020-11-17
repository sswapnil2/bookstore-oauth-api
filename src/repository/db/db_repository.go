package db

import (
	"bookstore-ouath-api/src/clients/cassandra"
	"bookstore-ouath-api/src/domain/access_token"
	"bookstore-ouath-api/src/utils"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens where access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
)

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *utils.RestErr)
	Create(token *access_token.AccessToken) *utils.RestErr
}

type dbRepository struct {
}

func NewRepository() DBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *utils.RestErr) {

	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		return nil, utils.NewInternalServerError(dbErr.Error())
	}
	defer session.Close()
	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {

		if err == gocql.ErrNotFound {
			return nil, utils.NewNotFoundError("Access token not found for id")
		}

		return nil, utils.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(token *access_token.AccessToken) *utils.RestErr {

	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		return utils.NewInternalServerError(dbErr.Error())
	}
	defer session.Close()

	if err := session.Query(queryCreateAccessToken,
		token.AccessToken,
		token.UserId,
		token.ClientId,
		token.Expires).Exec(); err != nil {
		return utils.NewInternalServerError(err.Error())
	}
	return nil

}
