package commands

import (
	"me/fast-cd/db"
	"me/fast-cd/util"
	"strings"
)

type Command interface {
	Validate(str string) (reason string)
	Process(database db.Database, str string) error
}

var availbleCommands = [] util.StringSearcherEntry[Command] {
	util.NewStringSearcherEntry[Command]("add", add{}),
	util.NewStringSearcherEntry[Command]("remove", remove{}),
}

var commandSearcher *util.StringSearcher[Command] = util.NewStringSearcherWith(availbleCommands)

func GetMatching(prefix string) []Command {
	prefix = prefix[1:]

	spaceInd := strings.Index(prefix, " ")
	if spaceInd != -1 {
		prefix = prefix[:spaceInd]
	}

	return commandSearcher.GetAvailableValues(prefix)
}
