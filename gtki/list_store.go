package gtki

import "github.com/coyim/gotk3adapter/glibi"

type ListStore interface {
	glibi.Object
	TreeModel

	Append() TreeIter
	Clear()
	Remove(TreeIter) bool
	Set2(TreeIter, []int, []interface{}) error
	SetValue(TreeIter, int, interface{}) error
	GetColumnType(index int) glibi.Type
}

func AssertListStore(_ ListStore) {}
