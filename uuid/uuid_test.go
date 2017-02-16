package uuid

import (
	"regexp"
	"testing"
)

const format = "^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$"

func TestNewID(t *testing.T) {
	u, err := NewID()
	if err != nil {
		t.Errorf("Expected to generate UUID without problems, error thrown: %s", err.Error())
		return
	}
	if u.Version() != 4 {
		t.Errorf("Expected to generate UUIDv4, given %d", u.Version())
	}
	if u.Variant() != ReservedRFC4122 {
		t.Errorf("Expected to generate UUIDv4 RFC4122 variant, given %x", u.Variant())
	}
	re := regexp.MustCompile(format)
	if !re.MatchString(u.String()) {
		t.Errorf("Expected string representation to be valid, given %s", u.String())
	}
}

func BenchmarkParseHex(b *testing.B) {
	s := "f3593cff-ee92-40df-4086-87825b523f13"
	for i := 0; i < b.N; i++ {
		_, err := ParseHex(s)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	b.ReportAllocs()
}
