package nanoid_test

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/aidarkhanov/nanoid"
)

func ExampleFormatString() {
	alphabet, size := "-0123456789ABCDEFGHNRVfgctiUvz_KqYTJkLxpZXIjQW", 21

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
	alphabet := "-0123456789ABCDEFGHNRVfgctiUvz_KqYTJkLxpZXIjQW"
	size := 21

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
