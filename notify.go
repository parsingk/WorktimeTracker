package main

import (
	"github.com/lxn/walk"
	"log"
)

func Notify() *walk.MainWindow {
	mw, err := walk.NewMainWindow()
	if err != nil {
		log.Panic(err)
	}

	ni, err := walk.NewNotifyIcon(mw)
	if err != nil {
		log.Panic(err)
	}
	defer ni.Dispose()

	if err := ni.SetIcon(GetIcon()); err != nil {
		log.Panic(err)
	}
	if err := ni.SetToolTip("Click for info or use the context menu to exit."); err != nil {
		log.Fatal(err)
	}

	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}

		if err := ni.ShowCustom(
			"퇴근 시간 알림",
			"wortimeTrakcer Notification",
			GetIcon()); err != nil {

			log.Fatal(err)
		}
	})

	exitAction := walk.NewAction()
	if err := exitAction.SetText("Exit"); err != nil {
		log.Fatal(err)
	}
	exitAction.Triggered().Attach(func() { walk.App().Exit(0) })
	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		log.Fatal(err)
	}

	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}

	if err := ni.ShowInfo("퇴근 시간 알림", "worktimeTracker Notification."); err != nil {
		log.Fatal(err)
	}

	return mw
}
