package platform

import (
	"github.com/rivo/tview"
)

func CreateTable(data []string) type {
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
