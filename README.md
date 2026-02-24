# go-fakeit

> Zero-dependency, seedable fake data generator for Go — deterministic, AI-agent-friendly, inspired by Faker.js and Python's Faker.

[![CI](https://github.com/njchilds90/go-fakeit/actions/workflows/ci.yml/badge.svg)](https://github.com/njchilds90/go-fakeit/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/njchilds90/go-fakeit.svg)](https://pkg.go.dev/github.com/njchilds90/go-fakeit)
[![Go Report Card](https://goreportcard.com/badge/github.com/njchilds90/go-fakeit)](https://goreportcard.com/report/github.com/njchilds90/go-fakeit)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## Why go-fakeit?

Python has [Faker](https://github.com/joke2k/faker) (15M+ monthly downloads).  
JavaScript has [@faker-js/faker](https://github.com/faker-js/faker) (5M+ weekly downloads).  
Go deserves an equivalent that is idiomatic, deterministic, and AI-agent-safe.

**go-fakeit** is:
- 🔒 **Deterministic** — same seed = same output, always
- 🧩 **Zero dependencies** — pure stdlib only
- 🤖 **AI-agent friendly** — no global state, pure functions, predictable types
- ⚡ **Rich API** — 40+ methods across 10+ categories, plus structured record types
- 🧪 **Well tested** — race-condition safe, edge-case tested, CI on 3 Go versions

## Install
```bash
go get github.com/njchilds90/go-fakeit
```

## Quick Start
```go
package main

import (
    "fmt"
    "github.com/njchilds90/go-fakeit"
)

func main() {
    g := fakeit.New(42) // deterministic

    fmt.Println(g.Name())             // e.g. "Alice Johnson"
    fmt.Println(g.Email())            // e.g. "alice.johnson@gmail.com"
    fmt.Println(g.UUID())             // e.g. "3f2c1a9b-..."
    fmt.Println(g.IPv4())             // e.g. "192.168.10.5"
    fmt.Println(g.HexColor())         // e.g. "#a3f0c2"
    fmt.Println(g.Price(10, 500))     // e.g. "249.99"
    fmt.Println(g.ISODate())          // e.g. "2021-07-14"

    p := g.PersonRecord()
    fmt.Println(p.FullName, p.JobTitle, p.Company)

    a := g.AddressRecord()
    fmt.Println(a.Street, a.City, a.Country)
}
```

## API Reference

| Category | Methods |
|---|---|
| **Identity** | `FirstName`, `LastName`, `Name`, `Username` |
| **Internet** | `Email`, `URL`, `IPv4`, `IPv6`, `UserAgent` |
| **Text** | `Word`, `Sentence`, `Paragraph`, `LoremIpsum`, `Slug`, `Hashtag` |
| **Numbers** | `Int`, `Float64`, `Bool`, `Digit` |
| **Identifiers** | `UUID`, `HexString`, `NumericalString` |
| **Dates** | `Date`, `PastDate`, `FutureDate`, `ISODate` |
| **Location** | `City`, `Country`, `CountryCode`, `Latitude`, `Longitude`, `TimeZone` |
| **Payment** | `CreditCardNumber`, `CreditCardExpiry`, `CreditCardCVV`, `Currency`, `Price` |
| **Color** | `HexColor`, `RGBColor`, `ColorName` |
| **Company** | `Company`, `JobTitle`, `Industry` |
| **File** | `FileName`, `FileExtension`, `MimeType` |
| **Network** | `MacAddress`, `Port`, `HTTPMethod`, `HTTPStatusCode` |
| **Records** | `PersonRecord`, `AddressRecord` |
| **Slice helpers** | `Sample`, `Shuffle`, `PickOne` |

Full docs: [pkg.go.dev/github.com/njchilds90/go-fakeit](https://pkg.go.dev/github.com/njchilds90/go-fakeit)

## Determinism
```go
g1 := fakeit.New(42)
g2 := fakeit.New(42)
// g1.Name() == g2.Name() always
```

For non-deterministic output:
```go
g := fakeit.NewRandom()
```

## License

MIT — see [LICENSE](LICENSE).
```

---


```
   ## go-fakeit v1.0.0

   Initial production release of go-fakeit — a zero-dependency, seedable fake data generator for Go.

   ### Highlights
   - 40+ methods across 10+ categories
   - Fully deterministic (seed-based)
   - Zero external dependencies
   - AI-agent and testing pipeline friendly
   - Structured record types: PersonRecord, AddressRecord
   - Full test suite with race detector
   - CI on Go 1.21, 1.22, 1.23

   ### Install
   go get github.com/njchilds90/go-fakeit@v1.0.0
