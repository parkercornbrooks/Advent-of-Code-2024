package main

import (
	"slices"
)

type treeNode struct {
	Before []string
	After  []string
}

func (tn *treeNode) addBefore(n string) {
	tn.Before = append(tn.Before, n)
}

func (tn *treeNode) addAfter(n string) {
	tn.After = append(tn.After, n)
}

type Tree map[string]*treeNode

func (t Tree) addEdge(b, a string) {
	t.addOrUpdate(b, a, true)
	t.addOrUpdate(a, b, false)
}

func (t Tree) addOrUpdate(name, rel string, after bool) {
	node, ok := t[name]
	if !ok {
		node = &treeNode{}
		t[name] = node
	}
	if after {
		node.addAfter(rel)
	} else {
		node.addBefore(rel)
	}
}

func (t Tree) aIsBeforeB(a, b string) bool {
	beforeList := t[a].Before
	return !slices.Contains(beforeList, b)
}

func (t Tree) orderIsValid(order []string) bool {
	for i, val := range order {
		for _, after := range order[i+1:] {
			if t.aIsBeforeB(after, val) {
				return false
			}
		}
	}
	return true
}

func (t Tree) arrange(order []string) []string {
	slices.SortFunc(order, func(a, b string) int {
		if t.aIsBeforeB(a, b) {
			return -1
		}
		return 1
	})
	return order
}
