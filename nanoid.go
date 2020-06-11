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
	// This alphabet uses `A-Za-z0-9_-` symbols.
	// The genetic algorithm helped optimize the gzip
	// compression for this alphabet.
	DefaultAlphabet = "ModuleSymbhasOwnPr-0123456789ABCDEFGHNRVfgctiUvz_KqYTJkLxpZXIjQW"

	// Default number of symbols in identifier
	DefaultSize = 21
)

// BytesGenerator represents random bytes buffer.
type BytesGenerator func(step int) (buffer []byte, err error)

func generateRandomBuffer(step int) (buffer []byte, err error) {
	buffer = make([]byte, step)
	if _, err = rand.Read(buffer); err != nil {
		return nil, err
	}

	return buffer, nil
}

// FormatString generates a random string based on
// BytesGenerator, alphabet and size.
// nolint: gomnd
func FormatString(generateRandomBuffer BytesGenerator, alphabet string, size int) (string, error) {
	var (
		mask = 2<<uint32(31-bits.LeadingZeros32(uint32(len(alphabet)-1|1))) - 1
		step = int(math.Ceil(1.6 * float64(mask*size) / float64(len(alphabet))))

		id = new(strings.Builder)
	)

	id.Grow(size)

	for {
		randomBuffer, err := generateRandomBuffer(step)
		if err != nil {
			return "", err
		}

		for i := step - 1; i >= 0; i-- {
			currentIndex := int(randomBuffer[i]) & mask

			if currentIndex >= len(alphabet) {
				continue
			}

			if err := id.WriteByte(alphabet[currentIndex]); err != nil {
				return "", err
			}

			if id.Len() == size {
				return id.String(), nil
			}
		}
	}
}

// GenerateString generates a random string based on
// alphabet and size.
func GenerateString(alphabet string, size int) (id string, err error) {
	return FormatString(generateRandomBuffer, alphabet, size)
}

// New generates a random string.
func New() (id string, err error) {
	return GenerateString(DefaultAlphabet, DefaultSize)
}
