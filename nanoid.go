// Package nanoid provides fast and convenient unique
// string ID generator.
package nanoid

import (
	"crypto/rand"
	"math"
	"strings"
)

const (
	alphabet = "-0123456789ABCDEFGHIJKLNQRTUVWXYZ_cfgijkpqtvxz"
	size     = 21
)

// BytesGenerator represents calculated random bytes
// buffer.
type BytesGenerator func(step int) ([]byte, error)

func random(step int) ([]byte, error) {
	buffer := make([]byte, step)
	if _, err := rand.Read(buffer); err != nil {
		return nil, err
	}
	return buffer, nil
}

// Format generates a random string with the passed in
// bytes generator.
func Format(random BytesGenerator, alphabet string, size int) (string, error) {
	mask := 1
	if len(alphabet) > 1 {
		mask = 2<<uint(math.Log(float64(len(alphabet)-1))/math.Ln2) - 1
	}
	step := int(math.Ceil(1.6 * float64(mask*size) / float64(len(alphabet))))

	var id strings.Builder
	id.Grow(size)
	for {
		randomBytes, err := random(step)
		if err != nil {
			return "", err
		}

		for i := 0; i < step; i++ {
			currentByte := int(randomBytes[i]) & mask

			if currentByte < len(alphabet) {
				if err := id.WriteByte(alphabet[currentByte]); err != nil {
					return "", err
				}
				if id.Len() == size {
					return id.String(), nil
				}
			}
		}
	}
}

// Generate generates a random string.
func Generate(alphabet string, size int) (string, error) {
	id, err := Format(random, alphabet, size)
	if err != nil {
		return "", err
	}
	return id, nil
}

// Must returns a random string if err is nil or panics
// otherwise.
func Must(id string, err error) string {
	if err != nil {
		panic(err)
	}
	return id
}

// MustFormat is like Format but panics if a random string
// cannot be generated.
func MustFormat(random BytesGenerator, alphabet string, size int) string {
	return Must(Format(random, alphabet, size))
}

// MustGenerate is like Generate but panics if a random
// string cannot be generated.
func MustGenerate(alphabet string, size int) string {
	return Must(Generate(alphabet, size))
}

// New generates a random string based on default alphabet
// and size.
func New() string {
	return Must(Generate(alphabet, size))
}
