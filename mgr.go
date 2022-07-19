package main

import (
	"container/list"
	"github.com/lxn/walk"
	"github.com/lxn/win"
	"time"
)

type WinResMgr struct {
	window     *walk.MainWindow
	parentList *list.List
}

func (m *WinResMgr) StartForeground() {
	m.Center()
	m.Foreground()
	m.window.Show()
	m.window.Run()
}

func (m *WinResMgr) Foreground() {
	win.SetForegroundWindow(m.window.Handle())
}

func (m *WinResMgr) Center() {
	var x, y, width, height int32
	var rtDesk, rtWindow win.RECT
	win.GetWindowRect(win.GetDesktopWindow(), &rtDesk)
	win.GetWindowRect(m.window.Handle(), &rtWindow)

	width = rtWindow.Right - rtWindow.Left
	height = rtWindow.Bottom - rtWindow.Top
	x = (rtDesk.Right - width) / 2
	y = (rtDesk.Bottom - height) / 2

	win.MoveWindow(m.window.Handle(), x, y, width, height, true)
}

func (m *WinResMgr) NoResize() {
	defStyle := win.GetWindowLong(m.window.Handle(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_THICKFRAME
	win.SetWindowLong(m.window.Handle(), win.GWL_STYLE, newStyle)
}

func (m *WinResMgr) GetParent() walk.Container {
	if m.parentList.Len() > 0 {
		parent := m.parentList.Back().Value.(walk.Container)
		return parent
	} else {
		return m.window
	}
}

func (m *WinResMgr) addObj(item walk.Widget) {
	if m.parentList.Len() == 0 {
		m.window.Children().Add(item)
	} else {
		parent := m.parentList.Back().Value.(walk.Container)
		parent.Children().Add(item)
	}
}

func (m *WinResMgr) Hide() {
	m.window.Hide()
}

func (m *WinResMgr) ShowForeground() {
	m.Center()
	m.Foreground()
	m.window.Show()
}

func (m *WinResMgr) HideStart() {
	m.window.Hide()
	m.window.Run()
}

func (m *WinResMgr) HideAndClose() {
	m.window.Synchronize(func() {
		m.window.SetVisible(false)
		m.window.Close()
	})
}

func (m *WinResMgr) DateLabel(date time.Time, format string) *walk.DateLabel {
	ne, _ := walk.NewDateLabel(m.GetParent())
	ne.SetDate(date)
	ne.SetTextAlignment(walk.AlignDefault)
	ne.SetFormat(format)

	m.addObj(ne)
	return ne
}

func (m *WinResMgr) DateLabelCenter(date time.Time, format string) *walk.DateLabel {
	ne, _ := walk.NewDateLabel(m.GetParent())
	ne.SetDate(date)
	ne.SetTextAlignment(walk.AlignCenter)
	ne.SetFormat(format)

	m.addObj(ne)
	return ne
}

func (m *WinResMgr) Label(text string) *walk.Label {
	ne, _ := walk.NewLabel(m.GetParent())
	ne.SetText(text)
	ne.SetTextAlignment(walk.AlignDefault)

	m.addObj(ne)
	return ne
}

func (m *WinResMgr) LabelCenter(text string) *walk.Label {
	ne, _ := walk.NewLabel(m.GetParent())
	ne.SetText(text)
	ne.SetTextAlignment(walk.AlignCenter)

	m.addObj(ne)
	return ne
}

func (m *WinResMgr) LabelRight(text string) *walk.Label {
	ne, _ := walk.NewLabel(m.GetParent())
	ne.SetText(text)
	ne.SetTextAlignment(walk.AlignFar)

	m.addObj(ne)
	return ne
}

func (m *WinResMgr) LabelLeft(text string) *walk.Label {
	ne, _ := walk.NewLabel(m.GetParent())
	ne.SetText(text)
	ne.SetTextAlignment(walk.AlignNear)

	m.addObj(ne)
	return ne
}

func (m *WinResMgr) PushButtonSimple(text string) *walk.PushButton {
	btn, _ := walk.NewPushButton(m.GetParent())
	btn.SetText(text)

	m.addObj(btn)
	return btn
}

func (m *WinResMgr) PushButton(text string, clickFunc func()) *walk.PushButton {
	btn, _ := walk.NewPushButton(m.GetParent())
	btn.SetText(text)
	btn.Clicked().Attach(clickFunc)

	m.addObj(btn)
	return btn
}

func (m *WinResMgr) HSplit() *walk.Splitter {
	hs, _ := walk.NewHSplitter(m.GetParent())
	m.parentList.PushBack(hs)
	return hs
}

func (m *WinResMgr) VSplit() *walk.Splitter {
	vs, _ := walk.NewVSplitter(m.GetParent())
	m.parentList.PushBack(vs)
	return vs
}

func (m *WinResMgr) EndSplit() {
	if m.parentList.Len() > 0 {
		popData := m.parentList.Remove(m.parentList.Back())
		parent := m.GetParent()
		parent.Children().Add(popData.(walk.Widget))
	}
}

func (m *WinResMgr) DisableCloseBox() {
	defStyle := win.GetWindowLong(m.window.Handle(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_SYSMENU
	win.SetWindowLong(m.window.Handle(), win.GWL_STYLE, newStyle)
}

func (m *WinResMgr) Minimize() {
	m.window.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		if m.window.Visible() {
			*canceled = true
			m.window.SetVisible(false)
			TrayMode(m)
		}
	})
}

func (m *WinResMgr) DefClosing() {
	m.window.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		if m.window.Visible() {
			*canceled = true
		}
	})
}

func MsgBox(msg string, window ...*walk.MainWindow) {
	if len(window) > 0 {
		walk.MsgBox(window[0], "알림", msg, walk.MsgBoxOK|walk.MsgBoxSetForeground)
	} else {
		walk.MsgBox(nil, "알림", msg, walk.MsgBoxOK|walk.MsgBoxSetForeground)
	}
}

func Confirm(msg string, window ...*walk.MainWindow) bool {
	if len(window) > 0 {
		if walk.MsgBox(window[0], "알림", msg, walk.MsgBoxYesNo|walk.MsgBoxSetForeground) == win.IDNO {
			return false
		} else {
			return true
		}
	} else {
		if walk.MsgBox(nil, "알림", msg, walk.MsgBoxYesNo|walk.MsgBoxSetForeground) == win.IDNO {
			return false
		} else {
			return true
		}
	}
	return false
}
