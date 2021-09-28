package ui_list

import (
	"fmt"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

type Entry struct {
	name  string
	value string
}

type ListModel struct {
	walk.ReflectListModelBase
	items []Entry
}

func (m *ListModel) Items() interface{} {
	return m.items
}

func (m *ListModel) ItemCount() int {
	return len(m.items)
}

func (m *ListModel) Value(index int) interface{} {
	return m.items[index].name
}

type List struct {
	MainList declarative.ListBox
	lb       *walk.ListBox
}

func NewList() *List {
	var items []Entry
	for i := 20; i > 0; i-- {
		items = append(items, Entry{
			fmt.Sprintf("%d", i),
			fmt.Sprintf("%d%d", i, i),
		})
	}

	list := &List{}
	list.lb = &walk.ListBox{}

	model := &ListModel{items: items}
	// itemStyler := Styler{}

	list.MainList = declarative.ListBox{
		StretchFactor:         1,
		AssignTo:              &list.lb,
		Model:                 model,
		OnCurrentIndexChanged: list.OnCurrentIndexChanged,
		// ItemStyler:            itemStyler,
	}
	return list
}

func (list *List) OnCurrentIndexChanged() {
	fmt.Println("当前点击的", list.lb.CurrentIndex())
}
