# 🐰 Animated bunny sign for your terminal (Go version)

This Go package lets you display animated messages with an adorable bunny holding a sign. Because why not?

```
 |￣￣￣￣￣￣￣￣￣￣￣|
    Hey, kids!
    don't you want 
    a bunny for your 
    terminal?
 |＿＿＿＿＿＿＿＿＿＿＿|
   (\__/) ||
   (•ㅅ•) ||
   /    づ
```

This project has been done for educational purposes, it's my first Go package, so expect some rough edges.
It is actually the Go port of the original [bunny-sign](https://github.com/fsgreco/bunny-sign) JavaScript package.

### Use Cases

- Announce important messages with the gravitas that only a bunny can provide
- Improve the mood of developers reviewing your logs
- Make your Go applications inexplicably delightful
- Just a new funny way to waste time on your terminal

## Installation

```bash
go get github.com/fsgreco/go-bunny-sign
```

## Usage

### Command Line

```bash
# Build and run
go build -o bunny-sign
./bunny-sign "Hello, World!"

# Or run directly
go run main.go "Hello, World!"

# Multiple messages
go run main.go "First message" "Second message"
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

# Build the binary
go build -o bunny-sign
```

## Contributing

Pull requests are welcome! Just ensure no bunnies are harmed during the process. 

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests if applicable
5. Commit your changes (`git commit -m 'Add some amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request.

## License

MIT

## Related Projects

- [bunny-sign](https://github.com/fsgreco/bunny-sign) - The original JavaScript version

---

*Made with ❤️ and questionable life choices by fsgreco*
