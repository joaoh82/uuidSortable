package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// The UUID reserved variants.
const (
	ReservedRFC4122 byte = 0x40
)

// UUID type declared as a slice of bytes with 16 bits
type UUID [16]byte

// Pattern used to parse hex string representation of the UUID.
const hexPattern = "^(urn\\:uuid\\:)?\\{?([a-z0-9]{8})-([a-z0-9]{4})-" +
	"([1-5][a-z0-9]{3})-([a-z0-9]{4})-([a-z0-9]{12})\\}?$"

var re = regexp.MustCompile(hexPattern)

// ParseHex creates a UUID object from given hex string
// representation. Formats:
//     uuid.ParseHex("6ba7b814-9dad-11d1-80b4-00c04fd430c8")
//     uuid.ParseHex("{6ba7b814-9dad-11d1-80b4-00c04fd430c8}")
//     uuid.ParseHex("urn:uuid:6ba7b814-9dad-11d1-80b4-00c04fd430c8")
func ParseHex(s string) (u *UUID, err error) {
	md := re.FindStringSubmatch(s)
	if md == nil {
		err = errors.New("Invalid UUID string")
		return
	}
	hash := md[2] + md[3] + md[4] + md[5] + md[6]
	b, err := hex.DecodeString(hash)
	if err != nil {
		return
	}
	u = new(UUID)
	copy(u[:], b)
	return
}

// NewIDSortable generates a random UUID.
func NewIDSortable() (ut string, err error) {
	u := new(UUID)
	// Set all bits to randomly (pseudo) chosen values.
	_, err = rand.Read(u[:])
	if err != nil {
		return
	}
	u.setVariant(ReservedRFC4122)
	u.setVersion(4)

	// Adds the timestamp to the beggining of the generated UUID
	ut = strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + u.String()

	return
}

// NewID generates a random UUID.
func NewID() (u *UUID, err error) {
	u = new(UUID)
	// Set all bits to randomly (pseudo) chosen values.
	_, err = rand.Read(u[:])
	if err != nil {
		return
	}
	u.setVariant(ReservedRFC4122)
	u.setVersion(4)
	return
}

// Set the bits 6 and 7 of the clock_seq_hi_and_reserved to zero and one.
func (u *UUID) setVariant(v byte) {
	u[8] = (u[8] | ReservedRFC4122) & 0x7F
}

// Variant returns the UUID Variant, to determine the internal
// layout of the UUID.
func (u *UUID) Variant() byte {
	return ReservedRFC4122
}

// Set the bits 12 through 15 of the time_hi_and_version field to the
// 4-bit version number.
func (u *UUID) setVersion(v byte) {
	u[6] = (u[6] & 0xF) | (v << 4)
}

// Version returns the version of the algorithm used to
// generate the UUID sequence.
func (u *UUID) Version() uint {
	return uint(u[6] >> 4)
}

// Returns string version of the generated UUID.
func (u *UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}
