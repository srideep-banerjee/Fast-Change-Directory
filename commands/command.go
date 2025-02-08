package commands

import (
	"strings"
)

type Command interface {
    Validate(str string) (reason string)
    Process(str string) error
}

var availbleCommands map[string]Command = map[string]Command {
	"add": add{},
	"remove": remove{},
}

var commandSearcher *StringSearcher[Command] = NewStringSearcherWith(availbleCommands)

func MatchesAny(str string) bool {
	str = str[1:]
	spaceInd := strings.Index(str, " ")
	if spaceInd != -1 {
		str = str[:spaceInd]
    }
	return len(commandSearcher.GetAvailableValues(str)) != 0
}
