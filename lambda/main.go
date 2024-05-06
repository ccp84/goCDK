package main

import (
	"fmt"
	"lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
}

// Take in a payload and do something with it
func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}

	return fmt.Sprintf("Successfully called by %s", event.Username), nil
}
func main() {
	// call the lambda SDK lambda invokation
	// pass the handler to call when lambda is invoked but do not call it here
	myApp := app.NewApp()
	lambda.Start(HandleRequest)
}
