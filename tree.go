// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gxui

// Tree is the interface of all controls that visualize a hierarchical tree
// structure of items.
type Tree interface {
	Focusable

	// SetAdapter binds the specified Adapter to this Tree control, replacing
	// any previously bound Adapter. Items returned by the Adapter that implement
	// the TreeNode interface and have a count greater than 0 will be expandable
	// in the tree.
	SetAdapter(Adapter)

	// Adapter returns the currently bound Adapter.
	Adapter() Adapter

	// Show makes the specified item visible, expanding the tree if necessary.
	Show(AdapterItem)

	// ExpandAll expands all tree nodes.
	ExpandAll()

	// CollapseAll collapses all tree nodes.
	CollapseAll()

	// Selected returns the currently selected item.
	Selected() AdapterItem

	// Select makes the specified item selected. The tree will not automatically
	// expand to the newly selected node.
	Select(AdapterItem)

	// OnSelectionChanged registers the function f to be called when the selection
	// changes.
	OnSelectionChanged(f func(AdapterItem)) EventSubscription
}

// TreeNode is the interface implemented by adapter items that hold sub-items.
type TreeNode interface {
	// Count returns the number of items under this node in the tree.
	Count() int

	// ItemAt returns the AdapterItem for the child item at index i. It is
	// important for the TreeNode to return consistent AdapterItems for the same,
	// data item, so that selections can be persisted, or re-ordering animations
	// can be played.
	// The AdapterItem returned must be equality-unique across the entire Adapter.
	ItemAt(index int) AdapterItem

	// ItemIndex returns the index of the child equal to item, or the index of the
	// child that indirectly contains item, or if the item is not found under this
	// node, -1.
	ItemIndex(item AdapterItem) int

	// Create returns a Control visualizing the item at the specified index.
	Create(theme Theme, index int) Control
}
