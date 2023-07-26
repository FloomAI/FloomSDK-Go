package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type FirePrompt struct {
	fireEngineURL string
	apiKey        string
}

// Constructor with environment variable-based initialization
func NewFirePrompt() (*FirePrompt, error) {
	fireEngineURL := os.Getenv("FIRE_ENGINE_URL")
	apiKey := os.Getenv("FIRE_ENGINE_API_KEY")

	if fireEngineURL == "" || apiKey == "" {
		return nil, fmt.Errorf("FIRE_ENGINE_URL and FIRE_ENGINE_API_KEY environment variables must be set")
	}

	return &FirePrompt{
		fireEngineURL: fireEngineURL,
		apiKey:        apiKey,
	}, nil
}

// Explicit constructor to pass fireEngineURL and apiKey as arguments
func NewFirePromptWithValues(fireEngineURL, apiKey string) *FirePrompt {
	return &FirePrompt{
		fireEngineURL: fireEngineURL,
		apiKey:        apiKey,
	}
}

func (fp *FirePrompt) Fire(pipelineID string, input interface{}, variables map[string]interface{}) (interface{}, error) {
	client := &http.Client{}

	url := fmt.Sprintf("%s/pipelines/fire?pipelineId=%s", fp.fireEngineURL, pipelineID)

	if input != nil {
		// You can customize the content based on your requirements here
		// For simplicity, let's assume we send the input as JSON
		jsonInput, err := json.Marshal(input)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonInput))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Api-Key", fp.apiKey)

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			var responseBody interface{}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			if err != nil {
				return nil, err
			}
			return responseBody, nil
		}

		return nil, fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Api-Key", fp.apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		responseBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		responseString := string(responseBytes)
		return responseString, nil
	}

	return nil, fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
}

func (fp *FirePrompt) FireAsync(pipelineID string, input interface{}, variables map[string]interface{}) (interface{}, error) {
	// Go does not have native support for async/await
	// Asynchronous behavior can be achieved using goroutines and channels

	// For the sake of simplicity, we will use the same implementation as the synchronous 'Fire' method
	return fp.Fire(pipelineID, input, variables)
}

func main() {
	// Example usage:
	fp, err := NewFirePrompt()
	if err != nil {
		fmt.Println("Error creating FirePrompt:", err)
		return
	}

	response, err := fp.Fire("my-pipeline", map[string]interface{}{"message": "Hello, FireSDK!"}, nil)
	if err != nil {
		fmt.Println("Error making Fire request:", err)
		return
	}

	fmt.Println("Response:", response)
}
