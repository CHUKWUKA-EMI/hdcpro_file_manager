package awslambda

import (
	"encoding/json"
	"fmt"
	"healthdecodepro_file_manager/internal/config"
	"healthdecodepro_file_manager/internal/handler"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func newLambdaHandler(appConfig *config.AppConfig) func(events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return func(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

		switch request.RawPath {
		case fmt.Sprintf("/%s/upload/presigned-url", appConfig.Env):
			response, err := handler.GetUploadURL(appConfig, request.Body)
			if err != nil {
				return handleResponse(400, true, err.Error())
			}
			return handleResponse(200, false, response)
		case fmt.Sprintf("/%s/delete/presigned-url", appConfig.Env):
			response, err := handler.GetDeleteURL(appConfig, request.Body)
			if err != nil {
				return handleResponse(400, true, err.Error())
			}
			return handleResponse(200, false, response)
		default:
			return events.APIGatewayV2HTTPResponse{
				StatusCode: 404,
				Body:       "Not Found",
			}, nil
		}
	}
}

func handleResponse(statusCode int, isError bool, responseBody interface{}) (events.APIGatewayV2HTTPResponse, error) {
	var b map[string]interface{}
	if isError {
		b = map[string]interface{}{"error": responseBody}
	} else {
		b = map[string]interface{}{"data": responseBody}
	}

	body, _ := json.Marshal(b)
	return events.APIGatewayV2HTTPResponse{
		Body:       string(body),
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

// StartLambdaHandler is a function that registers the aws lambda handler
func StartLambdaHandler(appConfig *config.AppConfig) {
	handler := newLambdaHandler(appConfig)
	lambda.Start(handler)
}
