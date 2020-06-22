// Package nanoid provides fast and convenient unique string generator.
package nanoid

import (
	"crypto/rand"
	"math"
	"math/bits"
	"strings"
)

const (
	// DefaultAlphabet is the default alphabet for Nano ID.
	DefaultAlphabet = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
	// DefaultSize is the default size for Nano ID.
	DefaultSize = 21
)

// BytesGenerator represents random bytes buffer.
type BytesGenerator func(step int) ([]byte, error)

func generateRandomBuffer(step int) ([]byte, error) {
	buffer := make([]byte, step)
	if _, err := rand.Read(buffer); err != nil {
		return nil, err
	}
	return buffer, nil
}

// Format generates a random string based on BytesGenerator, alphabet and size.
func Format(generateRandomBuffer BytesGenerator, alphabet string, size int) (string, error) {
	mask := 2<<uint32(31-bits.LeadingZeros32(uint32(len(alphabet)-1|1))) - 1
	step := int(math.Ceil(1.6 * float64(mask*size) / float64(len(alphabet))))

	id := new(strings.Builder)
	id.Grow(size)

	for {
		randomBuffer, err := generateRandomBuffer(step)
		if err != nil {
			return "", err
		}

		for i := 0; i < step; i++ {
			currentIndex := int(randomBuffer[i]) & mask

			if currentIndex < len(alphabet) {
				if err := id.WriteByte(alphabet[currentIndex]); err != nil {
					return "", err
				} else if id.Len() == size {
					return id.String(), nil
				}
			}
		}
	}
}

// Generate generates a random string based on alphabet and size.
func Generate(alphabet string, size int) (string, error) {
	id, err := Format(generateRandomBuffer, alphabet, size)
	if err != nil {
		return "", err
	}
	return id, nil
}

// Must returns a random string if err is nil or panics otherwise.
func Must(id string, err error) string {
	if err != nil {
		panic(err)
	}
	return id
}

// MustFormat is like Format but panics if a random string cannot be generated.
func MustFormat(generateRandomBuffer BytesGenerator, alphabet string, size int) string {
	return Must(Format(generateRandomBuffer, alphabet, size))
}

// MustGenerate is like Generate but panics if a random string cannot be generated.
func MustGenerate(alphabet string, size int) string {
	return Must(Generate(DefaultAlphabet, DefaultSize))
}

// New generates a random string.
func New() string {
	return Must(Generate(DefaultAlphabet, DefaultSize))
}
