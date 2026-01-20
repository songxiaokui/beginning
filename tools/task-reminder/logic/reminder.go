package logic

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"task-reminder/model"
	"task-reminder/repo"
	"time"
)

func StartReminderLoop(app fyne.App) {
	go func() {
		for {
			tasks, _ := repo.LoadTasks()
			now := time.Now()

			for _, t := range tasks {
				if !t.IsCompleted && now.After(t.RemindAt) {
					// 为了避免闭包引用问题，拷贝一个局部变量
					taskCopy := t

					// ❗把 UI 操作丢到 Fyne 主线程中执行
					fyne.Do(func() {
						showReminder(app, taskCopy)
					})

					// 标记完成可以留在 goroutine 中，跟 UI 无关
					markTaskDone(t.ID)
				}
			}

			time.Sleep(30 * time.Second)
		}
	}()
}

func showReminder(app fyne.App, t model.Task) {
	win := app.NewWindow("任务提醒")

	// ❗不要再使用 SetMaster()
	// win.SetMaster() // ❌ 这个会导致关闭弹窗时整个 App 退出

	win.SetFixedSize(true)
	win.Resize(fyne.NewSize(400, 240))
	win.CenterOnScreen()

	// 🔥 强制窗口置顶显示（Fyne 官方推荐）
	win.RequestFocus()

	label := widget.NewLabel("🔔 任务到点啦：\n" + t.Title + "\n\n" + t.Notes)
	label.Wrapping = fyne.TextWrapWord

	btn := widget.NewButton("我知道了", func() {
		win.Close() // 只关闭弹窗，不退出 App
	})

	win.SetContent(container.NewVBox(
		label,
		btn,
	))

	win.Show()

	// ⛑️ 确保点击 X 也不会退出 app
	win.SetCloseIntercept(func() {
		win.Close() // 只是关闭弹窗
	})
}

func markTaskDone(id int64) {
	tasks, _ := repo.LoadTasks()
	updated := []model.Task{}
	for _, t := range tasks {
		if t.ID == id {
			t.IsCompleted = true
		}
		updated = append(updated, t)
	}
	_ = repo.SaveTasks(updated)
}
