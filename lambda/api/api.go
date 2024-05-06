package api

import (
	"lambda-func/database"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewAPIHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}
