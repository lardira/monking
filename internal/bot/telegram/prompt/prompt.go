package prompt

import (
	"fmt"
	"strings"

	"github.com/lardira/monking/internal/domain"
)

const (
	promptTextDefault                  string = "Didn't get it, but you can write /help and I'll show you the list of possible commands"
	promptTextHelp                     string = "Possible commands:\n/start \\- shows your jungle info\n/jungle \\- shows your jungle info\n/raid \\- shows the list of jungles that can be raided\n/buy \\- shows the list of buildings you can buy for bananas\n/use \\- shows the list of abilities you can use for coconuts"
	promptTextJungle                   string = "*%v Jungle*\n\t🍌Bananas: %v\n\t🥥Coconuts: %v\n\t🐒Monkeys: %v"
	promptTextRaidList                 string = "You can raid:\n%v"
	promptTextRaidListItem             string = "\\- \\[id:%v\\] %v\n"
	promptTextRaidUnavailableSmallArmy string = "You have too small of an army (need at least 10)"
	promptTextBuy                      string = "Buildings you can buy \\(Costs %v 🍌\\):\n\\- 🐒Monkey Hideout\n\\- 🥥Coconut Tree\n\\- 🍌Banana Palm"
)

const (
	buildingCost = 25
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

func JungleFromModel(model *domain.Jungle) string {
	return Jungle(
		model.Title,
		model.Bananas,
		model.Coconuts,
		model.Monkeys,
	)
}

func RaidList(jungles []domain.Jungle) string {
	var raidList strings.Builder

	for _, jungle := range jungles {
		raidList.WriteString(fmt.Sprintf(promptTextRaidListItem, jungle.ID, jungle.Title))
	}
	return fmt.Sprintf(promptTextRaidList, raidList.String())
}

func RaidUnavailable() string {
	return promptTextRaidUnavailableSmallArmy
}

func Buy() string {
	return fmt.Sprintf(promptTextBuy, buildingCost)
}
