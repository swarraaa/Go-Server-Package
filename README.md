# GoServerPackage

GoServerPackage is a lightweight and efficient Go package that provides essential functionalities to easily set up a Go HTTP server with an optimized Trie-based router, customizable middleware support, graceful shutdown handling, and a cron job system for periodic tasks.

This package is designed for developers who want to build a performant Go HTTP server with concurrency control and flexibility for task scheduling.

## Key Features

- **Optimized Trie-based Router**: The package implements a fast, memory-efficient Trie-based router for route matching, leading to faster route lookups compared to traditional routing mechanisms.
- **Middleware Support**: Easily add middleware for logging and concurrency control.
- **Graceful Shutdown**: Supports safe and graceful shutdown of the server, ensuring all active requests are processed before termination.
- **Cron Job Scheduler**: Includes a lightweight cron scheduler for running periodic tasks like health checks, backups, or keeping the server alive during deployment.
- **Concurrency Control**: Manage concurrent request handling via a customizable concurrency middleware.

## Optimizations

### 1. Trie-based Router

- **Why Trie?**: The router uses a Trie data structure to store routes, which improves lookup performance, especially with deep route structures.
- **Memory Efficiency**: The Trie structure reduces memory consumption by sharing common path segments across routes.
- **Faster Lookups**: Compared to linear matching or tree-based structures, the Trie reduces the time complexity of route lookups, leading to faster responses.

### 2. Concurrency Middleware

- The concurrency middleware uses a channel to limit the number of concurrent requests being processed. By controlling the number of requests that can be handled simultaneously, the package avoids overloading the server, especially in high-traffic scenarios.

### 3. Graceful Shutdown

- The graceful shutdown mechanism ensures that the server waits for all ongoing requests to complete before shutting down, preventing abrupt terminations of requests and allowing for a clean closure.

### 4. Cron Job Scheduler

- The cron scheduler is built using Go's `time.Ticker` for lightweight periodic task scheduling. An optional integration with `robfig/cron/v3` is available for more complex scheduling requirements (e.g., using cron expressions).

## Installation

To install this package, run:

```bash
go get github.com/swarraaa/go-server-package.git
```

## Usage

To use this package, import it into your Go application:

```bash
import "github.com/swarraaa/go-server-package.git"
```

To checkout the example usage of the package, refer to the [example](github.com/swarraaa/go-server-package.git/USAGE.md) file.

## Contribution

Contributions are welcome! Feel free to open issues or submit pull requests to help improve this package.
