package model

// PresignedUploadURLResponse is the response returned by the GetPresignedURL handler
type PresignedUploadURLResponse struct {
	UploadURL   string `json:"upload_url"`
	DownloadURL string `json:"download_url"`
}

// PresignedDeleteURLResponse is the response returned by the GetPresignedURL handler
type PresignedDeleteURLResponse struct {
	DeleteURL string `json:"delete_url"`
}

// PresignedURLRequest is the request body for the GetPresignedURL handler
type PresignedURLRequest struct {
	Location string `json:"location" validate:"required"`
	FileName string `json:"file_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	UserID   string `json:"user_id" validate:"required"`
}
