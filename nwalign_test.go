package wunsch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanAlign(t *testing.T) {
	assert.Equal(t, int('-'), 45)
	assert.Equal(t, CanAlignString("ABCDEF", "ABCDEF"), true)
	assert.Equal(t, CanAlignString("ABCDEF", "ABCDEF"), true)
	assert.Equal(t, true, CanAlignString("ab--e", "-bcde"))
	assert.Equal(t, true, CanAlignString("abe--", "--ecd"))
	assert.Equal(t, false, CanAlignString("a-be", "ax-e"))
	assert.Equal(t, false, CanAlignString("a---be", "axoo-e"))
	assert.Equal(t, true, CanAlignString("abxe", "ab-e"))
	assert.Equal(t, true, CanAlignString("-a--a--a-------", "0a12a34a12abcde"))
	assert.Equal(t, false, CanAlignString("-ABCDEF", "Q-BCDEF"))
	assert.Equal(t, false, CanAlignString("-ABCDEF-----", "Q-BCDEF"))

}

func TestMultipleAlignment(t *testing.T) {
	data := []string{"Hello", "Hallo", "Hillo", "Hwello", "lloha", "habar"}
	assert.Equal(t, AlignManyString(data),
		[]string{
			"H---ello-----",
			"H--a-llo-----",
			"H-i--llo-----",
			"-----lloha---",
			"--------habar",
			"Hw--ello-----"})

	assert.Equal(t, AlignManyString([]string{"hhhh", "hhhh", "hhhh", "hh22"}),
		[]string{
			"hh--hh",
			"hh--hh",
			"hh--hh",
			"hh22--"})

}

func TestMakeFingerPrint(t *testing.T) {
	data := []string{"Hello", "Hallo", "Hillo", "Hwello", "lloha", "habar"}
	assert.Equal(t, item_to_string(MakeFingerPrint(strings_to_items(data))), "Hwiaellohabar")
}
