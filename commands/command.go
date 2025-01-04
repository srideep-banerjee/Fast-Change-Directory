package commands

type Command interface {
    Matches(str string) bool
    Validate(str string) (reason string)
    Process(str string) error
}

var commandList = []Command {
    add {},
    remove {},
}

func MatchesAny(str string) bool {
    for _, cmd := range commandList {
        if cmd.Matches(str) {
            return true
        }
    }
    return false
}
