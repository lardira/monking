package mock

import "github.com/lardira/monking/internal/model"

var (
	Jungles = []model.Jungle{
		{
			ID:       1,
			Title:    "Test",
			Monkeys:  10,
			Bananas:  100,
			Coconuts: 5,
		},
	}
)
