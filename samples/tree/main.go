// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"
)

// item is the adapter item type for the tree adapter. item conforms to the
// TreeNode interface, making the items with children expandable.
type item struct {
	name     string
	children []*item
}

// String returns the item name. It is used by the DefaultAdapter to visualize
// the item as a Label.
func (n item) String() string {
	return n.name
}

func (n item) Count() int {
	return len(n.children)
}

func (n item) ItemAt(index int) gxui.AdapterItem {
	return n.children[index]
}

func (n item) ItemIndex(item gxui.AdapterItem) int {
	for i, c := range n.children {
		if c == item || c.ItemIndex(item) >= 0 {
			return i
		}
	}
	return -1
}

func (n item) Create(theme gxui.Theme, index int) gxui.Control {
	l := theme.CreateLabel()
	l.SetText(n.children[index].name)
	return l
}

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	// helper function for building nodes using variadics for the children.
	newItem := func(name string, children ...*item) *item {
		return &item{name, children}
	}

	layout := theme.CreateLinearLayout()
	layout.SetOrientation(gxui.Vertical)

	mammals := newItem("Mammals",
		newItem("Cats"),
		newItem("Dogs"),
		newItem("Horses"),
		newItem("Duck-billed platypuses"),
	)

	birds := newItem("Birds",
		newItem("Peacocks"),
		newItem("Doves"),
	)

	reptiles := newItem("Reptiles",
		newItem("Lizards"),
		newItem("Turtles"),
		newItem("Crocodiles"),
		newItem("Snakes"),
	)

	amphibians := newItem("Amphibians",
		newItem("Frogs"),
		newItem("Toads"),
	)

	arthropods := newItem("Arthropods",
		newItem("Crustaceans",
			newItem("Crabs"),
			newItem("Lobsters"),
		),
		newItem("Insects",
			newItem("Ants"),
			newItem("Bees"),
		),
		newItem("Arachnids",
			newItem("Spiders"),
			newItem("Scorpions"),
		),
	)

	adapter := gxui.CreateDefaultAdapter()
	adapter.SetSize(math.Size{W: math.MaxSize.W, H: 18})
	adapter.SetItems(newItem("Animals", mammals, birds, reptiles, amphibians, arthropods))

	tree := theme.CreateTree()
	tree.SetAdapter(adapter)
	tree.Select(amphibians)
	tree.Show(tree.Selected())

	layout.AddChild(tree)

	row := theme.CreateLinearLayout()
	row.SetOrientation(gxui.Horizontal)
	layout.AddChild(row)

	expandAll := theme.CreateButton()
	expandAll.SetText("Expand All")
	expandAll.OnClick(func(gxui.MouseEvent) { tree.ExpandAll() })
	row.AddChild(expandAll)

	collapseAll := theme.CreateButton()
	collapseAll.SetText("Collapse All")
	collapseAll.OnClick(func(gxui.MouseEvent) { tree.CollapseAll() })
	row.AddChild(collapseAll)

	window := theme.CreateWindow(800, 600, "Tree view")
	window.AddChild(layout)
	window.OnClose(driver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
	gxui.EventLoop(driver)
}

func main() {
	gl.StartDriver(appMain)
}
