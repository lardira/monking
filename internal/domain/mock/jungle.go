package mock

import (
	"math"

	"github.com/lardira/monking/internal/domain"
)

var HeavenJungle = domain.Jungle{
	ID:       777,
	Title:    "Heaven",
	Monkeys:  math.MaxInt32,
	Bananas:  math.MaxInt32,
	Coconuts: math.MaxInt32,
}

var (
	Jungles = []domain.Jungle{
		{
			ID:       1,
			Title:    "Test",
			Monkeys:  10,
			Bananas:  100,
			Coconuts: 5,
		},
		HeavenJungle,
	}
)
