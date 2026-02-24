package fakeit_test

import (
	"strings"
	"testing"
	"time"

	"github.com/njchilds90/go-fakeit"
)

// ─── Determinism ─────────────────────────────────────────────────────────────

func TestDeterminism(t *testing.T) {
	g1 := fakeit.New(42)
	g2 := fakeit.New(42)

	for i := 0; i < 100; i++ {
		n1, n2 := g1.Name(), g2.Name()
		if n1 != n2 {
			t.Fatalf("seed 42 produced different Name() at call %d: %q vs %q", i, n1, n2)
		}
	}
}

func TestSeedReset(t *testing.T) {
	g := fakeit.New(99)
	first := g.Name()
	g.Seed(99)
	second := g.Name()
	if first != second {
		t.Fatalf("Seed() did not reset sequence: %q vs %q", first, second)
	}
}

func TestDifferentSeeds(t *testing.T) {
	g1 := fakeit.New(1)
	g2 := fakeit.New(2)
	// It's astronomically unlikely the first UUID matches with different seeds.
	u1, u2 := g1.UUID(), g2.UUID()
	if u1 == u2 {
		t.Fatalf("different seeds produced identical UUID: %q", u1)
	}
}

// ─── Identity ────────────────────────────────────────────────────────────────

func TestFirstName(t *testing.T) {
	g := fakeit.New(1)
	for i := 0; i < 20; i++ {
		v := g.FirstName()
		if v == "" {
			t.Fatal("FirstName returned empty string")
		}
	}
}

func TestLastName(t *testing.T) {
	g := fakeit.New(2)
	for i := 0; i < 20; i++ {
		v := g.LastName()
		if v == "" {
			t.Fatal("LastName returned empty string")
		}
	}
}

func TestName(t *testing.T) {
	g := fakeit.New(3)
	name := g.Name()
	parts := strings.Fields(name)
	if len(parts) < 2 {
		t.Fatalf("Name() should have at least 2 parts, got %q", name)
	}
}

func TestUsername(t *testing.T) {
	g := fakeit.New(4)
	for i := 0; i < 20; i++ {
		u := g.Username()
		if !strings.Contains(u, "_") {
			t.Fatalf("Username %q missing underscore separator", u)
		}
	}
}

// ─── Internet ────────────────────────────────────────────────────────────────

func TestEmail(t *testing.T) {
	g := fakeit.New(5)
	for i := 0; i < 20; i++ {
		e := g.Email()
		if !strings.Contains(e, "@") {
			t.Fatalf("Email %q missing @", e)
		}
		if !strings.Contains(e, ".") {
			t.Fatalf("Email %q missing domain dot", e)
		}
	}
}

func TestURL(t *testing.T) {
	g := fakeit.New(6)
	for i := 0; i < 20; i++ {
		u := g.URL()
		if !strings.HasPrefix(u, "http") {
			t.Fatalf("URL %q does not start with http", u)
		}
	}
}

func TestIPv4(t *testing.T) {
	g := fakeit.New(7)
	for i := 0; i < 20; i++ {
		ip := g.IPv4()
		parts := strings.Split(ip, ".")
		if len(parts) != 4 {
			t.Fatalf("IPv4 %q does not have 4 octets", ip)
		}
	}
}

func TestIPv6(t *testing.T) {
	g := fakeit.New(8)
	for i := 0; i < 20; i++ {
		ip := g.IPv6()
		parts := strings.Split(ip, ":")
		if len(parts) != 8 {
			t.Fatalf("IPv6 %q does not have 8 groups", ip)
		}
	}
}

func TestUserAgent(t *testing.T) {
	g := fakeit.New(9)
	ua := g.UserAgent()
	if ua == "" {
		t.Fatal("UserAgent returned empty string")
	}
}

// ─── Text ─────────────────────────────────────────────────────────────────────

func TestWord(t *testing.T) {
	g := fakeit.New(10)
	w := g.Word()
	if w == "" {
		t.Fatal("Word returned empty string")
	}
}

func TestSentence(t *testing.T) {
	g := fakeit.New(11)
	s := g.Sentence(5)
	if !strings.HasSuffix(s, ".") {
		t.Fatalf("Sentence does not end with '.': %q", s)
	}
	if len(strings.Fields(s)) != 5 {
		t.Fatalf("Sentence(5) should have 5 words, got: %q", s)
	}
}

func TestParagraph(t *testing.T) {
	g := fakeit.New(12)
	p := g.Paragraph(3)
	if p == "" {
		t.Fatal("Paragraph returned empty string")
	}
}

func TestLoremIpsum(t *testing.T) {
	g := fakeit.New(13)
	lorem := g.LoremIpsum(10)
	words := strings.Fields(lorem)
	// Words + trailing period means at least 10 words before period
	if len(words) < 10 {
		t.Fatalf("LoremIpsum(10) returned too few words: %q", lorem)
	}
}

func TestSlug(t *testing.T) {
	g := fakeit.New(14)
	for i := 0; i < 20; i++ {
		sl := g.Slug()
		if strings.Contains(sl, " ") {
			t.Fatalf("Slug %q contains space", sl)
		}
		parts := strings.Split(sl, "-")
		if len(parts) != 3 {
			t.Fatalf("Slug %q should have 3 hyphen-separated parts", sl)
		}
	}
}

func TestHashtag(t *testing.T) {
	g := fakeit.New(15)
	h := g.Hashtag()
	if !strings.HasPrefix(h, "#") {
		t.Fatalf("Hashtag %q does not start with #", h)
	}
}

// ─── Numbers ─────────────────────────────────────────────────────────────────

func TestInt(t *testing.T) {
	g := fakeit.New(16)
	for i := 0; i < 100; i++ {
		v := g.Int(10, 20)
		if v < 10 || v > 20 {
			t.Fatalf("Int(10,20) returned %d out of range", v)
		}
	}
}

func TestFloat64(t *testing.T) {
	g := fakeit.New(17)
	for i := 0; i < 100; i++ {
		v := g.Float64(0, 1)
		if v < 0 || v >= 1 {
			t.Fatalf("Float64(0,1) returned %f out of range", v)
		}
	}
}

func TestBool(t *testing.T) {
	g := fakeit.New(18)
	seen := map[bool]bool{}
	for i := 0; i < 100; i++ {
		seen[g.Bool()] = true
	}
	if !seen[true] || !seen[false] {
		t.Fatal("Bool() never returned both true and false in 100 calls")
	}
}

// ─── UUID ─────────────────────────────────────────────────────────────────────

func TestUUID(t *testing.T) {
	g := fakeit.New(19)
	seen := map[string]bool{}
	for i := 0; i < 100; i++ {
		u := g.UUID()
		parts := strings.Split(u, "-")
		if len(parts) != 5 {
			t.Fatalf("UUID %q wrong format", u)
		}
		if seen[u] {
			t.Fatalf("UUID collision at iteration %d: %q", i, u)
		}
		seen[u] = true
	}
}

func TestHexString(t *testing.T) {
	g := fakeit.New(20)
	for _, n := range []int{4, 8, 16, 32} {
		h := g.HexString(n)
		if len(h) != n {
			t.Fatalf("HexString(%d) returned length %d", n, len(h))
		}
	}
}

func TestNumericalString(t *testing.T) {
	g := fakeit.New(21)
	ns := g.NumericalString(10)
	if len(ns) != 10 {
		t.Fatalf("NumericalString(10) length = %d", len(ns))
	}
	for _, c := range ns {
		if c < '0' || c > '9' {
			t.Fatalf("NumericalString contains non-digit: %c", c)
		}
	}
}

// ─── Dates ───────────────────────────────────────────────────────────────────

func TestDate(t *testing.T) {
	g := fakeit.New(22)
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 20; i++ {
		d := g.Date(start, end)
		if d.Before(start) || d.After(end) {
			t.Fatalf("Date out of range: %v", d)
		}
	}
}

func TestISODate(t *testing.T) {
	g := fakeit.New(23)
	d := g.ISODate()
	if len(d) != 10 || d[4] != '-' || d[7] != '-' {
		t.Fatalf("ISODate format wrong: %q", d)
	}
}

// ─── Location ────────────────────────────────────────────────────────────────

func TestCity(t *testing.T) {
	g := fakeit.New(24)
	c := g.City()
	if c == "" {
		t.Fatal("City returned empty string")
	}
}

func TestLatLon(t *testing.T) {
	g := fakeit.New(25)
	for i := 0; i < 20; i++ {
		lat := g.Latitude()
		lon := g.Longitude()
		if lat < -90 || lat > 90 {
			t.Fatalf("Latitude %f out of range", lat)
		}
		if lon < -180 || lon > 180 {
			t.Fatalf("Longitude %f out of range", lon)
		}
	}
}

// ─── Payment ─────────────────────────────────────────────────────────────────

func TestCreditCardNumber(t *testing.T) {
	g := fakeit.New(26)
	cc := g.CreditCardNumber()
	parts := strings.Split(cc, "-")
	if len(parts) != 4 {
		t.Fatalf("CreditCardNumber %q wrong format", cc)
	}
}

func TestPrice(t *testing.T) {
	g := fakeit.New(27)
	p := g.Price(1, 999)
	if p == "" {
		t.Fatal("Price returned empty string")
	}
}

// ─── Color ───────────────────────────────────────────────────────────────────

func TestHexColor(t *testing.T) {
	g := fakeit.New(28)
	for i := 0; i < 20; i++ {
		c := g.HexColor()
		if !strings.HasPrefix(c, "#") || len(c) != 7 {
			t.Fatalf("HexColor %q wrong format", c)
		}
	}
}

func TestRGBColor(t *testing.T) {
	g := fakeit.New(29)
	rgb := g.RGBColor()
	_ = rgb // just ensure it compiles and doesn't panic
}

// ─── Company ─────────────────────────────────────────────────────────────────

func TestCompany(t *testing.T) {
	g := fakeit.New(30)
	c := g.Company()
	if len(strings.Fields(c)) < 2 {
		t.Fatalf("Company %q too short", c)
	}
}

func TestJobTitle(t *testing.T) {
	g := fakeit.New(31)
	j := g.JobTitle()
	if j == "" {
		t.Fatal("JobTitle returned empty string")
	}
}

// ─── File ────────────────────────────────────────────────────────────────────

func TestFileName(t *testing.T) {
	g := fakeit.New(32)
	f := g.FileName()
	if !strings.Contains(f, ".") {
		t.Fatalf("FileName %q has no extension", f)
	}
}

func TestMimeType(t *testing.T) {
	g := fakeit.New(33)
	m := g.MimeType()
	if !strings.Contains(m, "/") {
		t.Fatalf("MimeType %q wrong format", m)
	}
}

// ─── Network ─────────────────────────────────────────────────────────────────

func TestMacAddress(t *testing.T) {
	g := fakeit.New(34)
	mac := g.MacAddress()
	parts := strings.Split(mac, ":")
	if len(parts) != 6 {
		t.Fatalf("MacAddress %q wrong format", mac)
	}
}

func TestPort(t *testing.T) {
	g := fakeit.New(35)
	for i := 0; i < 50; i++ {
		p := g.Port()
		if p < 1024 || p > 65535 {
			t.Fatalf("Port %d out of range", p)
		}
	}
}

func TestHTTPMethod(t *testing.T) {
	g := fakeit.New(36)
	valid := map[string]bool{"GET": true, "POST": true, "PUT": true, "PATCH": true, "DELETE": true, "HEAD": true, "OPTIONS": true}
	for i := 0; i < 20; i++ {
		m := g.HTTPMethod()
		if !valid[m] {
			t.Fatalf("HTTPMethod returned invalid value: %q", m)
		}
	}
}

// ─── Structured Records ──────────────────────────────────────────────────────

func TestPersonRecord(t *testing.T) {
	g := fakeit.New(37)
	p := g.PersonRecord()
	if p.FullName == "" {
		t.Fatal("PersonRecord.FullName empty")
	}
	if !strings.Contains(p.Email, "@") {
		t.Fatalf("PersonRecord.Email invalid: %q", p.Email)
	}
	if p.JobTitle == "" {
		t.Fatal("PersonRecord.JobTitle empty")
	}
}

func TestAddressRecord(t *testing.T) {
	g := fakeit.New(38)
	a := g.AddressRecord()
	if a.Street == "" {
		t.Fatal("AddressRecord.Street empty")
	}
	if a.City == "" {
		t.Fatal("AddressRecord.City empty")
	}
	if a.Lat < -90 || a.Lat > 90 {
		t.Fatalf("AddressRecord.Lat out of range: %f", a.Lat)
	}
}

// ─── Slice Helpers ───────────────────────────────────────────────────────────

func TestSample(t *testing.T) {
	g := fakeit.New(39)
	source := []string{"a", "b", "c", "d", "e"}
	result := g.Sample(source, 3)
	if len(result) != 3 {
		t.Fatalf("Sample(3) returned %d elements", len(result))
	}
}

func TestSampleExceedsLength(t *testing.T) {
	g := fakeit.New(40)
	source := []string{"x", "y"}
	result := g.Sample(source, 10)
	if len(result) != len(source) {
		t.Fatalf("Sample capped to %d, expected %d", len(result), len(source))
	}
}

func TestPickOne(t *testing.T) {
	g := fakeit.New(41)
	result := g.PickOne([]string{"alpha", "beta", "gamma"})
	if result == "" {
		t.Fatal("PickOne returned empty string")
	}
}

func TestPickOneEmpty(t *testing.T) {
	g := fakeit.New(42)
	result := g.PickOne([]string{})
	if result != "" {
		t.Fatalf("PickOne on empty slice should return empty, got %q", result)
	}
}

func TestShuffle(t *testing.T) {
	g := fakeit.New(43)
	source := []string{"a", "b", "c", "d", "e"}
	result := g.Shuffle(source)
	if len(result) != len(source) {
		t.Fatalf("Shuffle length mismatch: %d vs %d", len(result), len(source))
	}
}

// ─── NewRandom ───────────────────────────────────────────────────────────────

func TestNewRandom(t *testing.T) {
	g := fakeit.NewRandom()
	v := g.Name()
	if v == "" {
		t.Fatal("NewRandom generator Name() returned empty")
	}
}

// ─── Edge cases ──────────────────────────────────────────────────────────────

func TestIntEqualMinMax(t *testing.T) {
	g := fakeit.New(44)
	v := g.Int(5, 5)
	if v != 5 {
		t.Fatalf("Int(5,5) should return 5, got %d", v)
	}
}

func TestSentenceZeroWords(t *testing.T) {
	g := fakeit.New(45)
	s := g.Sentence(0)
	if s == "" {
		t.Fatal("Sentence(0) returned empty, expected default behavior")
	}
}
