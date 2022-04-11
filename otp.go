package otp

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash"
var (
	ErrUnsupportedAlgorithm = errors.New("unsupported algorithm")
	ErrEmptyIssuer          = errors.New("empty issuer")
	ErrCodeLengthMismatch   = errors.New("code length mismatch")
	ErrCodeIsNotValid       = errors.New("code is not valid")
	ErrEncodingNotValid     = errors.New("encoding is not valid")
)

// Algorithm represents the hashing function to use for OTP.
type Algorithm int

const (
	AlgorithmUnknown Algorithm = 0
	AlgorithmSHA1    Algorithm = 1
	AlgorithmSHA256  Algorithm = 2
	AlgorithmSHA512  Algorithm = 3
	algorithmMax     Algorithm = 4
)

func (a Algorithm) String() string {
	switch a {
	case AlgorithmUnknown:
		return ""
	case AlgorithmSHA1:
		return "SHA1"
	case AlgorithmSHA256:
		return "SHA256"
	case AlgorithmSHA512:
		return "SHA512"
	default:
		panic(fmt.Sprintf("otp: unsupported algorithm: %d", int(a)))
	}
}

func (a Algorithm) Hash() hash.Hash {
	switch a {
	case AlgorithmUnknown:
		return nil
	case AlgorithmSHA1:
		return sha1.New()
	case AlgorithmSHA256:
		return sha256.New()
	case AlgorithmSHA512:
		return sha512.New()
	default:
		panic(fmt.Sprintf("otp: unsupported algorithm: %d", int(a)))
	}
}

// Digits is the number of digits in the OTP passcode.
type Digits int

// Six and Eight are the most common values.
const (
	DigitsSix   Digits = 6
	DigitsEight Digits = 8
)

func (d Digits) String() string { return fmt.Sprintf("%d", d) }

// Length of the passcode.
func (d Digits) Length() int { return int(d) }

// Format the number to a digit format (zero-filled upto digits size).
func (d Digits) Format(n int) string {
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", d), n)
}
