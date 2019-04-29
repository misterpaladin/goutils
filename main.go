package goutils

import (
	"math/rand"
	"time"
)

// Random get random number between two integers
func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
