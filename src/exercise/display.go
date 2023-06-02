package exercise

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/leveltype/src/problemwords"
	"github.com/rivo/tview"
)

var (
	exer          = Exercise{}
	numSelections = 0
	typebox       = &tview.TextView{}
)

func GenerateLesson() string {

	return problemwords.E200BuildFocusLesson()
}

func NewExercise(app *tview.Application) tview.Primitive {

	typebox = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	statsbox := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	statsbox.SetBorder(true)

	typebox.SetInputCapture(func(event *tcell.EventKey) (outevent *tcell.EventKey) {

		// User want to quit?
		if event.Modifiers() == tcell.ModCtrl {
			if event.Key() == tcell.KeyCtrlX {
				app.Stop()
			}
		}

		if event.Key() == tcell.KeyRune {
			exer.WordEntry(event.Rune())
			typebox.SetText(exer.Render())

			// Update stats on word change
			if event.Rune() == ' ' {
				statsbox.SetText(exer.RenderStats())
			}
		}

		return event
	})

	typebox.SetDoneFunc(func(key tcell.Key) {
		currentSelection := typebox.GetHighlights()
		if key == tcell.KeyEnter {
			if len(currentSelection) > 0 {
				typebox.Highlight()
			} else {
				typebox.Highlight("0").ScrollToHighlight()
			}
		} else if len(currentSelection) > 0 {
			index, _ := strconv.Atoi(currentSelection[0])
			if key == tcell.KeyTab {
				index = (index + 1) % numSelections
			} else if key == tcell.KeyBacktab {
				index = (index - 1 + numSelections) % numSelections
			} else {
				return
			}
			typebox.Highlight(strconv.Itoa(index)).ScrollToHighlight()
		}
	})

	flex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(false).SetTitle("left"), 2, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(typebox, 10, 1, true).AddItem(statsbox, 6, 1, false), 75, 1, true).
		AddItem(tview.NewBox().SetBorder(false).SetTitle("right"), 3, 1, false)

	go func() {
		exer.Load(GenerateLesson())
		typebox.SetText(exer.Render())
		statsbox.SetText(exer.RenderStats())
		app.ForceDraw()
	}()
	return flex

}

//Example showing how to make hyperlinks in the text selection.
//go func() {
// for _, word := range strings.Split(exer.Render(), " ") {
// 	if word == "the" {
// 		//word = "[red]the[white]"
// 	}
// 	if word == "to" {
// 		word = fmt.Sprintf(`["%d"]to[""]`, numSelections)
// 		numSelections++
// 	}
// 	fmt.Fprintf(typebox, "%s ", word)
// 	time.Sleep(200 * time.Millisecond)
// }
//typebox.SetText(exer.Render())
//}()
