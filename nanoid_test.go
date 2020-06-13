package nanoid_test

import (
	"math"
	"strings"
	"testing"

	"github.com/aidarkhanov/nanoid"
)

func TestGeneratesURLFriendlyIDs(t *testing.T) {
	alphabet := "-0123456789ABCDEFGHIJKLNQRTUVWXYZ_cfgijkpqtvxz"
	for i := 0; i < 100; i++ {
		id := nanoid.New()
		for j := 0; j < len(id); j++ {
			if !strings.Contains(alphabet, string(id[j])) {
				t.Errorf("ID does contain not URL-friendly char: %v", string(id[j]))
			}
		}
	}
}

func TestChangesIDLength(t *testing.T) {
	alphabet := "-0123456789ABCDEFGHIJKLNQRTUVWXYZ_cfgijkpqtvxz"
	id := nanoid.MustGenerate(alphabet, 10)
	if len(id) != 10 {
		t.Errorf("Expected ID length to be %v, got %v", 10, len(id))
	}
}

func TestHasNoCollisions(t *testing.T) {
	count := 100 * 1000
	used := make(map[string]bool)
	for i := 0; i < count; i++ {
		id := nanoid.New()
		if v, ok := used[id]; ok {
			t.Errorf("Repeated ID has been generated: %v", v)
		}
		used[id] = true
	}
}

func TestHasFlatDistribution(t *testing.T) {
	count := 100 * 1000
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	size := 5

	chars := make(map[string]int)
	for i := 0; i < count; i++ {
		id := nanoid.MustGenerate(alphabet, size)
		for j := 0; j < len(id); j++ {
			char := string(id[j])
			if _, ok := chars[char]; !ok {
				chars[char] = 0
			}
			chars[char]++
		}
	}
	if len(chars) != len(alphabet) {
		t.Error("Expected chars length to be equal to alphabet length")
	}

	max := 0.0
	min := float64(math.MaxInt64)
	for _, v := range chars {
		distribution := float64((v * len(alphabet))) / float64((count * size))
		if distribution > max {
			max = distribution
		}
		if distribution < min {
			min = distribution
		}
	}
	distribution := max - min
	if distribution >= 0.05 {
		t.Errorf("Algorithm does not provide enough distribution: %v", distribution)
	}
}

func TestHasOptions(t *testing.T) {
	id := nanoid.MustGenerate("a", 5)
	target := "aaaaa"
	if id != target {
		t.Errorf("Expected %v, got %v", target, id)
	}
}

func TestGeneratesRandomString(t *testing.T) {
	sequence := []byte{2, 255, 3, 7, 7, 7, 7, 7, 0, 1}
	generateBytesBuffer := func(step int) ([]byte, error) {
		buffer := make([]byte, 0, step)
		for i := 0; i < step; i += len(sequence) {
			buffer = append(buffer, sequence[:step-i]...)
		}
		return buffer, nil
	}

	id := nanoid.MustFormat(generateBytesBuffer, "abcde", 4)
	target := "cdac"
	if id != target {
		t.Errorf("Expected %v, got %v", target, id)
	}
}
