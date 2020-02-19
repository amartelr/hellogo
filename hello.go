package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	MainWindow{
		Title:  "Hola",
		Layout: VBox{},
		Children: []Widget{
			Label{Text: "Hola Mundo!"},
			PushButton{
				Text: "Salir",
				OnClicked: func() {
					walk.App().Exit(0)
				},
			},
		},
	}.Run()
}
