# go-agentic-ai

> Go language learning journey — Internship at **AgenticGokit**
> Learning Track: Go for Agentic AI

---

## About

This repository documents my weekly progress learning the Go programming language
during my internship at AgenticGokit. The goal is to build a strong Go foundation
and apply it to agentic AI systems.

**Reference:** [go.dev/doc](https://go.dev/doc) · [go.dev/ref/spec](https://go.dev/ref/spec) · [go.dev/tour](https://go.dev/tour)

---

## Weekly Progress

| Week | Focus | Status |
|------|-------|--------|
| Week 1 | Go environment setup, Tour of Go, modules, concurrency intro | ✅ Done |
| Week 2 | Go Language Specification — lexical elements, types, constants | ✅ Done |
| Week 3 | net/http, encoding/json, context, Anthropic API integration | 🔄 In Progress |
| Week 4 | Agent loop, tool use, multi-agent patterns | 🔜 Upcoming |

---

## Repository Structure

```
go-agentic-ai/
├── go.mod
├── README.md
├── .gitignore
│
├── week1/                  # Tour of Go — basics
│   ├── hello/              # Hello world, first module
│   ├── basics/             # Variables, loops, functions, structs
│   └── pointers/           # Pointers and memory
│
├── week2/                  # Go Language Specification study
│   ├── types/              # All type system examples
│   ├── literals/           # Integer, float, rune, string literals
│   └── constants/          # Typed vs untyped constants, iota
│
├── week3/                  # HTTP, JSON, context
│   ├── httpclient/         # net/http REST client
│   ├── jsonparse/          # encoding/json examples
│   └── context/            # context.Context timeout & cancel
│
└── week4/                  # Agentic AI
    ├── agent/              # Basic LLM agent loop
    └── tools/              # Tool use & function dispatch
```

---

## How to Run

```bash
# Clone the repository
git clone https://github.com/your-username/go-agentic-ai.git
cd go-agentic-ai

# Run any example (replace path as needed)
go run week1/hello/main.go
go run week2/types/main.go
go run week3/httpclient/main.go
```

---

## Resources

- [The Go Programming Language Docs](https://go.dev/doc/)
- [A Tour of Go](https://go.dev/tour/)
- [Go Language Specification](https://go.dev/ref/spec)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Standard Library](https://pkg.go.dev/std)
- [Anthropic API Docs](https://docs.anthropic.com)

---

*AgenticGokit Internship · Go for Agentic AI*
