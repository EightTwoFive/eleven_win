package ui_list

import "github.com/lxn/walk"

type Styler struct {
}

func (s Styler) DefaultItemHeight() int {
	return 100
}

func (s Styler) ItemHeight(index int, width int) int {
	return 200
}

func (s Styler) ItemHeightDependsOnWidth() bool {
	return true
}

func (s Styler) StyleItem(style *walk.ListItemStyle) {

}
