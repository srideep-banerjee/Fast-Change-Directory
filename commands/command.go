package commands

import (
	"me/fast-cd/db"
	"strings"
)

type Command interface {
	Validate(str string) (reason string)
	Process(database db.Database, str string) error
}

var availbleCommands = []StringSearcherEntry[Command] {
	{key: "add", value: add{}},
	{key: "remove", value: remove{}},
}

var commandSearcher *StringSearcher[Command] = NewStringSearcherWith(availbleCommands)

func GetMatching(prefix string) []Command {
	prefix = prefix[1:]

	spaceInd := strings.Index(prefix, " ")
	if spaceInd != -1 {
		prefix = prefix[:spaceInd]
	}

	return commandSearcher.GetAvailableValues(prefix)
}
