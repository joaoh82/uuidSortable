package main

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"time"
)

// The UUID reserved variants.
const (
	ReservedRFC4122 byte = 0x40
)

// UUID type declared as a slice of bytes with 16 bits
type UUID [16]byte

// NewIDSortable generates a random UUID.
func NewIDSortable() (ut string, err error) {
	u := new(UUID)
	// Set all bits to randomly (or pseudo-randomly) chosen values.
	_, err = rand.Read(u[:])
	if err != nil {
		return
	}
	u.setVariant(ReservedRFC4122)
	u.setVersion(4)

	ut = strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + u.String()

	return
}

// NewID generates a random UUID.
func NewID() (u *UUID, err error) {
	u = new(UUID)
	// Set all bits to randomly (or pseudo-randomly) chosen values.
	_, err = rand.Read(u[:])
	if err != nil {
		return
	}
	u.setVariant(ReservedRFC4122)
	u.setVersion(4)
	return
}

// Set the two most significant bits (bits 6 and 7) of the
// clock_seq_hi_and_reserved to zero and one, respectively.
func (u *UUID) setVariant(v byte) {
	u[8] = (u[8] | ReservedRFC4122) & 0x7F
}

// Variant returns the UUID Variant, which determines the internal
// layout of the UUID.
func (u *UUID) Variant() byte {
	return ReservedRFC4122
}

// Set the four most significant bits (bits 12 through 15) of the
// time_hi_and_version field to the 4-bit version number.
func (u *UUID) setVersion(v byte) {
	u[6] = (u[6] & 0xF) | (v << 4)
}

// Version returns a version number of the algorithm used to
// generate the UUID sequence.
func (u *UUID) Version() uint {
	return uint(u[6] >> 4)
}

// Returns unparsed version of the generated UUID sequence.
func (u *UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}
