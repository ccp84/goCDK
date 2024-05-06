package database

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDBClient() DynamoDBClient {
	// Open a new session
	dbSession := session.Must(session.NewSession())
	// Create a db client from the session
	db := dynamodb.New(dbSession)
	return DynamoDBClient{
		databaseStore: db,
	}
}
