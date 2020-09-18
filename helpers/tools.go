package helpers

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

// ParseInt64 helper to avoid code repetition
func ParseUInt64(stringToParse string) (uint64, error) {
	intID, err := strconv.ParseInt(stringToParse, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse string to int")
	}
	return uint64(intID), nil
}

func RandomString10() (string){
	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)

	return s
}