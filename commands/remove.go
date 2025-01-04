package commands

import (
    "strings"
)

type remove struct {}

func (r remove) Matches(str string) bool {
    spaceInd := strings.IndexRune(str, ' ')

    if spaceInd != -1 {
        str = str[:spaceInd]
    }

    return strings.HasPrefix("remove", str)
}

func (r remove) Validate(str string) string {
    str = str[strings.IndexRune(str, ' ') + 1:]
    if len(strings.Trim(str, " ")) == 0 {
        return "Provided tag cannot be empty"
    }
    return ""
}

func (r remove) Process(str string) error {
    return nil
}
