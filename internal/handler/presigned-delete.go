package handler

import (
	"encoding/json"
	"errors"
	"healthdecodepro_file_manager/internal/config"
	"healthdecodepro_file_manager/internal/model"
	"healthdecodepro_file_manager/internal/storage"
	"strings"
)

// GetDeleteURL returns a pre-signed URL that can be used to delete a file from S3
func GetDeleteURL(app *config.AppConfig, requestBody string) (*model.PresignedDeleteURLResponse, error) {
	var requestData model.PresignedURLRequest
	err := json.Unmarshal([]byte(requestBody), &requestData)
	if err != nil {
		return nil, err
	}

	if requestData.Location == "" || requestData.FileName == "" || requestData.Email == "" || requestData.UserID == "" {
		return nil, errors.New("location, file_name, email and user_id are required")
	}

	if !strings.Contains(requestData.Email, "@") {
		return nil, errors.New("invalid email")
	}

	userName := strings.Split(requestData.Email, "@")[0] + requestData.UserID
	fileDirectory := requestData.Location + "/" + userName
	return storage.NewStorage(app, storage.S3StorageType).
		GetDeleteURL(
			fileDirectory,
			requestData.FileName,
		)
}
