package atinternet

// ErrorResponse stores general Ridder API error response
//
type ErrorResponse struct {
	ErrorCode    string `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`
	ErrorName    string `json:"ErrorName"`
}
