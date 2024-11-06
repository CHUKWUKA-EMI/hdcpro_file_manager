package main

import (
	"healthdecodepro_file_manager/internal/config"
	awslambda "healthdecodepro_file_manager/internal/platform/aws_lambda"
)

var appConfig *config.AppConfig

func init() {
	app := config.InitializeApp()
	appConfig = app
}

func main() {
	awslambda.StartLambdaHandler(appConfig)
}
