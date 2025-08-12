package session

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	adjectives = []string{
		"happy",
		"clever",
		"brave",
		"swift",
		"gentle",
		"bright",
		"calm",
		"eager",
		"bold",
		"wise",
		"kind",
		"quick",
		"proud",
		"fresh",
		"cool",
		"warm",
	}

	animals = []string{
		"fox",
		"bear",
		"wolf",
		"hawk",
		"owl",
		"deer",
		"cat",
		"dog",
		"lion",
		"tiger",
		"dove",
		"crow",
		"fish",
		"frog",
		"bee",
		"ant",
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GeneratePetName returns a random adjective-animal combination
func GeneratePetName() string {
	adjective := adjectives[rand.Intn(len(adjectives))]
	animal := animals[rand.Intn(len(animals))]
	return fmt.Sprintf("%s-%s", adjective, animal)
}