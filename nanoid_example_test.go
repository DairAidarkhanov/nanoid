package nanoid_test

import (
	"crypto/rand"
	"fmt"

	"github.com/aidarkhanov/nanoid"
)

func ExampleFormat() {
	alphabet := nanoid.DefaultAlphabet
	size := nanoid.DefaultSize

	// generateBytesBuffer returns random bytes buffer
	generateBytesBuffer := func(step int) ([]byte, error) {
		buffer := make([]byte, step)
		if _, err := rand.Read(buffer); err != nil {
			return nil, err
		}

		return buffer, nil
	}

	id, err := nanoid.Format(generateBytesBuffer, alphabet, size)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}

func ExampleGenerate() {
	alphabet := nanoid.DefaultAlphabet
	size := nanoid.DefaultSize

	id, err := nanoid.Generate(alphabet, size)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}

func ExampleMustFormat() {
	alphabet := nanoid.DefaultAlphabet
	size := nanoid.DefaultSize

	// generateBytesBuffer returns random bytes buffer
	generateBytesBuffer := func(step int) ([]byte, error) {
		buffer := make([]byte, step)
		if _, err := rand.Read(buffer); err != nil {
			return nil, err
		}

		return buffer, nil
	}

	id := nanoid.MustFormat(generateBytesBuffer, alphabet, size)

	fmt.Println(id)
}

func ExampleMustGenerate() {
	alphabet := nanoid.DefaultAlphabet
	size := nanoid.DefaultSize

	id := nanoid.MustGenerate(alphabet, size)

	fmt.Println(id)
}

func ExampleNew() {
	id := nanoid.New()

	fmt.Println(id)
}
