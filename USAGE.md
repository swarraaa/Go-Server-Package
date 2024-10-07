# Usage of GoServerPackage

Below are example use cases to help you integrate and use the GoServerPackage in your project.

## Basic Usage

## 1. Creating and Starting the Server

The server is built with an efficient Trie-based router and supports graceful shutdown. Here’s how you can create and start a server.

```go
package main

import (
    "github.com/yourusername/go-server-package"
    "net/http"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    // Initialize the server
    server := goServer.NewServer()

    // Add a simple route
    server.AddRoute("/", func(w http.ResponseWriter, r *http.Request) {
        goServer.JSONResponse(w, map[string]string{"message": "Hello, World!"}, http.StatusOK)
    })

    // Start the server on port 8080
    server.StartServer("8080", &wg)

    // Wait for shutdown signal
    wg.Wait()
}
```

### Server Details

- `NewServer()`: Initializes the server with a Trie-based router for faster route matching.
- `AddRoute()`: Adds a route to the server.
- `StartServer()`: Starts the server on the specified port and handles incoming requests.

## 2. Middleware Support

This package includes middleware support to extend the functionality of the server. Two default middlewares are provided: logging and concurrency control.

```go
package main

import (
    "github.com/yourusername/go-server-package"
    "net/http"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    // Create a new server
    server := goServer.NewServer()

    // Add a simple route
    server.AddRoute("/", func(w http.ResponseWriter, r *http.Request) {
        goServer.JSONResponse(w, map[string]string{"message": "Welcome to the Go Server!"}, http.StatusOK)
    })

    // Apply middlewares
    server.AddMiddleware(goServer.LoggingMiddleware)         // Logs incoming requests
    server.AddMiddleware(goServer.ConcurrencyMiddleware(5))  // Limits concurrent requests to 5

    // Start the server on port 8080
    server.StartServer("8080", &wg)

    // Wait for shutdown signal
    wg.Wait()
}

```

### Middleware Details

#### `LoggingMiddleware`

- **Purpose**: Logs each incoming HTTP request for visibility.
- **Usage**: The middleware automatically logs the method and URL of the request.
- **Example Output**:

  ```
  Processing request GET /
  Processing request POST /submit
  ```

#### `ConcurrencyMiddleware`

- **Purpose**: Limits the number of concurrent requests that the server can process.
- **Implementation**: The middleware uses a buffered channel (semaphore pattern) to limit concurrent access. Requests beyond the defined limit will be blocked until capacity is available.
- **Parameter**: maxConcurrency (int) – defines the maximum number of concurrent requests allowed.

## 3. Routing

The server uses a Trie-based router for efficient routing. Below is a detailed explanation of how to add and use routes.

```go
package main

import (
    "github.com/yourusername/go-server-package"
    "net/http"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    // Create a new server
    server := goServer.NewServer()

    // Add a route for the root path
    server.AddRoute("/", func(w http.ResponseWriter, r *http.Request) {
        goServer.JSONResponse(w, map[string]string{"message": "Welcome to the Go Server!"}, http.StatusOK)
    })

    // Add a route for /about path
    server.AddRoute("/about", func(w http.ResponseWriter, r *http.Request) {
        goServer.JSONResponse(w, map[string]string{"message": "About Page"}, http.StatusOK)
    })

    // Start the server on port 8080
    server.StartServer("8080", &wg)

    // Wait for shutdown signal
    wg.Wait()
}
```

### Route Details

- **Adding Routes**: Use the `AddRoute()` method to add routes to the server. The method takes two parameters: the route path and the handler function.
- **Handler Function**: The handler function is executed when a request matches the specified route. It takes two parameters: the response writer and the request object.

## 4. JSON and Error Responses

This package provides utility functions to send JSON responses and error messages to clients.

1. **JSON Response**: Use the `JSONResponse()` function to send JSON data as a response.

```go
package main

import (
    "github.com/yourusername/go-server-package"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Sending a successful JSON response
    goServer.JSONResponse(w, map[string]string{"message": "Success!"}, http.StatusOK)
}
```

2. **Error Response**: Use the `ErrorResponse()` function to send error messages with the appropriate status code.

```go
package main

import (
    "github.com/yourusername/go-server-package"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Sending an error response
    goServer.ErrorResponse(w, "Not Found", http.StatusNotFound)
}
```

## 5. Cron Job Management

This package provides functionality to schedule and manage cron jobs using the robfig/cron library.

```go
package main

import (
    "github.com/yourusername/go-server-package"
    "log"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup

    // Example task to keep the server alive
    keepAliveTask := func() {
        log.Println("Keep-alive task executed!")
    }

    // Start a cron job that runs every 10 seconds
    goServer.StartCronWithRobfig("*/10 * * * * *", keepAliveTask)

    // Keep the main function running
    wg.Add(1)
    go func() {
        defer wg.Done()
        time.Sleep(1 * time.Minute) // Keep alive for 1 minute
    }()
    wg.Wait()
}
```

### Cron Job Details

- **StartCronWithRobfig()**: Starts a cron job with the specified schedule and task. The schedule follows the cron format.

# Conclusion

This package provides a simple and efficient way to create and manage HTTP servers in Go. With middleware support, routing, and cron job management, you can build robust server applications with ease. Feel free to explore the package further and customize it to suit your project requirements.
