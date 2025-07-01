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

func IsCommand(input string) bool {
	return strings.HasPrefix(input, "/")
}

func GetMatching(input string) []Command {
	if !IsCommand(input) {
		return [] Command {}
	}
	
	input = input[1:]

	spaceInd := strings.Index(input, " ")
	if spaceInd != -1 {
		input = input[:spaceInd]
	}

	return commandSearcher.GetAvailableValues(input)
}
