# data-diff

A production-ready command-line tool for comparing data structures, built in Go.

## Features

- **Multiple Input Formats**: Support for JSON, YAML, and text files
- **Flexible Output**: Output results in JSON, YAML, or table format
- **Configurable Comparison**: Case-insensitive comparison, field ignoring, and more
- **Structured Logging**: JSON and text logging with configurable levels
- **Configuration Management**: Support for config files, environment variables, and CLI flags
- **Production Ready**: Docker support, CI/CD pipeline, comprehensive testing

## Installation

### From Source

```bash
git clone https://github.com/renepersau/data-diff.git
cd data-diff
make build
```

### Using Docker

```bash
docker build -t data-diff .
docker run --rm -v $(pwd):/data data-diff compare /data/file1.json /data/file2.json
```

## Usage

### Basic Comparison

```bash
# Compare two JSON files
./data-diff compare file1.json file2.json

# Compare with specific output format
./data-diff compare file1.json file2.json --format yaml

# Save output to file
./data-diff compare file1.json file2.json --output result.json
```

### Advanced Options

```bash
# Case-insensitive comparison
./data-diff compare file1.json file2.json --ignore-case

# Ignore specific fields
./data-diff compare file1.json file2.json --ignore-fields timestamp,id

# Show unchanged fields in output
./data-diff compare file1.json file2.json --show-unchanged
```

### Configuration

```bash
# Show current configuration
./data-diff config show

# Initialize configuration file
./data-diff config init

# Use custom config file
./data-diff --config /path/to/config.yaml compare file1.json file2.json
```

### Environment Variables

```bash
export DATA_DIFF_LOG_LEVEL=debug
export DATA_DIFF_LOG_FORMAT=json
export DATA_DIFF_DEBUG=true
```

## Configuration File

Create a `config.yaml` file in your home directory (`~/.data-diff/config.yaml`) or current directory:

```yaml
log_level: info
log_format: json
debug: false
format: json
```

## Examples

### Compare JSON Files

```bash
# Using the example files
./data-diff compare examples/sample1.json examples/sample2.json
```

### Compare with Different Output Formats

```bash
# JSON output (default)
./data-diff compare file1.json file2.json

# YAML output
./data-diff compare file1.json file2.json --format yaml

# Table output
./data-diff compare file1.json file2.json --format table
```

## Development

### Prerequisites

- Go 1.21 or later
- Make (optional, for using Makefile)

### Setup

```bash
# Clone the repository
git clone https://github.com/renepersau/data-diff.git
cd data-diff

# Install dependencies
make deps

# Run tests
make test

# Build the application
make build
```

### Available Make Targets

- `make build` - Build the binary
- `make test` - Run tests
- `make coverage` - Run tests with coverage
- `make lint` - Run linter
- `make clean` - Clean build artifacts
- `make docker-build` - Build Docker image
- `make release-build` - Build release binaries for multiple platforms

### Project Structure

```
data-diff/
├── cmd/data-diff/          # Main application entry point
├── internal/               # Private application code
│   ├── app/               # Application initialization
│   ├── commands/          # CLI commands
│   ├── config/            # Configuration management
│   └── logger/            # Logging utilities
├── pkg/                   # Public packages
│   └── diff/              # Core diff functionality
├── examples/              # Example files
├── .github/workflows/     # CI/CD configuration
├── Dockerfile             # Docker configuration
├── Makefile              # Build automation
└── README.md             # This file
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Version Information

```bash
./data-diff version
```

This will display the current version, build time, and git commit information.