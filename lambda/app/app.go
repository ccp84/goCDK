package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	// Initialise the db store
	// gets passeed the API handler
	db := database.NewDynamoDBClient()
	apiHandler := api.NewAPIHandler(db)
	return App{
		ApiHandler: apiHandler,
	}
}
