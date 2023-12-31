package floom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type DataType int

type DataTransferType int

const (
	Base64 DataTransferType = iota + 1
)

type ResponseValue struct {
	Type   DataType
	Format string
	Value  string
	B64    string
	URL    string
}

type FloomResponse struct {
	MessageID      string
	ChatID         string
	Values         []ResponseValue
	ProcessingTime int
}

type FloomRequest struct {
	PipelineID   string
	ChatID       string
	Input        string
	Variables    map[string]string
	DataTransfer DataTransferType
}

type FloomClient struct {
	URL    string
	APIKey string
}

func NewFloomClient(endpoint, apiKey string) *FloomClient {
	return &FloomClient{
		URL:    endpoint,
		APIKey: apiKey,
	}
}

func (c *FloomClient) Run(pipelineID, chatID string, input interface{}, variables map[string]string, dataTransfer DataTransferType) (*FloomResponse, error) {
	httpClient := &http.Client{Timeout: time.Second * 160}

	floomRequest := FloomRequest{
		PipelineID:   pipelineID,
		ChatID:       chatID,
		Input:        fmt.Sprintf("%v", input),
		Variables:    variables,
		DataTransfer: dataTransfer,
	}

	requestBody, err := json.Marshal(floomRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/Pipelines/Run", c.URL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Api-Key", c.APIKey)
	request.Header.Set("Content-Type", "application/json")

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		var floomResponse FloomResponse
		err = json.NewDecoder(response.Body).Decode(&floomResponse)
		if err != nil {
			return nil, err
		}
		return &floomResponse, nil
	}

	// Read the response body and include it in the error
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// Unable to read the response body
		return nil, fmt.Errorf("request failed with status %d and unable to read response body", response.StatusCode)
	}

	// Log or return the detailed error
	errorMessage := fmt.Sprintf("request failed with status %d: %s", response.StatusCode, string(bodyBytes))
	fmt.Println(errorMessage) // You can log it or just return as an error
	return nil, fmt.Errorf(errorMessage)
}
