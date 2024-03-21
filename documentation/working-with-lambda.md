# Working with Lambda

This guide will walk you through the process of developing and deploying a Lambda function using the AWS SAM CLI.

## Prerequisites

- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)
- [AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Development

Any custom configuration can be defined via the Lambda struct in `config.go` - the starter kit simply provides a basic example for LogLevel to determine the verbosity of the custom slog logger. Once your config is properly defined you can begin defining your Lambda methods via the provided `BaseLambdaHandler()` singleton which is bootstrapped during app initialization. This function takes a `LambdaEvent` and returns an `APIGatewayProxyResponse`. This may or may not fit your use case, but it is a good starting point for most Lambda functions. 

A simple `exampleLambdaHandler()` is provided that simply return a 200 status code and a "Hello, World!" message as a proof of concept. You can use this as a starting point for your own Lambda functions that run via API Gateway. To build and run the Lambda functions locally you can use AWS SAM CLI, which is a CLI tool for local development and testing of Serverless applications. It uses Docker to simulate the AWS Lambda environment locally. The SAM resources are located in the `./lambda` directory, mainly a template.yaml file that defines the Lambda function and the API Gateway endpoint.

To build and run the Lambda function locally, you can use the following commands:

```bash
$ ./lambda/lambda_build.sh
```
You should see a new `.aws-sam` directory created in the root of the project which contains all of the Lambda build artifacts.

To invoke the Lambda function locally with arguments to test the `exampleLambdaHandler` run the following:

```bash
$ sam local invoke -e lambda/example.json
```

