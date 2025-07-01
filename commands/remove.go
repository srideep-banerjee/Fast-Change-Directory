package commands

import (
	"me/fast-cd/db"
	"strings"
)

type remove struct {}

func (r remove) Validate(str string) string {
    str = str[strings.IndexRune(str, ' ') + 1:]
    if len(strings.Trim(str, " ")) == 0 {
        return "Provided tag cannot be empty"
    }
    return ""
}

func (r remove) Process(database db.Database, str string) error {
    return nil
}
