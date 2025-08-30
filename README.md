# 🐰 Animated bunny sign for your terminal (Go version)

This Go package lets you display animated messages with an adorable bunny holding a sign. Because, why not?

```
 |￣￣￣￣￣￣￣￣￣￣￣￣￣|
    Hey kids! Don't 
    you want a bunny 
    for your terminal?
 |＿＿＿＿＿＿＿＿＿＿＿＿＿|
   (\__/) ||
   (•ㅅ•) ||
   /    づ
```
> [!NOTE]
> This project has been done for educational purposes, it's my first Go package.
>
> It is actually the Go porting of the original [bunny-sign](https://github.com/fsgreco/bunny-sign) JavaScript package.


### Use Cases

- Announce important messages with the gravitas that only a bunny can provide
- Improve the mood of developers reviewing your logs
- Make your Go applications inexplicably delightful
- Just a new funny way to waste time on your terminal

## Installation

### Using go install (Recommended)

```bash
go install github.com/fsgreco/go-bunny-sign/cmd/bunnysign@latest
```

### Download Pre-built Binaries

Download the latest binary for your platform from [GitHub Releases](https://github.com/fsgreco/go-bunny-sign/releases):

- **Windows**: `bunnysign_v1.0.0_windows_amd64.zip`
- **macOS**: `bunnysign_v1.0.0_darwin_amd64.tar.gz` (Intel) or `bunnysign_v1.0.0_darwin_arm64.tar.gz` (Apple Silicon)
- **Linux**: `bunnysign_v1.0.0_linux_amd64.tar.gz` or `bunnysign_v1.0.0_linux_arm64.tar.gz`

Extract the binary and add it to your PATH.

### As a Library

```bash
go get github.com/fsgreco/go-bunny-sign
```

## Usage

### Command Line
```bash
# After installing with go install
bunnysign "Hello, World!"

# Multiple messages
bunnysign "First message" "Second message"

# Clear after display
bunnysign -c "This message will disappear"
```

### In Your Go Code

Import the package and use it in your applications:

```go
package main

import (
	"github.com/fsgreco/go-bunny-sign/bunnysign"
)

func main() {

	// Show multiple messages in sequence
	messages := []string{"First message", "Then this one", "And finally this"}

	bunnysign.Display(messages, true) // pass false if you want to clear also the last message
}
```

## Building from Source

```bash
# Clone the repository
git clone https://github.com/fsgreco/go-bunny-sign.git
cd go-bunny-sign

# Build the CLI binary
go build -o bin/bunnysign ./cmd/bunnysign
# OR use make if you have it: `make build`

# Run the binary
bin/bunnysign "Hello, Friend!"

# Or install locally
go install ./cmd/bunnysign
```

## Contributing

Pull requests are welcome! Just ensure no bunnies are harmed during the process. 

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes (Add tests if applicable)
4. Commit and push the branch 
5. Open a Pull Request.

## License

MIT

## Related Projects

- [bunny-sign](https://github.com/fsgreco/bunny-sign) - The original JavaScript version

---

*Made with ❤️ and questionable life choices by fsgreco*
