# PHP Filter Chain Generator (Go Version)

This project is a Go reimplementation of [Synacktiv's PHP Filter Chain Generator](https://github.com/synacktiv/php_filter_chain_generator/), providing similar functionality but optimized for use as a command-line binary on Linux systems.

## Features:
- **Binary Executable**: After compiling the Go code, you can place the binary in `/usr/bin/phpchain` for easy access.
- **No Dependencies**: Unlike the Python version, this Go version doesn't require any dependencies to run, making it lightweight and fast.
- **Same Functionality**: Provides the same PHP filter chain generation capabilities as the original project.

## Installation:
1. Clone this repository.
   ```bash
   git clone https://github.com/NumeXx/PHP-Filter-Chain-Generator.git
   cd PHP-Filter-Chain-Generator
   ```
2. Compile the code into a binary
   ```bash
   go build -o phpchain up.go
   ```
3. Move the binary to `/usr/bin`
   ```bash
   sudo mv phpchain /usr/bin/
   ```
## Usage
```bash
phpchain --chain "your_payload_here"
```
## Credits:
This project is inspired by the original [PHP Filter Chain Generator by Synacktiv](https://github.com/synacktiv/php_filter_chain_generator/).
