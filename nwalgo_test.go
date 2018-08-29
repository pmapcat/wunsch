// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leachim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2018-29-08 18:19<mklimoff222@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

package wunsch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlignment(t *testing.T) {
	a := "GAATTCAGTTA"
	b := "GGATCGA"
	as, bs, score := NeedlemanWunschString(1, 0, 0, a, b)
	assert.Equal(t, 6, score)
	assert.Equal(t, "G-AATTCAGTTA", as)
	assert.Equal(t, "GGA-T-C-G--A", bs)

	as, bs, _ = NeedlemanWunschString(1, 0, 0, "ABCDEF", "QBCDEF")
	assert.Equal(t, "-ABCDEF", as)
	assert.Equal(t, "Q-BCDEF", bs)

	as, bs, _ = NeedlemanWunschString(1, 0, 0, "bbppb", "bbbppb")
	assert.Equal(t, "bb-ppb", as)
	assert.Equal(t, "bbbppb", bs)

	as, bs, _ = NeedlemanWunschString(1, 0, 0, "trtrtdaatd", "trtdaaatd")
	assert.Equal(t, "trtrtdaa-td", as)
	assert.Equal(t, "trt--daaatd", bs)

	as, bs, _ = NeedlemanWunschString(1, 0, 0, "abe", "bcde")
	assert.Equal(t, "ab--e", as)
	assert.Equal(t, "-bcde", bs)

	as, bs, _ = NeedlemanWunschString(1, 0, 0, "abe", "ecd")
	assert.Equal(t, "abe--", as)
	assert.Equal(t, "--ecd", bs)

	as, bs, _ = NeedlemanWunschString(1, 0, 0, "abe", "axe")
	assert.Equal(t, "a-be", as)
	assert.Equal(t, "ax-e", bs)

	as, bs, _ = NeedlemanWunschString(1, 0, 0, "abe", "axooe")
	assert.Equal(t, "a---be", as)
	assert.Equal(t, "axoo-e", bs)

	as, bs, _ = NeedlemanWunschString(1, 0, 0, "abxe", "abe")
	assert.Equal(t, "abxe", as)
	assert.Equal(t, "ab-e", bs)

	as, bs, _ = NeedlemanWunschString(1, 0, 0, "aaa", "0a12a34a12abcde")
	assert.Equal(t, "-a--a--a-------", as)
	assert.Equal(t, "0a12a34a12abcde", bs)

}
