package nanoid_test

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/aidarkhanov/nanoid/v2"
)

func ExampleFormatString() {
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

	id, err := nanoid.FormatString(generateBytesBuffer, alphabet, size)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(id)
}

func ExampleGenerateString() {
	alphabet := nanoid.DefaultAlphabet
	size := nanoid.DefaultSize

	id, err := nanoid.GenerateString(alphabet, size)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(id)
}

func ExampleNew() {
	id, err := nanoid.New()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(id)
}
