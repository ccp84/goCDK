# Welcome to your CDK Go project!

## This project was built using cdk init --language go

This is a blank project for CDK development with Go.

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests
 
 * `go get` - get package imports

 * `go get github.com/aws/aws-lambda-go/lambda` - get lambda SDK

 * `"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"` - lambda new function 

 * - `Runtime: awslambda.Runtime_PROVIDED_AL2023(),` - runtime since changes as go not native 

 * `GOOS=linux GOARCH=amd64 go build -o bootstrap` - to build executable

 * `zip function.zip bootstrap` - to zip the created binary this must match Code element in props of the NewFunction 

 ## Write a Makefile to zip lambda code

 ```
 build:
	@GOOS=linux GOARCH=amd64 go build -o bootstrap
	@zip function.zip bootstrap

```
to run `make build`

 
