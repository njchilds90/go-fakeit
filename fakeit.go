// Package fakeit provides a zero-dependency, seedable fake data generator for Go.
//
// It is designed to be deterministic, AI-agent-friendly, and safe for concurrent use.
// Every Generator instance is isolated — no global state is used.
//
// # Quick Start
//
//	g := fakeit.New(42) // seed for deterministic output
//	fmt.Println(g.Name())
//	fmt.Println(g.Email())
//	fmt.Println(g.UUID())
//
// # Determinism
//
// The same seed always produces the same sequence of values.
// This makes go-fakeit ideal for reproducible tests and AI training data pipelines.
package fakeit

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Generator is a seeded fake data generator. All methods are safe for sequential use.
// Create one with New or NewRandom.
type Generator struct {
	r *rand.Rand
}

// New creates a Generator with a fixed seed, producing deterministic output.
// Use the same seed to reproduce identical sequences.
//
//	g := fakeit.New(42)
func New(seed int64) *Generator {
	return &Generator{r: rand.New(rand.NewSource(seed))} //nolint:gosec
}

// NewRandom creates a Generator seeded from the current time (non-deterministic).
//
//	g := fakeit.NewRandom()
func NewRandom() *Generator {
	return New(time.Now().UnixNano())
}

// Seed resets the generator with a new seed, restarting its sequence.
func (g *Generator) Seed(seed int64) {
	g.r = rand.New(rand.NewSource(seed)) //nolint:gosec
}

// intn returns a random int in [0, n).
func (g *Generator) intn(n int) int {
	return g.r.Intn(n)
}

// pick returns a random element from the slice.
func (g *Generator) pick(s []string) string {
	return s[g.intn(len(s))]
}

// ─── Identity ────────────────────────────────────────────────────────────────

// FirstName returns a random first name.
func (g *Generator) FirstName() string { return g.pick(firstNames) }

// LastName returns a random last name.
func (g *Generator) LastName() string { return g.pick(lastNames) }

// Name returns a random full name (first + last).
func (g *Generator) Name() string {
	return g.FirstName() + " " + g.LastName()
}

// Username returns a random lowercase username (e.g. "blue_fox42").
func (g *Generator) Username() string {
	adj := strings.ToLower(g.pick(adjectives))
	noun := strings.ToLower(g.pick(animals))
	num := g.intn(100)
	return fmt.Sprintf("%s_%s%d", adj, noun, num)
}

// ─── Internet ────────────────────────────────────────────────────────────────

// Email returns a random email address.
func (g *Generator) Email() string {
	user := strings.ToLower(g.FirstName()) + "." + strings.ToLower(g.LastName())
	domain := g.pick(emailDomains)
	return fmt.Sprintf("%s@%s", user, domain)
}

// URL returns a random URL.
func (g *Generator) URL() string {
	scheme := g.pick([]string{"https", "http"})
	sub := g.pick([]string{"www", "api", "app", "dev", "blog"})
	domain := g.pick(domains)
	path := strings.ToLower(g.pick(adjectives)) + "-" + strings.ToLower(g.pick(nouns))
	return fmt.Sprintf("%s://%s.%s/%s", scheme, sub, domain, path)
}

// IPv4 returns a random IPv4 address string.
func (g *Generator) IPv4() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		g.intn(256), g.intn(256), g.intn(256), g.intn(256))
}

// IPv6 returns a random IPv6 address string.
func (g *Generator) IPv6() string {
	parts := make([]string, 8)
	for i := range parts {
		parts[i] = fmt.Sprintf("%04x", g.intn(0xffff+1))
	}
	return strings.Join(parts, ":")
}

// UserAgent returns a random HTTP User-Agent string.
func (g *Generator) UserAgent() string { return g.pick(userAgents) }

// ─── Text ─────────────────────────────────────────────────────────────────────

// Word returns a random English word.
func (g *Generator) Word() string { return g.pick(nouns) }

// Sentence returns a random sentence of n words.
func (g *Generator) Sentence(words int) string {
	if words <= 0 {
		words = 6
	}
	ws := make([]string, words)
	ws[0] = strings.Title(g.pick(adjectives)) //nolint:staticcheck
	for i := 1; i < words; i++ {
		ws[i] = g.pick(nouns)
	}
	return strings.Join(ws, " ") + "."
}

// Paragraph returns a random paragraph of n sentences.
func (g *Generator) Paragraph(sentences int) string {
	if sentences <= 0 {
		sentences = 3
	}
	ss := make([]string, sentences)
	for i := range ss {
		ss[i] = g.Sentence(g.intn(6) + 4)
	}
	return strings.Join(ss, " ")
}

// LoremIpsum returns classic Lorem Ipsum placeholder text of approximately n words.
func (g *Generator) LoremIpsum(words int) string {
	lorem := strings.Fields(loremText)
	result := make([]string, words)
	for i := range result {
		result[i] = lorem[g.intn(len(lorem))]
	}
	return strings.Join(result, " ") + "."
}

// Slug returns a URL-safe slug (e.g. "quick-brown-fox").
func (g *Generator) Slug() string {
	return strings.ToLower(g.pick(adjectives)) + "-" +
		strings.ToLower(g.pick(adjectives)) + "-" +
		strings.ToLower(g.pick(nouns))
}

// Hashtag returns a random hashtag string.
func (g *Generator) Hashtag() string {
	return "#" + strings.ToLower(g.pick(adjectives)) + strings.Title(g.pick(nouns)) //nolint:staticcheck
}

// ─── Numbers ─────────────────────────────────────────────────────────────────

// Int returns a random int in [min, max].
func (g *Generator) Int(min, max int) int {
	if max <= min {
		return min
	}
	return min + g.intn(max-min+1)
}

// Float64 returns a random float64 in [min, max).
func (g *Generator) Float64(min, max float64) float64 {
	return min + g.r.Float64()*(max-min)
}

// Bool returns a random boolean.
func (g *Generator) Bool() bool { return g.intn(2) == 1 }

// Digit returns a random digit character ('0'–'9').
func (g *Generator) Digit() byte { return byte('0' + g.intn(10)) }

// ─── Identifiers ─────────────────────────────────────────────────────────────

// UUID returns a random UUID v4 string.
func (g *Generator) UUID() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = byte(g.intn(256))
	}
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant bits
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// HexString returns a random hex string of length n.
func (g *Generator) HexString(n int) string {
	const hexChars = "0123456789abcdef"
	b := make([]byte, n)
	for i := range b {
		b[i] = hexChars[g.intn(len(hexChars))]
	}
	return string(b)
}

// NumericalString returns a string of n random digits.
func (g *Generator) NumericalString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = g.Digit()
	}
	return string(b)
}

// ─── Dates & Times ───────────────────────────────────────────────────────────

// Date returns a random time.Time between start and end.
func (g *Generator) Date(start, end time.Time) time.Time {
	delta := end.Unix() - start.Unix()
	if delta <= 0 {
		return start
	}
	sec := g.r.Int63n(delta)
	return time.Unix(start.Unix()+sec, 0).UTC()
}

// PastDate returns a random date in the past N days.
func (g *Generator) PastDate(days int) time.Time {
	now := time.Now().UTC()
	return g.Date(now.AddDate(0, 0, -days), now)
}

// FutureDate returns a random date in the next N days.
func (g *Generator) FutureDate(days int) time.Time {
	now := time.Now().UTC()
	return g.Date(now, now.AddDate(0, 0, days))
}

// ISODate returns a random ISO 8601 date string (YYYY-MM-DD).
func (g *Generator) ISODate() string {
	t := g.PastDate(3650) // up to ~10 years ago
	return t.Format("2006-01-02")
}

// ─── Location ────────────────────────────────────────────────────────────────

// City returns a random city name.
func (g *Generator) City() string { return g.pick(cities) }

// Country returns a random country name.
func (g *Generator) Country() string { return g.pick(countries) }

// CountryCode returns a random ISO 3166-1 alpha-2 country code.
func (g *Generator) CountryCode() string { return g.pick(countryCodes) }

// Latitude returns a random latitude in [-90, 90].
func (g *Generator) Latitude() float64 { return g.Float64(-90, 90) }

// Longitude returns a random longitude in [-180, 180].
func (g *Generator) Longitude() float64 { return g.Float64(-180, 180) }

// TimeZone returns a random IANA timezone string.
func (g *Generator) TimeZone() string { return g.pick(timeZones) }

// ─── Payment ─────────────────────────────────────────────────────────────────

// CreditCardNumber returns a fake 16-digit credit card number string.
func (g *Generator) CreditCardNumber() string {
	return g.NumericalString(4) + "-" +
		g.NumericalString(4) + "-" +
		g.NumericalString(4) + "-" +
		g.NumericalString(4)
}

// CreditCardExpiry returns a fake card expiry date (MM/YY format).
func (g *Generator) CreditCardExpiry() string {
	month := g.intn(12) + 1
	year := time.Now().Year()%100 + g.intn(5) + 1
	return fmt.Sprintf("%02d/%02d", month, year)
}

// CreditCardCVV returns a fake 3-digit CVV string.
func (g *Generator) CreditCardCVV() string { return g.NumericalString(3) }

// Currency returns a random ISO 4217 currency code.
func (g *Generator) Currency() string { return g.pick(currencies) }

// Price returns a random price string like "19.99".
func (g *Generator) Price(min, max float64) string {
	return fmt.Sprintf("%.2f", g.Float64(min, max))
}

// ─── Color ───────────────────────────────────────────────────────────────────

// HexColor returns a random hex color string (e.g. "#a3f0c2").
func (g *Generator) HexColor() string {
	return fmt.Sprintf("#%06x", g.intn(0xffffff+1))
}

// RGBColor returns a random RGB color as [3]uint8.
func (g *Generator) RGBColor() [3]uint8 {
	return [3]uint8{uint8(g.intn(256)), uint8(g.intn(256)), uint8(g.intn(256))}
}

// ColorName returns a random English color name.
func (g *Generator) ColorName() string { return g.pick(colorNames) }

// ─── Company ─────────────────────────────────────────────────────────────────

// Company returns a random company name.
func (g *Generator) Company() string {
	return g.pick(companyPrefixes) + " " + g.pick(companySuffixes)
}

// JobTitle returns a random job title.
func (g *Generator) JobTitle() string {
	return g.pick(jobLevels) + " " + g.pick(jobRoles)
}

// Industry returns a random industry name.
func (g *Generator) Industry() string { return g.pick(industries) }

// ─── File ────────────────────────────────────────────────────────────────────

// FileName returns a random filename with extension (e.g. "report_draft.pdf").
func (g *Generator) FileName() string {
	base := strings.ToLower(g.pick(nouns)) + "_" + strings.ToLower(g.pick(adjectives))
	ext := g.pick(fileExtensions)
	return base + ext
}

// FileExtension returns a random file extension (e.g. ".json").
func (g *Generator) FileExtension() string { return g.pick(fileExtensions) }

// MimeType returns a random MIME type string.
func (g *Generator) MimeType() string { return g.pick(mimeTypes) }

// ─── Network / System ────────────────────────────────────────────────────────

// MacAddress returns a random MAC address string.
func (g *Generator) MacAddress() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = byte(g.intn(256))
	}
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		b[0], b[1], b[2], b[3], b[4], b[5])
}

// Port returns a random port number in [1024, 65535].
func (g *Generator) Port() int { return g.Int(1024, 65535) }

// HTTPMethod returns a random HTTP method string.
func (g *Generator) HTTPMethod() string { return g.pick(httpMethods) }

// HTTPStatusCode returns a random HTTP status code.
func (g *Generator) HTTPStatusCode() int { return g.pick(httpStatusCodes) }

// ─── Structured Records ──────────────────────────────────────────────────────

// Person holds fake identity data.
type Person struct {
	FirstName string
	LastName  string
	FullName  string
	Email     string
	Username  string
	JobTitle  string
	Company   string
}

// PersonRecord returns a fully populated Person struct.
func (g *Generator) PersonRecord() Person {
	first := g.FirstName()
	last := g.LastName()
	return Person{
		FirstName: first,
		LastName:  last,
		FullName:  first + " " + last,
		Email:     strings.ToLower(first) + "." + strings.ToLower(last) + "@" + g.pick(emailDomains),
		Username:  g.Username(),
		JobTitle:  g.JobTitle(),
		Company:   g.Company(),
	}
}

// Address holds fake location data.
type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
	Lat     float64
	Lon     float64
}

// AddressRecord returns a fully populated Address struct.
func (g *Generator) AddressRecord() Address {
	num := g.Int(1, 9999)
	street := g.pick(streetNames)
	suffix := g.pick(streetSuffixes)
	return Address{
		Street:  fmt.Sprintf("%d %s %s", num, street, suffix),
		City:    g.City(),
		Country: g.Country(),
		ZipCode: g.NumericalString(5),
		Lat:     g.Latitude(),
		Lon:     g.Longitude(),
	}
}

// ─── Slice helpers ───────────────────────────────────────────────────────────

// Sample returns n unique random elements from slice s.
// If n >= len(s), returns a shuffled copy of s.
func (g *Generator) Sample(s []string, n int) []string {
	cp := make([]string, len(s))
	copy(cp, s)
	g.r.Shuffle(len(cp), func(i, j int) { cp[i], cp[j] = cp[j], cp[i] })
	if n > len(cp) {
		n = len(cp)
	}
	return cp[:n]
}

// Shuffle returns a shuffled copy of the string slice.
func (g *Generator) Shuffle(s []string) []string {
	return g.Sample(s, len(s))
}

// PickOne returns a random element from the provided slice.
func (g *Generator) PickOne(s []string) string {
	if len(s) == 0 {
		return ""
	}
	return g.pick(s)
}
