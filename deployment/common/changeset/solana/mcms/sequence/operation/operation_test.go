package operation

import (
	"errors"
	"strings"
	"testing"
)

var errEntropyUnavailable = errors.New("entropy unavailable")

func TestRandomSeed(t *testing.T) {
	seed, err := randomSeed(strings.NewReader(strings.Repeat("\x00", 32)))
	if err != nil {
		t.Fatalf("randomSeed returned an error: %v", err)
	}

	for _, character := range seed {
		if !strings.ContainsRune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", rune(character)) {
			t.Errorf("randomSeed returned invalid character %q", character)
		}
	}
}

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) {
	return 0, errEntropyUnavailable
}

func TestRandomSeedReturnsEntropyError(t *testing.T) {
	_, err := randomSeed(failingReader{})
	if !errors.Is(err, errEntropyUnavailable) {
		t.Fatalf("randomSeed error = %v, want %v", err, errEntropyUnavailable)
	}
}
