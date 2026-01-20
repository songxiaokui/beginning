package logic

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"task-reminder/model"
	"task-reminder/repo"
	"time"
)

func BuildMainUI(app fyne.App, w fyne.Window) fyne.CanvasObject {
	tasks, _ := repo.LoadTasks()

	list := widget.NewList(
		func() int { return len(tasks) },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(i int, o fyne.CanvasObject) {
			t := tasks[i]
			o.(*widget.Label).SetText(
				t.Title + " | 提醒时间: " + t.RemindAt.Format("2006-01-02 15:04"),
			)
		},
	)

	// 定义刷新函数（关键！）
	refreshList := func() {
		newTasks, _ := repo.LoadTasks()
		tasks = newTasks
		list.Refresh()
	}

	btnAdd := widget.NewButton("添加任务", func() {
		openCreateTaskDialog(app, w, refreshList) // 传 refresh
	})

	return container.NewBorder(btnAdd, nil, nil, nil, list)
}

func openCreateTaskDialog(app fyne.App, w fyne.Window, refresh func()) {
	today := time.Now().Format("2006-01-02")

	now := time.Now()
	minute := (now.Minute() / 5) * 5
	defaultTime := fmt.Sprintf("%02d:%02d", now.Hour(), minute)

	title := widget.NewEntry()
	notes := widget.NewMultiLineEntry()
	dateEntry := widget.NewEntry()
	dateEntry.SetText(today)

	timeSelect := widget.NewSelectEntry(generateTimeOptions())
	timeSelect.SetText(defaultTime)

	form := dialog.NewForm(
		"创建任务",
		"保存",
		"取消",
		[]*widget.FormItem{
			widget.NewFormItem("标题", title),
			widget.NewFormItem("备注", notes),
			widget.NewFormItem("日期", dateEntry),
			widget.NewFormItem("时间", timeSelect),
		},
		func(ok bool) {
			if !ok {
				return
			}

			loc := time.Local
			dt := dateEntry.Text + " " + timeSelect.Text
			remindAt, err := time.ParseInLocation("2006-01-02 15:04", dt, loc)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			tasks, _ := repo.LoadTasks()
			tasks = append(tasks, model.Task{
				ID:        time.Now().UnixNano(),
				Title:     title.Text,
				Notes:     notes.Text,
				RemindAt:  remindAt,
				CreatedAt: time.Now(),
			})

			repo.SaveTasks(tasks)

			// 🔥🔥🔥 添加完任务自动刷新
			if refresh != nil {
				refresh()
			}
		},
		w,
	)

	form.Resize(fyne.NewSize(400, 330))
	form.Show()
}

// ====== 生成 96 个时间点（每 15 分钟） ======
func generateTimeOptions() []string {
	var options []string
	for h := 0; h < 24; h++ {
		for m := 0; m < 60; m += 15 {
			options = append(options, fmt.Sprintf("%02d:%02d", h, m))
		}
	}
	return options
}
