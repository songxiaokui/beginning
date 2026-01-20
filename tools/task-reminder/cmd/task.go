package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"task-reminder/logic"
)

func main() {
	a := app.NewWithID("task.reminder.app")
	w := a.NewWindow("任务提醒")

	w.Resize(fyne.NewSize(900, 600))
	w.SetContent(
		container.NewMax(logic.BuildMainUI(a, w)),
	)

	logic.StartReminderLoop(a)

	w.ShowAndRun()
}
