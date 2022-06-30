package _7_binary_tree

import (
	list "algo/level_02/02_list"
	"strconv"
)

// PreSerialize 先序序列化
func PreSerialize(t *Tree) []string {
	s := make([]string, 0, 10)
	s = preSerialize(t, s)
	return s
}

func preSerialize(n list.Node, s []string) []string {
	if n == nil {
		s = append(s, "")
	} else {
		s = append(s, strconv.Itoa(int(n.Value())))
		s = preSerialize(n.Left(), s)
		s = preSerialize(n.Right(), s)
	}

	return s
}
