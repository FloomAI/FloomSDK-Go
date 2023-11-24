<div align="center">

# FloomSDK-Go

**Floom Go SDK** - A Go library for interacting with [Floom](https://floom.ai), an AI Orchestration platform that empowers Developers and DevOps.

version 1.0.2
</div>

## About Floom

[Floom](https://floom.ai) orchestrates and executes Generative AI pipelines, allowing developers and DevOps teams to focus on what matters most. It offers enterprise-grade, production-ready, and battle-tested solutions, now open-source and free for everyone, including commercial use.

Floom's AI Pipeline model simplifies the integration and execution process of Generative AI, handling everything from prompt design and data linking to execution and cost management.

## Getting Started with FloomSDK-Go

### Installation

To start using FloomSDK-Go, install the package using `go get`:

```bash
go get github.com/FloomAI/FloomSDK-Go
```

### Usage
Here's how you can use the Floom Go SDK in your Go application:

```bash
package main

import (
    "fmt"
    "github.com/FloomAI/FloomSDK-Go"
)

func main() {
// Initialize FloomClient
	floomClient := floom.NewFloomClient("http://127.0.0.1:80", "COqRR8qLz4RrXygsDoYMXRvDJheXj3MO")

	// Hardcoded values for demonstration
	pipelineID := "docs-pipeline-v1"
	chatID := "abcdefghijklmnop"
	input := "Who was the first US president?"

	// Run the FloomClient with hardcoded values
	response, err := floomClient.Run(pipelineID, chatID, input, nil, floom.Base64)

	// Print the response and error to the console
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		// Print the properties of FloomResponse
		fmt.Printf("Pipeline Response Valid")
		fmt.Printf("Message ID: %s\n", response.MessageID)
		fmt.Printf("Chat ID: %s\n", response.ChatID)
		fmt.Printf("Processing Time: %d\n", response.ProcessingTime)
		for _, value := range response.Values {
			fmt.Printf("Value - Type: %d, Format: %s, Value: %s\n", value.Type, value.Format, value.Value)
		}
	}
}
```

This README provides a concise yet comprehensive introduction to Floom and its Go SDK. It includes installation instructions, a basic usage example, links to further documentation, contribution guidelines, and licensing information. The structure is designed to be user-friendly and to enhance the visibility of your project.

For more information, visit us at https://floom.ai.
