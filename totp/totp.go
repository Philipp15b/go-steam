package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"time"
)

// ErrInvalidSharedSecret is returned when shared secret isn't in base64 form
var ErrInvalidSharedSecret error = errors.New("invalid base64 shared secret")

const (
	// Range of possible chars for auth code.
	chars    string = "23456789BCDFGHJKMNPQRTVWXY"
	charsLen int    = len(chars)
)

// Totp is structure holding data needed for generating TOTP token.
type Totp struct {
	sharedSecret string
	Time         time.Time
}

// NewTotp creates new Totp structure with current time.
func NewTotp(sharedSecret string) *Totp {
	return &Totp{sharedSecret, time.Now()}
}

// NewTimedTotp creates new Totp structure with custom time.
func NewTimedTotp(sharedSecret string, time time.Time) *Totp {
	return &Totp{sharedSecret, time}
}

// GenerateCode generates Steam TOTP code which is always 5 symbols.
func (totp *Totp) GenerateCode() (string, error) {
	return GenerateTotpCode(totp.sharedSecret, totp.Time)
}

// SharedSecret returns shared secret of Totp structure.
func (totp *Totp) SharedSecret() string {
	return totp.sharedSecret
}

// Function below is originally made by https://github.com/fortis/ and used in https://github.com/fortis/go-steam-totp/.

// GenerateTotpCode generates steam TOTP code which is always 5 symbols.
func GenerateTotpCode(sharedSecret string, time time.Time) (string, error) {
	key, err := base64.StdEncoding.DecodeString(sharedSecret)
	if err != nil {
		return "", ErrInvalidSharedSecret
	}

	// Converting time for any reason
	// 00 00 00 00 00 00 00 00
	// 00 00 00 00 xx xx xx xx
	ut := uint64(time.Unix()) / 30
	tb := make([]byte, 8)
	binary.BigEndian.PutUint64(tb, ut)

	// Evaluate hash code for `tb` by key
	mac := hmac.New(sha1.New, key)
	mac.Write(tb)
	hashcode := mac.Sum(nil)

	// Last 4 bits provide initial position
	// len(hashcode) = 20 bytes
	start := hashcode[19] & 0xf

	// Extract 4 bytes at `start` and drop first bit
	fc32 := binary.BigEndian.Uint32(hashcode[start : start+4])
	fc32 &= 1<<31 - 1
	fullcode := int(fc32)

	// Generate auth code
	code := make([]byte, 5)
	for i := range code {
		code[i] = chars[fullcode%charsLen]
		fullcode /= charsLen
	}

	return string(code[:]), nil
}
