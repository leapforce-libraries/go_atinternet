package atinternet

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	APIURL string = "https://api.atinternet.io/v3/data"
)

// type
//
type Service struct {
	accessKey   string
	secretKey   string
	httpService *go_http.Service
}

type ServiceConfig struct {
	AccessKey             string
	SecretKey             string
	MaxRetries            *uint
	SecondsBetweenRetries *uint32
}

func NewService(config ServiceConfig) (*Service, *errortools.Error) {
	if config.AccessKey == "" {
		return nil, errortools.ErrorMessage("AccessKey not provided")
	}

	if config.SecretKey == "" {
		return nil, errortools.ErrorMessage("SecretKey not provided")
	}

	httpServiceConfig := go_http.ServiceConfig{
		MaxRetries:            config.MaxRetries,
		SecondsBetweenRetries: config.SecondsBetweenRetries,
	}

	return &Service{
		accessKey:   config.AccessKey,
		secretKey:   config.SecretKey,
		httpService: go_http.NewService(httpServiceConfig),
	}, nil
}

func (service *Service) httpRequest(httpMethod string, requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add authentication header
	header := http.Header{}
	header.Set("x-api-key", service.apiKey())
	(*requestConfig).NonDefaultHeaders = &header

	// add error model
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request, response, e := service.httpService.HTTPRequest(httpMethod, requestConfig)
	if errorResponse.ErrorMessage != "" {
		e.SetMessage(errorResponse.ErrorMessage)
	}

	return request, response, e
}

func (service *Service) apiKey() string {
	return fmt.Sprintf("%s_%s", service.accessKey, service.secretKey)
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", APIURL, path)
}

func (service *Service) get(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodGet, requestConfig)
}

func (service *Service) post(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPost, requestConfig)
}

func (service *Service) put(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPut, requestConfig)
}

func (service *Service) delete(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodDelete, requestConfig)
}
