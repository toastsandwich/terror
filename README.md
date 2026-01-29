# TracedError

A Go error wrapper that captures stack traces to help with debugging and error tracking.

## Overview

TracedError enhances standard Go errors by capturing the call stack at the point where the error is created or wrapped. This makes it easier to trace the origin of errors through complex call chains, especially in larger applications where errors may pass through multiple layers.

## Features

- **Stack Trace Capture**: Automatically captures the call stack when errors are created or wrapped
- **Standard Error Interface**: Fully implements the `error` interface
- **Error Wrapping**: Supports Go 1.13+ error wrapping with `Unwrap()` method
- **Formatted Messages**: Supports both simple and formatted error messages
- **Configurable Stack Depth**: Adjustable stack trace depth for performance tuning
- **Type Compatibility**: Works with `errors.Is()` and `errors.As()` for error type checking

## Installation

```bash
go get github.com/toastsandwich/error
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/toastsandwich/error"
)

func businessLogic() error {
    return error.New(fmt.Errorf("database connection failed"))
}

func main() {
    err := businessLogic()
    if err != nil {
        tracedErr, ok := err.(*error.TracedError)
        if ok {
            fmt.Printf("Error: %s\n", tracedErr)
            fmt.Printf("Stack trace:\n%s", tracedErr.Trace())
        }
    }
}
```

## API Reference

### Constructors

#### `New(err error) *TracedError`
Creates a new TracedError from an existing error, capturing the call stack.

```go
err := error.New(fmt.Errorf("something went wrong"))
```

#### `Newf(format string, args ...any) *TracedError`
Creates a new TracedError with a formatted message.

```go
err := error.Newf("invalid input: %s is not allowed", userInput)
```

#### `Wrap(err error, msg string) *TracedError`
Wraps an existing error with an additional message.

```go
err := error.Wrap(originalErr, "failed to process user request")
```

#### `Wrapf(err error, format string, args ...any) *TracedError`
Wraps an existing error with a formatted message.

```go
err := error.Wrapf(originalErr, "processing failed for user %d", userID)
```

### Configuration

#### `Init(depth int)`
Sets the maximum depth of the captured stack trace. Default is 32.

```go
error.Init(64) // Capture up to 64 stack frames
```

### Methods

#### `Error() string`
Returns the error message. If a custom message was provided during wrapping, it's prepended to the original error.

#### `Unwrap() error`
Returns the underlying wrapped error, enabling compatibility with Go's error wrapping features.

#### `Trace() string`
Returns the formatted stack trace showing the call path leading to the error creation.

## Usage Examples

### Basic Error Creation

```go
func processFile(filename string) error {
    if filename == "" {
        return error.Newf("filename cannot be empty")
    }
    // ... file processing logic
    return nil
}
```

### Error Wrapping

```go
func readConfig(path string) error {
    data, err := os.ReadFile(path)
    if err != nil {
        return error.Wrap(err, "failed to read config file")
    }
    // ... parse config
    return nil
}
```

### Stack Trace Analysis

```go
func handleError(err error) {
    if tracedErr, ok := err.(*error.TracedError); ok {
        fmt.Printf("Error occurred: %s\n", tracedErr.Error())
        fmt.Println("Stack trace:")
        fmt.Println(tracedErr.Trace())
    }
}
```

### Error Type Checking

```go
var (
    ErrNotFound = errors.New("resource not found")
    ErrInvalid  = errors.New("invalid input")
)

func processRequest(id string) error {
    if id == "" {
        return error.Wrap(ErrInvalid, "request validation failed")
    }
    // ... processing logic
    return nil
}

func main() {
    err := processRequest("")
    if errors.Is(err, ErrInvalid) {
        fmt.Println("Invalid input error occurred")
        if tracedErr, ok := err.(*error.TracedError); ok {
            fmt.Println(tracedErr.Trace())
        }
    }
}
```

## Performance Considerations

- Stack trace capture has a small performance overhead
- Use `Init()` to adjust the stack depth based on your needs
- Consider disabling in production hot paths if performance is critical

## Best Practices

1. **Use at Boundaries**: Wrap errors at layer boundaries (API handlers, service boundaries)
2. **Add Context**: Include relevant context when wrapping errors
3. **Don't Overwrap**: Avoid wrapping the same error multiple times unnecessarily
4. **Consider Performance**: Be mindful of stack trace overhead in performance-critical code

## Testing

Run the test suite:

```bash
go test -v
```

## License

This project is licensed under the MIT License.