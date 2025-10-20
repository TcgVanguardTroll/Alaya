# Alaya Programming Language

A calculator-like programming language interpreter written in Go. This project serves as a learning experience for both Go programming and interpreter design, implementing a complete interpreter from scratch following classic compiler construction principles.

## About

**Alaya** is an interpreted programming language built using traditional compiler architecture: Lexer → Parser → Abstract Syntax Tree → Evaluator. The project demonstrates fundamental concepts in language implementation including tokenization, recursive descent parsing, tree-walking interpretation, and REPL design.

**Author:** Jordan Grant ([@TcgVanguardTroll](https://github.com/TcgVanguardTroll))
**License:** MIT
**Language:** Go 1.19+

## Features

### Currently Implemented
- ✅ Lexical analysis with 20+ token types
- ✅ Recursive descent parser with operator precedence
- ✅ Abstract Syntax Tree (AST) representation
- ✅ Tree-walking evaluator for arithmetic expressions
- ✅ Interactive REPL (Read-Eval-Print Loop)
- ✅ Comment support (`#` line comments)
- ✅ Symbol table for variable storage
- ✅ Arithmetic operators: `+`, `-`, `*`, `/`
- ✅ Comparison operators: `==`, `!=`, `<`, `>`
- ✅ Parenthesized expressions

### Planned Features
- ⏳ Variables and assignments (`name` keyword)
- ⏳ Functions (`cmnd` keyword)
- ⏳ Conditionals (`is`, `else` keywords)
- ⏳ Boolean types (`true`, `false`)
- ⏳ Return statements
- ⏳ Error handling and recovery

## Project Structure

```
Alaya/
├── go.mod                          # Go module definition
├── LICENSE                         # MIT License
├── README.md                       # This file
├── .travis.yml                     # CI configuration
└── main/
    ├── main.go                     # Entry point - launches REPL
    ├── alaya_token/                # Token type definitions
    │   └── token.go                # Token struct and constants
    ├── alaya_tokenizer/            # Lexical analyzer
    │   ├── tokenizer.go            # Tokenizer implementation
    │   └── tokenizer_test.go       # Tokenizer tests
    ├── alaya_parser/               # Syntax analyzer
    │   ├── parser.go               # Recursive descent parser
    │   └── parser_test.go          # Parser tests
    ├── alaya_ast/                  # Abstract Syntax Tree
    │   ├── ast.go                  # AST node definitions
    │   └── ast_test.go             # AST tests
    ├── alaya_repl/                 # Interactive shell
    │   └── repl.go                 # REPL implementation
    └── alaya/                      # Interpreter/Evaluator
        └── alaya.go                # Tree-walking evaluator
```

## Architecture

### 1. Tokenizer (Lexer)
Converts source code into a stream of tokens. Handles:
- Single-character operators: `+`, `-`, `*`, `/`, `=`, `<`, `>`
- Multi-character operators: `==`, `!=`
- Keywords: `cmnd`, `name`, `true`, `false`, `is`, `else`, `return`
- Identifiers and integers
- Whitespace and comments

**Location:** `main/alaya_tokenizer/tokenizer.go`

### 2. Parser
Converts token stream into an Abstract Syntax Tree using recursive descent parsing:
```
expr   → term ((+ | -) term)*
term   → factor ((* | /) factor)*
factor → number | identifier | (expr)
```

**Location:** `main/alaya_parser/parser.go`

### 3. Abstract Syntax Tree (AST)
Represents program structure with nodes for:
- `Num` - Numeric literals
- `BinOp` - Binary operations (Left Operator Right)
- `UnaryOp` - Unary operations (Operator Expression)
- `Identifier` - Variable references
- `AsStatement` - Assignment statements

**Location:** `main/alaya_ast/ast.go`

### 4. Evaluator
Tree-walking interpreter that recursively traverses and evaluates the AST using the Visitor pattern.

**Location:** `main/alaya/alaya.go`

### 5. REPL
Interactive shell for testing expressions with prompt `" ** "`

**Location:** `main/alaya_repl/repl.go`

## Getting Started

### Prerequisites
- Go 1.19 or higher
- Git

### Installation

```bash
# Clone the repository
git clone https://github.com/TcgVanguardTroll/Alaya.git
cd Alaya

# Download dependencies
go mod tidy

# Build the interpreter
go build -o alaya ./main
```

### Running the REPL

```bash
./alaya
```

You'll see:
```
Hello <username>! Welcome to the Alaya Programming Language!
 **
```

### Example Usage

```
 ** 2 + 3 * 4
14

 ** (2 + 3) * 4
20

 ** 10 / 2 - 3
2
```

## Language Syntax

### Comments
```python
# This is a comment
```

### Arithmetic Expressions
```
2 + 3           # Addition
5 - 2           # Subtraction
4 * 6           # Multiplication
10 / 2          # Division
(2 + 3) * 4     # Parentheses for grouping
```

### Operators (By Precedence)
1. Parentheses `( )`
2. Unary `+`, `-`
3. Multiplication `*`, Division `/`
4. Addition `+`, Subtraction `-`

## Development Status

### Current State
The project is in active development. The core arithmetic calculator functionality is implemented and the architecture is in place for more advanced features.

### Known Issues
The following functions need to be implemented for full functionality:

**In `main/alaya_tokenizer/tokenizer.go`:**
- `isLetter(ch byte) bool` - Check if character is a letter
- `isDigit(ch byte) bool` - Check if character is a digit
- `readIdentifier() Token` - Read complete identifier
- `readNumber() Token` - Read complete number

**In `main/alaya_ast/ast.go`:**
- Define `Statement` interface for statement nodes

These are intentionally left for learning purposes.

## Learning Goals

This project teaches:

### Go Programming Concepts
- Structs and methods
- Interfaces and polymorphism
- Pointers and references
- Type switches and type assertions
- Maps for symbol tables
- String manipulation
- Package organization
- Testing with `go test`

### Interpreter Design Concepts
- Lexical analysis and tokenization
- Recursive descent parsing
- Operator precedence handling
- Abstract Syntax Trees (AST)
- Tree-walking interpretation
- Visitor pattern
- REPL design
- Symbol tables

## Testing

```bash
# Run all tests
go test ./...

# Test specific package
go test ./main/alaya_tokenizer

# Run with verbose output
go test -v ./...
```

## Implementation Guide

### Implementing Missing Functions

**1. `isLetter(ch byte) bool`**
```go
func isLetter(ch byte) bool {
    // Return true if ch is a-z, A-Z, or _
    // Hint: Use byte comparison like ch >= 'a' && ch <= 'z'
}
```

**2. `isDigit(ch byte) bool`**
```go
func isDigit(ch byte) bool {
    // Return true if ch is 0-9
}
```

**3. `readIdentifier()`**
```go
func (t *Tokenizer) readIdentifier() alaya_token.Token {
    // Build a string while isLetter(t.currentCharacter)
    // Check if it's a keyword using Keywords map
    // Return appropriate token (keyword or IDENT)
}
```

**4. `readNumber()`**
```go
func (t *Tokenizer) readNumber() alaya_token.Token {
    // Build a string while isDigit(t.currentCharacter)
    // Return INTEGER token
}
```

## Contributing

This is a personal learning project, but suggestions and feedback are welcome! Feel free to:
- Open issues for bugs or suggestions
- Submit pull requests for fixes
- Share your own interpreter implementations

## Resources

Recommended reading for understanding this project:
- [Writing An Interpreter In Go](https://interpreterbook.com/) by Thorsten Ball
- [Crafting Interpreters](https://craftinginterpreters.com/) by Robert Nystrom
- [The Go Programming Language](https://www.gopl.io/) by Donovan & Kernighan

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by the classic books on interpreter and compiler design
- Built as a learning project to understand Go and language implementation
- Thanks to the Go community for excellent documentation and tools

---

**Happy Interpreting!** 🚀

For questions or feedback, open an issue on GitHub.
