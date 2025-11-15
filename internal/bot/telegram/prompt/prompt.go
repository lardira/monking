package prompt

import (
	"fmt"

	"github.com/lardira/monking/internal/model"
)

const (
	promptTextDefault string = "Didn't get it, but you can write /help and I'll show you the list of possible commands"
	promptTextHelp    string = "Possible commands:\n/start \\- shows your jungle info\n/jungle \\- shows your jungle info\n/raid \\- shows the list of jungles that can be raided\n/buy \\- shows the list of buildings you can buy for bananas\n/use \\- shows the list of abilities you can use for coconuts"
	promptTextJungle  string = "*%v Zoo*\n\t🍌Bananas: %v\n\t🥥Coconuts: %v\n\t🐒Monkeys: %v"
)

func Default() string {
	return promptTextDefault
}

func Help() string {
	return promptTextHelp
}

func Jungle(title string, bananas int, coconuts int, monkeys int) string {
	return fmt.Sprintf(
		promptTextJungle,
		title,
		bananas,
		coconuts,
		monkeys,
	)
}

func JungleFromModel(model *model.Jungle) string {
	return Jungle(
		model.Title,
		model.Bananas,
		model.Coconuts,
		model.Monkeys,
	)
}
