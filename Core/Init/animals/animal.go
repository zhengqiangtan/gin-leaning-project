package animals

import (
	"math/rand"
	"time"
)

var animals = []string{"Cat", "Dog", "Dolphin", "Eagle", "Shark"}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func Random() string {
	i := rand.Intn(len(animals))
	return animals[i]
}
