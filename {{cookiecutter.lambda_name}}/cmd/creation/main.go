package main

import (
	{%- if cookiecutter.use_dynamodb == "y" %}
	"github.com/Ryanair/goaws"
	"github.com/Ryanair/goaws/dynamodb"
	{%- endif %}
	"github.com/Ryanair/goaws/lambda/apigw"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	{%- if cookiecutter.use_dynamodb == "y" %}
	"github.com/kelseyhightower/envconfig"
	{%- endif %}

	"stash.ryanair.com/{{cookiecutter.project_name}}/{{cookiecutter.lambda_name}}/internal/logger"
	"stash.ryanair.com/{{cookiecutter.project_name}}/{{cookiecutter.lambda_name}}/pkg/{{cookiecutter.lambda_name}}"
	"stash.ryanair.com/{{cookiecutter.project_name}}/{{cookiecutter.lambda_name}}/pkg/http/rest"
	"stash.ryanair.com/{{cookiecutter.project_name}}/{{cookiecutter.lambda_name}}/pkg/storage"
)

{% if cookiecutter.use_dynamodb == "y" -%}
const prefix = "{{cookiecutter.project_abbreviation | upper}}"

type EnvVariables struct {
	TableName        string
	DynamoDBEndpoint string
}
{% endif -%}
var (
	{%- if cookiecutter.use_dynamodb == "y" %}
	envVars         EnvVariables
	dynamodbClient  *dynamodb.Client
	repository      *storage.DynamoDBAdapter	
	{%- else %}
	repository      *storage.InMemoryDBAdapter  
	{%- endif %}
	service         *{{cookiecutter.lambda_name}}.Service
	creationHandler *rest.CreationHandler
)

// nolint
func init() {
	logger.Initialize()
	{%-if cookiecutter.use_dynamodb == "y" %}
	processErr := envconfig.Process(prefix, &envVars)
	if processErr != nil {
		panic("cannot properly process environment variables")
	}
	awsConfig, configErr := goaws.NewConfig()
	if configErr != nil {
		panic("cannot properly configure goaws.Config")
	}
	dynamodbClient = dynamodb.NewClient(awsConfig, dynamodb.Endpoint(envVars.DynamoDBEndpoint))
	repository = storage.NewDynamoDBAdapter(dynamodbClient, envVars.TableName)
	{%- else %}
	repository = storage.NewInMemoryDBAdapter()
	{%- endif %}
	service = {{cookiecutter.lambda_name}}.NewService(repository)
	creationHandler = rest.NewCreationHandler(service)
}

func main() {
	lambda.Start(handler)
}

func handler(event *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := apigw.NewRequest(event.Resource, event.HTTPMethod, apigw.RequestBody(event.Body))
	res, err := creationHandler.Handle(req)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return res.Convert(), nil
}