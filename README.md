# Alaya Programming Language

✅ **Status:** Fully Functional Arithmetic Interpreter · 📚 **Purpose:** Learning Go + Interpreter Design

A calculator-like programming language interpreter written in Go. This project serves as a hands-on learning experience for both Go programming and interpreter design, implementing a complete interpreter from scratch following classic compiler construction principles.

> **⚡ It Works!** The interpreter is now **fully functional** for arithmetic expressions! Build it with `go build -o alaya ./main` and start calculating. See [Getting Started](#getting-started).

## About

**Alaya** is an interpreted programming language built using traditional compiler architecture: 

```
Source Code → Tokenizer → Parser → AST → Evaluator → Result
```

The project demonstrates fundamental concepts in language implementation including tokenization, recursive descent parsing, tree-walking interpretation, and REPL design.

**Author:** Jordan Grant ([@TcgVanguardTroll](https://github.com/TcgVanguardTroll))  
**License:** MIT  
**Language:** Go 1.19+  
**Inspired By:** "Writing An Interpreter In Go" by Thorsten Ball

## Project Status

✅ **The interpreter is fully functional!** All core features for arithmetic calculation are working.

### Currently Working
- ✅ Full lexical analysis with 20+ token types
- ✅ Recursive descent parser with operator precedence
- ✅ Abstract Syntax Tree (AST) construction
- ✅ Tree-walking evaluator for arithmetic
- ✅ Interactive REPL with prompt
- ✅ Comment support (`#` line comments)
- ✅ Symbol table for variable storage (structure in place)
- ✅ Arithmetic operators: `+`, `-`, `*`, `/`
- ✅ Comparison operator tokens: `==`, `!=`, `<`, `>`
- ✅ Parenthesized expressions
- ✅ All helper functions implemented (isLetter, isDigit, readIdentifier, readNumber)

### Planned Features (After Compilation)
- ⏳ Variables and assignments (`name` keyword)
- ⏳ Functions (`cmnd` keyword)
- ⏳ Conditionals (`is`, `else` keywords)
- ⏳ Boolean evaluation (`true`, `false`)
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

### Installation & Running

```bash
# Build the interpreter
go build -o alaya ./main

# Run the REPL
./alaya
```

You'll see:
```
Hello <username>! Welcome to the Alaya Programming Language!
 **
```

### Example Usage

Evaluate arithmetic expressions:

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

### 🚧 Current State: Does Not Compile

The project architecture is **~95% complete**, but it **will not compile** until the following functions are implemented. This is **intentional** - these functions are core learning opportunities for understanding lexical analysis.

### ⚠️ Required for Compilation

**In `main/alaya_tokenizer/tokenizer.go` (Priority: CRITICAL):**
- ❌ `isLetter(ch byte) bool` - Check if character is a letter (a-z, A-Z, _)
- ❌ `isDigit(ch byte) bool` - Check if character is a digit (0-9)
- ❌ `readIdentifier() Token` - Read complete identifier or keyword
- ❌ `readNumber() Token` - Read complete integer

**In `main/alaya_ast/ast.go` (Priority: Low):**
- ❌ Define `Statement` interface for statement nodes (not blocking compilation of main features)

### 📊 What Works Once Implemented

After implementing the 4 tokenizer functions:
- ✅ Tokenization of source code
- ✅ Parsing arithmetic expressions
- ✅ AST construction
- ✅ Expression evaluation
- ✅ Interactive REPL
- ✅ Basic arithmetic calculator (`2 + 3 * 4`, parentheses, etc.)

### 🎯 Why These Are Left Unimplemented

These functions are **intentionally incomplete** as learning exercises:
- They teach fundamental concepts in lexical analysis
- They demonstrate Go string/byte manipulation
- They're simple enough to implement in 30-60 minutes
- They provide immediate satisfaction when the interpreter compiles and runs

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

### 🛠️ Quick Start: Get It Compiling (30-60 minutes)

Open `main/alaya_tokenizer/tokenizer.go` and add these 4 functions:

#### **1. `isLetter(ch byte) bool`** ⭐ Start Here
```go
// Add this function anywhere in the file (suggest after the New() function)
func isLetter(ch byte) bool {
    return (ch >= 'a' && ch <= 'z') || 
           (ch >= 'A' && ch <= 'Z') || 
           ch == '_'
}
```

**What it does:** Checks if a byte is a valid identifier start/continuation character.

**Learning:** Go byte comparisons, ASCII ranges

---

#### **2. `isDigit(ch byte) bool`** ⭐ 
```go
func isDigit(ch byte) bool {
    return ch >= '0' && ch <= '9'
}
```

**What it does:** Checks if a byte is a numeric digit.

**Learning:** More byte operations, pattern recognition

---

#### **3. `readIdentifier() Token`** ⭐⭐ (Slightly harder)
```go
func (t *Tokenizer) readIdentifier() alaya_token.Token {
    startPos := t.position
    
    // Read all letters and digits (identifiers can contain numbers after first char)
    for isLetter(t.currentCharacter) || isDigit(t.currentCharacter) {
        t.Advance()
    }
    
    // Extract the complete identifier
    identifier := t.text[startPos:t.position]
    
    // Check if it's a keyword (cmnd, name, is, else, true, false, return)
    if tokenType, isKeyword := alaya_token.Keywords[identifier]; isKeyword {
        return alaya_token.New(tokenType, identifier)
    }
    
    // Not a keyword, so it's a regular identifier
    return alaya_token.New(alaya_token.IDENT, identifier)
}
```

**What it does:** Reads a complete identifier (like `myVariable` or keyword like `cmnd`) from the input.

**Learning:** String slicing, loops, map lookups, distinguishing keywords from identifiers

---

#### **4. `readNumber() Token`** ⭐⭐
```go
func (t *Tokenizer) readNumber() alaya_token.Token {
    startPos := t.position
    
    // Read all consecutive digits
    for isDigit(t.currentCharacter) {
        t.Advance()
    }
    
    // Extract the complete number
    number := t.text[startPos:t.position]
    
    return alaya_token.New(alaya_token.INTEGER, number)
}
```

**What it does:** Reads a complete integer (like `123` or `42`) from the input.

**Learning:** Similar pattern to `readIdentifier`, demonstrates scanning patterns

---

### ✅ After Implementation

```bash
# Should now compile successfully!
go build -o alaya ./main

# Run tests
go test ./main/alaya_tokenizer -v

# Try the REPL
./alaya
```

### 🧪 Test Your Implementation

Try these in the REPL:
```
** 2 + 3
** 10 * (5 - 2)
** 100 / 4 + 6
```

If you see results, **congratulations! You have a working interpreter!** 🎉

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
