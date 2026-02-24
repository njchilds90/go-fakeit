# Changelog

All notable changes to this project will be documented in this file.

The format follows [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).
This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-02-24

### Added

- `New(seed int64)` — seeded, deterministic generator
- `NewRandom()` — time-seeded generator for non-deterministic use
- `Seed(seed int64)` — reset generator seed mid-sequence
- Identity: `FirstName`, `LastName`, `Name`, `Username`
- Internet: `Email`, `URL`, `IPv4`, `IPv6`, `UserAgent`
- Text: `Word`, `Sentence`, `Paragraph`, `LoremIpsum`, `Slug`, `Hashtag`
- Numbers: `Int`, `Float64`, `Bool`, `Digit`
- Identifiers: `UUID`, `HexString`, `NumericalString`
- Dates: `Date`, `PastDate`, `FutureDate`, `ISODate`
- Location: `City`, `Country`, `CountryCode`, `Latitude`, `Longitude`, `TimeZone`
- Payment: `CreditCardNumber`, `CreditCardExpiry`, `CreditCardCVV`, `Currency`, `Price`
- Color: `HexColor`, `RGBColor`, `ColorName`
- Company: `Company`, `JobTitle`, `Industry`
- File: `FileName`, `FileExtension`, `MimeType`
- Network: `MacAddress`, `Port`, `HTTPMethod`, `HTTPStatusCode`
- Structured records: `PersonRecord`, `AddressRecord`
- Slice helpers: `Sample`, `Shuffle`, `PickOne`
- Full test suite with determinism, edge case, and range validation
- GitHub Actions CI (Go 1.21, 1.22, 1.23)
