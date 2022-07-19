package main

import (
	"container/list"
	"github.com/lxn/walk"
)

const (
	winName = "worktimeTracker"
)

func NewWindowMgr(title string, width int, height int, icon *walk.Icon) (*WinResMgr, *walk.MainWindow) {
	rd := WinResMgr{}
	win, _ := walk.NewMainWindowWithName(winName)
	rd.window = win
	rd.parentList = list.New()

	win.SetTitle(title)
	win.SetLayout(walk.NewVBoxLayout())
	win.SetWidth(width)
	win.SetHeight(height)

	if icon != nil {
		win.SetIcon(icon)
	}
	return &rd, win
}

func NewWindowMgrNoResize(title string, width int, height int, icon *walk.Icon) (*WinResMgr, *walk.MainWindow) {
	rd := WinResMgr{}
	win, _ := walk.NewMainWindowWithName(winName)
	rd.window = win
	rd.parentList = list.New()

	win.SetTitle(title)
	win.SetLayout(walk.NewVBoxLayout())
	win.SetWidth(width)
	win.SetHeight(height)

	if icon != nil {
		win.SetIcon(icon)
	}

	win.SetMinMaxSize(walk.Size{Width: width, Height: height}, walk.Size{Width: width, Height: height})
	rd.NoResize()
	return &rd, win
}
