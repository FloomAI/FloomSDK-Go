<div align="center">

# FloomSDK-Go

**Floom Go SDK** - A Go library for interacting with [Floom](https://floom.ai), an AI Orchestration platform that empowers Developers and DevOps.

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
    floomClient := floom.NewFloomClient("your_endpoint", "your_api_key")

    // Example: Running a pipeline
    response, err := floomClient.Run("your_pipeline_id", "your_chat_id", "your_input", nil, floom.Base64)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Process the response
    fmt.Printf("Response: %+v\n", response)
}
```

This README provides a concise yet comprehensive introduction to Floom and its Go SDK. It includes installation instructions, a basic usage example, links to further documentation, contribution guidelines, and licensing information. The structure is designed to be user-friendly and to enhance the visibility of your project.

For more information, visit us at https://floom.ai.
