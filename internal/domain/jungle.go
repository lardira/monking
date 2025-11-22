package domain

const (
	minRaidSize = 10
)

type Jungle struct {
	ID       int
	Title    string
	Monkeys  int
	Bananas  int
	Coconuts int
}

func (j *Jungle) CanRaid(other *Jungle) bool {
	return j.Monkeys < minRaidSize
}
