package main

import (
	"strings"

	"eleven.cn/eleven_win/ui/ui_list"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit
	//菜单栏
	mainWD := declarative.MainWindow{
		Title:     "golang桌面应用程序",
		Layout:    declarative.VBox{},
		Icon:      "./assets/icon.png",
		MinSize:   declarative.Size{Width: 900, Height: 300},
		Size:      declarative.Size{Width: 900, Height: 300},
		MaxSize:   declarative.Size{Width: 900, Height: 300},
		MenuItems: createMenuItems(),
		ToolBar:   createToolBars(),
		Children: []declarative.Widget{
			declarative.HSplitter{
				Children: []declarative.Widget{
					ui_list.NewList().MainList,
					declarative.TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			declarative.VSplitter{ //排列方式 VSplitter纵向 HSplitter横向
				Children: []declarative.Widget{
					declarative.TextEdit{AssignTo: &outTE, ReadOnly: true},
					declarative.TextEdit{AssignTo: &inTE, MaxLength: 10},
				},
			},
			declarative.PushButton{
				Text: "点击按钮",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}
	mainWD.Run()
}

func createMenuItems() []declarative.MenuItem {
	allMenuItems := []declarative.MenuItem{
		declarative.Menu{
			Text: "文件",
			Items: []declarative.MenuItem{
				declarative.Action{
					Text: "打开文件",
					Shortcut: declarative.Shortcut{ //定义快捷键后会有响应提示显示
						Modifiers: walk.ModControl,
						Key:       walk.KeyO,
					},
					OnTriggered: func() {}, //点击动作触发响应函数
				},
				declarative.Action{
					Text: "另存为",
					Shortcut: declarative.Shortcut{
						Modifiers: walk.ModControl | walk.ModShift,
						Key:       walk.KeyS,
					},
					OnTriggered: func() {

					},
				},
				declarative.Action{
					Text: "退出",
					OnTriggered: func() {
						//mw.Close()
					},
				},
			},
		},
		declarative.Menu{
			Text: "帮助",
			Items: []declarative.MenuItem{
				declarative.Action{
					Text: "关于",
					OnTriggered: func() {
						// walk.MsgBox(mw, "关于", "这是一个菜单和工具栏的实例",
						// 	walk.MsgBoxIconInformation|walk.MsgBoxDefButton1)
					},
				},
			},
		},
	}
	return allMenuItems
}

func createToolBars() declarative.ToolBar {
	toolBar := declarative.ToolBar{ //工具栏
		ButtonStyle: declarative.ToolBarButtonTextOnly,
		Items: []declarative.MenuItem{
			declarative.Menu{
				Text: "工具箱1",
				Items: []declarative.MenuItem{
					declarative.Action{
						Text:        "A",
						OnTriggered: func() {},
					},
					declarative.Action{
						Text:        "B",
						OnTriggered: func() {},
					},
				},
				OnTriggered: func() {}, //在菜单中不可如此定义，会无响应
			},
			declarative.Separator{}, //分隔符
			declarative.Menu{
				Text: "工具箱2",
				Items: []declarative.MenuItem{
					declarative.Action{
						Text:        "A",
						OnTriggered: func() {},
					},
					declarative.Action{
						Text:        "B",
						OnTriggered: func() {},
					},
				},
				OnTriggered: func() {}, //在菜单中不可如此定义，会无响应
			},
		},
	}
	return toolBar
}
