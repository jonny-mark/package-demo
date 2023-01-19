package crypto

import (
	"strconv"
	"testing"
)

func TestBcryptingIsEasy(t *testing.T) {
	expectedHash := []byte("$2a$10$XajjQvNhvvRt5GSeFk1xFeyqRrsxkhBkUiQeg0dt.wU1qD4aFDcga")
	cost, _ := decodeCost(expectedHash)
	println(cost)
}

func decodeCost(sbytes []byte) (int, error) {
	cost, err := strconv.Atoi(string(sbytes[0:2]))
	if err != nil {
		return -1, err
	}

	return cost, nil
}
