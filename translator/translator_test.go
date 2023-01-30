package translator

import (
	"math"
	"testing"
)

func TestEmojiToAsciiTranslation(t *testing.T) {

	got := EmojiToAscii("4️2️+🎱+25✖️2")
	want := "42+8+25*2"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got2 := EmojiToAscii("🔟x🔟")
	want2 := "10x10"

	if got2 != want2 {
		t.Errorf("got %q, wanted %q", got2, want2)
	}

}

func TestNumberToEmojiTranslation(t *testing.T) {

	got := NumberToEmoji(100)
	want := EMOJI_DIGIT_100

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got2 := NumberToEmoji(123)
	want2 := EMOJI_DIGIT_1 + EMOJI_DIGIT_2 + EMOJI_DIGIT_3

	if got2 != want2 {
		t.Errorf("got %q, wanted %q", got2, want2)
	}

	got3 := NumberToEmoji(math.NaN())
	want3 := EMOJI_NO_RESULT

	if got3 != want3 {
		t.Errorf("got %q, wanted %q", got3, want3)
	}
}
