package main

import (
	"github.com/leveltype/src/config"
	"github.com/leveltype/src/exercise"
	"github.com/rivo/tview"
)

var (
	GlobalConfig config.Config
)

func main() {

	GlobalConfig.ReadConfiguration()

	app := tview.NewApplication()
	pages := tview.NewPages()

	// pages.AddPage("main",
	// 	tview.NewModal().
	// 		SetText("this the main menu").
	// 		AddButtons([]string{"Ok", "Quit"}).
	// 		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
	// 			if buttonIndex == 0 {
	// 				app.Stop()
	// 			} else {
	// 				pages.SwitchToPage("typing")
	// 			}
	// 		}),
	// 	false,
	// 	true)

	pages.AddPage("typing", exercise.NewExercise(app), true, true)

	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}

}
