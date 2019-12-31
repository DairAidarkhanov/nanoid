// Package nanoid provides fast and convenient unique
// string generator.
package nanoid

import (
	"crypto/rand"
	"math"
	"math/bits"
	"strings"
)

const (
	alphabet = "-0123456789ABCDEFGHNRVfgctiUvz_KqYTJkLxpZXIjQW"
	size     = 21
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

// FormatString generates a random string based on
// BytesGenerator, alphabet and size.
func FormatString(generateRandomBuffer BytesGenerator, alphabet string, size int) (string, error) {
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

// GenerateString generates a random string based on
// alphabet and size.
func GenerateString(alphabet string, size int) (string, error) {
	id, err := FormatString(generateRandomBuffer, alphabet, size)
	if err != nil {
		return "", err
	}
	return id, nil
}

// New generates a random string.
func New() (string, error) {
	id, err := FormatString(generateRandomBuffer, alphabet, size)
	if err != nil {
		return "", err
	}
	return id, nil
}
