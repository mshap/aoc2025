# AoC2025

A Go project for Advent of Code 2025 solutions.

## Project Structure

- `cmd/` - Main application entry points
- `pkg/` - Public library code
- `internal/` - Private application code

## Getting Started

### Prerequisites

- Go 1.21 or later

### Building and Running

1. Clone or navigate to this repository
2. Build the project:
   ```bash
   go build -o bin/aoc2025 ./cmd
   ```

3. Run the project:
   ```bash
   ./bin/aoc2025
   ```

   Or run directly:
   ```bash
   go run ./cmd
   ```

### Development

- Run tests: `go test ./...`
- Format code: `go fmt ./...`
- Lint code: `go vet ./...`
- Update dependencies: `go mod tidy`

## License

This project is licensed under the MIT License.