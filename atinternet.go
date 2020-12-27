package atinternet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	utilities "github.com/leapforce-libraries/go_utilities"
)

const (
	APIURL string = "https://api.atinternet.io/v3/data"
)

// type
//
type ATInternet struct {
	accessKey             string
	secretKey             string
	maxRetries            uint
	secondsBetweenRetries uint32
}

type ATInternetConfig struct {
	AccessKey             string
	SecretKey             string
	MaxRetries            *uint
	SecondsBetweenRetries *uint32
}

func NewATInternet(config ATInternetConfig) (*ATInternet, *errortools.Error) {
	atInternet := new(ATInternet)

	if config.AccessKey == "" {
		return nil, errortools.ErrorMessage("ATInternet AccessKey not provided")
	}
	atInternet.accessKey = config.AccessKey

	if config.SecretKey == "" {
		return nil, errortools.ErrorMessage("ATInternet SecretKey not provided")
	}
	atInternet.secretKey = config.SecretKey

	if config.MaxRetries != nil {
		atInternet.maxRetries = *config.MaxRetries
	} else {
		atInternet.maxRetries = 0
	}

	if config.SecondsBetweenRetries != nil {
		atInternet.secondsBetweenRetries = *config.SecondsBetweenRetries
	} else {
		atInternet.secondsBetweenRetries = 3
	}

	return atInternet, nil
}

func (ai *ATInternet) apiKey() string {
	return fmt.Sprintf("%s_%s", ai.accessKey, ai.secretKey)
}

// generic Get method
//
func (ai *ATInternet) Get(urlPath string, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return ai.httpRequest(http.MethodGet, urlPath, nil, responseModel)
}

// generic Post method
//
func (ai *ATInternet) Post(urlPath string, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return ai.httpRequest(http.MethodPost, urlPath, bodyModel, responseModel)
}

func (ai *ATInternet) httpRequest(httpMethod string, urlPath string, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	client := new(http.Client)

	url := fmt.Sprintf("%s/%s", APIURL, urlPath)
	fmt.Println(url)

	buffer := new(bytes.Buffer)
	buffer = nil

	if bodyModel != nil {
		b, err := json.Marshal(bodyModel)
		if err != nil {
			return nil, nil, errortools.ErrorMessage(err)
		}
		//fmt.Println(string(b)) //temp
		buffer = bytes.NewBuffer(b)
	}

	ee := new(errortools.Error)

	request, err := func() (*http.Request, error) {
		// function necessary because a Buffer nil pointer differs from a nil value
		if buffer == nil {
			return http.NewRequest(httpMethod, url, nil)
		}
		return http.NewRequest(httpMethod, url, buffer)
	}()

	ee.SetRequest(request)

	if err != nil {
		ee.SetMessage(err)
		return request, nil, ee
	}

	// Add authorization token to header
	request.Header.Set("Accept", "application/json")
	request.Header.Set("x-api-key", ai.apiKey())

	if bodyModel != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	// Send out the HTTP request
	response, e := utilities.DoWithRetry(client, request, ai.maxRetries, ai.secondsBetweenRetries)
	ee.SetResponse(response)

	if response != nil {
		// Check HTTP StatusCode
		if response.StatusCode < 200 || response.StatusCode > 299 {
			fmt.Println(fmt.Sprintf("ERROR in %s", httpMethod))
			fmt.Println("url", url)
			fmt.Println("StatusCode", response.StatusCode)

			ee.SetMessage(fmt.Sprintf("Server returned statuscode %v", response.StatusCode))
		}

		if response.Body != nil {

			defer response.Body.Close()

			b, err := ioutil.ReadAll(response.Body)
			if err != nil {
				ee.SetMessage(err)
				return request, response, ee
			}

			if e != nil {
				// try to unmarshal to ErrorResponse
				errorResponse := ErrorResponse{}
				errError := json.Unmarshal(b, &errorResponse)

				if errError == nil {
					if errorResponse.ErrorMessage != "" {
						ee.SetMessage(errorResponse.ErrorMessage)
					}
				} else {
					// try to unmarshal to string
					errorString := ""
					errError = json.Unmarshal(b, &errorString)

					if errorString != "" {
						ee.SetMessage(errorString)
					}
				}

				errortools.CaptureInfo(errError)

				return request, response, ee
			}

			if responseModel != nil {
				err = json.Unmarshal(b, &responseModel)
				if err != nil {
					ee.SetMessage(err)
					return request, response, ee
				}
			}
		}
	}

	if e != nil {
		return request, response, e
	}

	return request, response, nil
}
