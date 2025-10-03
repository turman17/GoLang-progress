# 🧮 Calculator CLI (Go)

This is my **first project in Go**, created to practice the basics of the language and get comfortable with Go’s project structure, error handling, and testing.  

It’s a simple command-line calculator that supports **addition, subtraction, multiplication, and division**.

---

## 📂 Project Structure

```
calculator/
├── cmd/
│   └── calculator/
│       └── main.go        # Entry point for the CLI
├── internal/
│   └── calc/              # Core calculator logic
│       ├── evaluate.go    # Evaluates expressions
│       ├── parse.go       # Parses CLI arguments into numbers and operator
│       ├── errors.go      # Usage & error messages
│       └── calc_test.go   # Unit tests for the calculator
├── go.mod                 # Go module definition
└── README.md              # Documentation (this file)
```

- **`cmd/calculator/`** → where the `main.go` lives (the actual program you run).  
- **`internal/calc/`** → private logic for parsing and evaluating expressions.  
- **`calc_test.go`** → unit tests to ensure correctness.  

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/turman17/GoLang-progress.git
cd GoLang-progress/calculator
```

### 2. Run the calculator

```bash
go run ./cmd/calculator <number> <operator> <number>
```

#### Example:

```bash
go run ./cmd/calculator 3 + 5
```

✅ Output:
```
Result: 8.000
```

---

## ➕ Supported Operators

| Operator | Action        |
|----------|---------------|
| `+`      | Addition      |
| `-`      | Subtraction   |
| `*`      | Multiplication|
| `/`      | Division (with zero-check) |

---

## 🧪 Running Tests

Unit tests are included for the calculator logic.

Run all tests:

```bash
go test ./...
```

Verbose mode:

```bash
go test -v ./internal/calc
```

With coverage:

```bash
go test -cover ./internal/calc
```

---

## 🎯 What I Learned

- How to structure a Go project (`cmd/`, `internal/`)  
- Parsing CLI arguments with `os.Args`  
- Converting strings to numbers with `strconv`  
- Error handling and user-friendly messages  
- Writing unit tests using Go’s `testing` package  

---

## 📌 Next Steps

- Add more math operators (modulo `%`, power `^`)  
- Better input validation and error output  
- Explore integration with external packages  
- Add CI (GitHub Actions) to run tests automatically  

---

✍️ **Author:** [@turman17](https://github.com/turman17)  
📅 **Started:** October 2025  
🔰 **Note:** This is my first Go project and part of my learning journey.  