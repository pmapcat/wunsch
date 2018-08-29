// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leachim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2018-29-08 18:56<mklimoff222@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

package wunsch

import (
	"bytes"
)

type Item struct {
	Id    int
	Index int
}

func NewItem(index, id int) Item {
	return Item{Id: id, Index: index}
}
func (i *Item) IsGap() bool {
	return i.Index == -1
}
func NewItemFromRune(index int, point byte) Item {
	if int(point) == 45 {
		return NewGapItem()
	}
	return NewItem(index, int(point))
}
func NewGapItem() Item {
	return Item{Id: -1, Index: -1}
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
		result = append(result, item_to_string(v))
	}
	return result
}

func string_to_item(a string) []Item {
	ai := []Item{}
	for index, id := range a {
		ai = append(ai, NewItem(index, int(id)))
	}
	return ai
}
func IsStringGap(input rune) bool {
	return int(input) == 45
}
func item_to_string(data []Item) string {
	var b bytes.Buffer
	for _, v := range data {
		if v.IsGap() {
			b.WriteString("-")
		} else {
			b.WriteString(string(v.Id))
		}
	}
	return b.String()
}

func min_of(a, b string) string {
	if len(a) < len(b) {
		return a
	}
	return b

}
