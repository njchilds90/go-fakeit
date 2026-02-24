# Contributing to go-fakeit

Thank you for your interest in contributing!

## Ground Rules

- All public functions must have GoDoc comments.
- All new features must include table-driven tests.
- No external dependencies — this is a zero-dependency library.
- Keep functions pure and stateless where possible.
- Do not use global state.

## How to Contribute

1. Fork the repository.
2. Create a branch: `git checkout -b feature/my-feature`
3. Make your changes with tests.
4. Run `go test ./...` and `go vet ./...` locally.
5. Open a Pull Request against `main`.

## Adding New Data

Word lists and lookup tables live in `data.go`.
If you add a new category, add it there and keep lists alphabetically sorted within categories.

## Coding Style

- Follow standard Go formatting (`gofmt`).
- Keep function signatures minimal and predictable.
- Return values, not errors, for pure data generation methods.
- Use `context.Context` only if a method may block.

## Reporting Issues

Open a GitHub Issue with a minimal reproduction case.
