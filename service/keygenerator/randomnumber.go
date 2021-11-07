package keygenerator

import (
	"math/rand"
	"time"
)

// Let's start with a simple random key generator

// RandomNumber generates new keys with seed as the current time.
// The Generate() returns a new int64 in the range specified by min and max
type RandomNumber struct {
	randomSource    rand.Source
	numberGenerator *rand.Rand

	min, max int64
}

func NewRandomNumber(min, max int64) *RandomNumber {
	r := RandomNumber{}
	r.randomSource = rand.NewSource(time.Now().UnixNano())
	r.numberGenerator = rand.New(r.randomSource)
	r.min = min
	r.max = max

	return &r
}

func (r *RandomNumber) Generate() int64 {
	return r.numberGenerator.Int63n(r.max-r.min) + r.min
}
