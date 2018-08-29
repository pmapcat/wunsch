// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leachim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2018-29-08 18:19<mklimoff222@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

package wunsch

import (
	"sort"
)

func CanAlign(data [][]Item) bool {
	// do not touch
	hasLetterA := false
	hasLetterB := false
	for _, v := range data {
		right := v[0]
		left := v[1]
		if right.Id() == left.Id() {
			// start item
			if hasLetterA == true && hasLetterB == true {
				return false
			}
			hasLetterA = false
			hasLetterB = false
			continue
		}
		if right.Id() != "-" {
			hasLetterA = true
		}
		if left.Id() != "-" {
			hasLetterB = true
		}
	}
	return !(hasLetterA == true && hasLetterB == true)
}

func SimpleAlignment(alignment [][]Item) []Item {
	// align two items
	result := []Item{}
	for _, v := range alignment {
		right := v[0]
		left := v[1]
		if right.Id() == "-" {
			result = append(result, left)
			continue
		}
		if left.Id() == "-" {
			result = append(result, right)
			continue
		}
		result = append(result, right)
	}
	return result
}

func MakeFingerPrint(data [][]Item) []Item {
	sort.Sort(ItemSlice(data))
	fingerPrint := data[0]
	for _, v := range data {
		data, _ := NeedlemanWunsch(fingerPrint, v, 1, 0, 0)
		fingerPrint = SimpleAlignment(data)
	}
	return fingerPrint
}

func AlignMany(data [][]Item) ([][]Item, []Item) {
	// Returns aligned data and it's fingerprint
	// Select largest item.
	populace_data := data
	fingerPrint := MakeFingerPrint(data)

	result := make([][]Item, 0)
	for _, v := range populace_data {
		data, _ := NeedlemanWunsch(fingerPrint, v, 1, 0, 0)
		shouldAppend := make([]Item, 0)
		for _, v := range data {
			shouldAppend = append(shouldAppend, v[1])
		}
		result = append(result, shouldAppend)
	}
	return result, fingerPrint
}

func AlignManyString(data []string) []string {
	datai := make([][]Item, 0)
	for _, v := range data {
		datai = append(datai, string_to_item(v))
	}
	result, _ := AlignMany(datai)
	must_return := make([]string, 0)
	for _, v := range result {
		must_return = append(must_return, stringify_items(v))
	}
	return must_return
}
func CanAlignString(a, b string) bool {
	result := make([][]Item, 0)
	for k, _ := range a {
		result = append(result, []Item{Gap(a[k]), Gap(b[k])})
	}
	return CanAlign(result)
}
