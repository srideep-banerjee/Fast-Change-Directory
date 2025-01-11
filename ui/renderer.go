package ui

import (
	"github.com/nsf/termbox-go"
)

func RenderInputBox(input string) {
	r, c := termbox.Size()
    line := "+"
    for i := 0; i < r - 2; i++ {
        line += "-"
    }
    line += "+"
    print(0, c - 3, line)
    print(0, c - 2, "|" + input)
    for i := len(input); i < c; i++ {
        print(i, c - 2, " ")
    }
    print(r - 1, c - 2, "|")
    print(0, c - 1, line)
    termbox.Flush()
}

func print(x int, y int, text string) {
    for i, ch := range text {
        termbox.SetCell(x + i, y, ch, termbox.ColorDefault, termbox.ColorDefault)
    }
}

