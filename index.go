package main

import (
	"os"
	"time"
)

var isWorking bool = false

func main() {
	isRunning := IsRunningProcess()
	if isRunning {
		os.Exit(0)
		return
	}

	LoadIconFromFile("./worktime.ico")

	mgr, _ := NewWindowMgrNoResize("worktimeTracker", 360, 300, GetIcon())

	mgr.Label("출근시간")
	startLabel := mgr.LabelCenter("-")

	mgr.Label("퇴근시간")
	endLabel := mgr.LabelCenter("-")

	btn := mgr.PushButtonSimple("출근")
	btn.Clicked().Attach(func() {
		if isWorking {
			if !Confirm("퇴근하시겠습니까?") {
				return
			}
			startLabel.SetText("-")
			endLabel.SetText("-")
			btn.SetText("출근")
			isWorking = false
		} else {
			if !Confirm("출근하시겠습니까?") {
				return
			}
			now := time.Now()
			workTime, _ := time.ParseDuration("9h")

			startLabel.SetText(now.Format("2006-01-02 15:04:05"))
			endLabel.SetText(now.Add(workTime).Format("2006-01-02 15:04:05"))

			btn.SetText("퇴근")
			isWorking = true

			go func() {
				timer := time.NewTimer(time.Hour * 9)
				<-timer.C
				Notify().Run()
			}()
		}
	})

	mgr.Minimize()
	mgr.StartForeground()
}
