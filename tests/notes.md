# Test suites

These are some notes on how to create and run test suites in Go.

## Writing subtests

There is an subtests example on [shapes/shapes_test.go](/interfaces/shapes/shapes\_test.go).

## Running a test suite

```bash
go test -v ./...
```

## Creating a coverage report

```bash
go test -coverprofile coverage.out -v ./...
```

## Viewing the coverage report with HTML

```bash
go tool cover -html=coverage.out
```
