// This is the modification of the implementation of the algorithm by Andrew E. Bruno.  @
// You can find original work here: https://github.com/aebruno/nwalgo/                  @
// Below is the retained copyright notice                                               @
// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// Copyright 2015 Andrew E. Bruno. All rights reserved.                                 @
// Use of this source code is governed by a MIT style                                   @
// license that can be found in the LICENSE file.                                       @
// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leachim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2018-29-08 18:16 <mklimoff222@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

package wunsch

const (
	UP   = 1
	LEFT = 2
	NW   = 3
	NONE = 0
)

func NeedlemanWunsch(a, b []Item, match, mismatch, gap int) ([][]Item, int) {

	alen := len(a) + 1
	blen := len(b) + 1

	f := make([][]int, alen)
	pointer := make([][]int, alen)
	for i := range f {
		f[i] = make([]int, blen)
		pointer[i] = make([]int, blen)
	}

	for i := 1; i < alen; i++ {
		f[i][0] = gap * i
		pointer[i][0] = UP
	}
	for j := 1; j < blen; j++ {
		f[0][j] = gap * j
		pointer[0][j] = LEFT
	}

	pointer[0][0] = NONE
	for i := 1; i < alen; i++ {
		for j := 1; j < blen; j++ {
			match_mismatch := mismatch
			if a[i-1].Id == b[j-1].Id {
				match_mismatch = match
			}

			max := f[i-1][j-1] + match_mismatch
			hgap := f[i-1][j] + gap
			vgap := f[i][j-1] + gap
			if hgap > max {
				max = hgap
			}
			if vgap > max {
				max = vgap
			}

			p := NW
			if max == hgap {
				p = UP
			} else if max == vgap {
				p = LEFT
			}

			pointer[i][j] = p
			f[i][j] = max
		}
	}

	i := alen - 1
	j := blen - 1
	score := f[i][j]

	result := make([][]Item, 0)
	for p := pointer[i][j]; p != NONE; p = pointer[i][j] {
		if p == NW {
			result = join_items(result, a[i-1], b[j-1])
			i--
			j--
		} else if p == UP {
			result = join_items(result, a[i-1], NewGapItem())
			i--
		} else if p == LEFT {
			result = join_items(result, NewGapItem(), b[j-1])
			j--
		}
	}
	return result, score
}

func UnzipListOfItems(items [][]Item) ([]Item, []Item) {
	a := []Item{}
	b := []Item{}
	for _, item := range items {
		a = append(a, item[0])
		b = append(b, item[1])
	}
	return a, b
}

func NeedlemanWunschString(match, mismatch, gap int, a, b string) (string, string, int) {
	ai := []Item{}
	bi := []Item{}
	ai = string_to_item(a)
	bi = string_to_item(b)
	result, score := NeedlemanWunsch(ai, bi, match, mismatch, gap)
	na, nb := UnzipListOfItems(result)
	return item_to_string(na), item_to_string(nb), score
}
