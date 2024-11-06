package config

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
)

// AppConfig holds the application settings
type AppConfig struct {
	Name          string
	Env           string
	AWS           *session.Session
	StorageBucket string
	// AssetsDistributionURL is the base URL where the uploaded assets are distributed
	AssetsDistributionURL string
}

func (app *AppConfig) setup() {
	if os.Getenv("ENV") == "LOCAL" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	if err != nil {
		log.Fatalf("Error initializing aws sdk session: %s", err.Error())
	}

	app.Name = "HEALTHDECODE_PRO_FILE_MANAGER"
	app.Env = os.Getenv("ENV")
	app.AWS = awsSession
	app.StorageBucket = os.Getenv("STORAGE_BUCKET")
	app.AssetsDistributionURL = os.Getenv("ASSETS_DISTRIBUTION_URL")

	log.Printf("Application %s is running in %s env\n", app.Name, app.Env)
}

// InitializeApp initializes the application and returns the its configuration
func InitializeApp() *AppConfig {
	app := &AppConfig{}
	app.setup()
	return app
}
