package wunsch

import (
	"bytes"
)

type Item interface {
	Id() string
}

type Gap string

func (i Gap) Id() string {
	return string(i)
}

func join_items(result [][]Item, a, b Item) [][]Item {
	return append([][]Item{[]Item{a, b}}, result...)
}

func strings_to_items(a []string) [][]Item {
	result := make([][]Item, 0)
	for _, v := range a {
		result = append(result, string_to_item(v))
	}
	return result
}
func items_to_strings(a [][]Item) []string {
	result := []string{}
	for _, v := range a {
		result = append(result, stringify_items(v))
	}
	return result
}

func string_to_item(a string) []Item {
	ai := make([]Item, 0)
	for _, v := range a {
		ai = append(ai, Gap(v))
	}
	return ai
}

func stringify_items(data []Item) string {
	var b bytes.Buffer
	for _, v := range data {
		b.WriteString(v.Id())
	}
	return b.String()
}

type ItemSlice [][]Item

func (slice ItemSlice) Len() int {
	return len(slice)
}

func (slice ItemSlice) Less(i, j int) bool {
	return len(slice[i]) < len(slice[j])
}

func (slice ItemSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
