package mock

import (
	"math"

	"github.com/lardira/monking/internal/model"
)

var HeavenJungle = model.Jungle{
	ID:       777,
	Title:    "Heaven",
	Monkeys:  math.MaxInt32,
	Bananas:  math.MaxInt32,
	Coconuts: math.MaxInt32,
}

var (
	Jungles = []model.Jungle{
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
