# Gin Header Enricher

A simple middleware for Gin that automatically enriches HTTP requests with custom headers.

## Features

- Automatically adds custom headers to all incoming requests
- Minimal configuration required
- Simple integration with Gin applications
- High performance with minimal overhead

## Installation

```bash
go get github.com/mithun/gin-headerenricher
```

## Usage

### Quick Start

The easiest way to use this middleware is to replace your `gin.Default()` call with `autogin.DefaultEngine()`:

```go
package main

import (
    "github.com/mithun/gin-headerenricher/autogin"
)

func main() {
    // Create a router with header enrichment already applied
    router := autogin.DefaultEngine()
    
    // Add your routes
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
    
    // Start the server
    router.Run(":8080")
}
```

### Headers Added

The middleware adds the following headers to all requests:

- `X-Request-ID`: A unique identifier for each request
- `X-Request-Timestamp`: The timestamp when the request was received
- `X-Service`: The name of the service
- `X-Processed-By`: Added to response headers

### Custom Configuration

You can customize the behavior by creating a configuration:

```go
// Create custom configuration
config := autogin.HeaderConfig{
    ServiceName:       "MyAwesomeService",
    AddRequestID:      true,
    AddTimestamp:      true,
    AddResponseHeader: true,
}

// Use with DefaultEngine
router := autogin.DefaultEngine(config)

// Or with NewEngine (no Logger/Recovery middlewares)
router := autogin.NewEngine(config)
```

## License

MIT 