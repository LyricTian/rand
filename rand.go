package rand

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
)

// define a flag that generates a random string
const (
	Ldigit = 1 << iota
	LlowerCase
	LupperCase
	LlowerAndUpperCase = LlowerCase | LupperCase
	LdigitAndLowerCase = Ldigit | LlowerCase
	LdigitAndUpperCase = Ldigit | LupperCase
	LdigitAndLetter    = Ldigit | LlowerCase | LupperCase
)

var (
	digits           = []byte("0123456789")
	lowerCaseLetters = []byte("abcdefghijklmnopqrstuvwxyz")
	upperCaseLetters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// definition error
var (
	ErrInvalidLength = errors.New("invalid random length")
	ErrInvalidFlag   = errors.New("invalid random flag")
)

// MustRandom generate a random string specifying the length of the random number
// and the random flag (panic if an error occurs)
func MustRandom(length, flag int) string {
	s, err := Random(length, flag)
	if err != nil {
		panic(err)
	}
	return s
}

// Random generate a random string specifying the length of the random number
// and the random flag
func Random(length, flag int) (string, error) {
	if length == 0 {
		return "", ErrInvalidLength
	}

	var source []byte

	if flag&Ldigit > 0 {
		source = append(source, digits...)
	}

	if flag&LlowerCase > 0 {
		source = append(source, lowerCaseLetters...)
	}

	if flag&LupperCase > 0 {
		source = append(source, upperCaseLetters...)
	}

	sourceLen := len(source)
	if sourceLen == 0 {
		return "", ErrInvalidFlag
	}

	b, err := randomBytesMod(length, byte(sourceLen))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	for _, c := range b {
		buf.WriteByte(source[c])
	}

	return buf.String(), nil
}

// MustUUID generate a random UUID (panic if an error occurs)
func MustUUID() string {
	s, err := UUID()
	if err != nil {
		panic(err)
	}
	return s
}

// UUID generate a random UUID,
// reference: https://github.com/google/uuid/blob/master/version4.go
func UUID() (string, error) {
	buf, err := randomBytes(16)
	if err != nil {
		return "", err
	}
	buf[6] = (buf[6] & 0x0f) | 0x40
	buf[8] = (buf[8] & 0x3f) | 0x80

	dst := make([]byte, 36)
	hex.Encode(dst, buf[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], buf[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], buf[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], buf[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], buf[10:])

	return string(dst), nil
}

func randomBytesMod(length int, mod byte) ([]byte, error) {
	b := make([]byte, length)
	max := 255 - 255%mod
	i := 0

LROOT:
	for {
		r, err := randomBytes(length + length/4)
		if err != nil {
			return nil, err
		}

		for _, c := range r {
			if c >= max {
				// Skip this number to avoid modulo bias
				continue
			}

			b[i] = c % mod
			i++
			if i == length {
				break LROOT
			}
		}

	}

	return b, nil
}

func randomBytes(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
