// This package is used to translate emoji characters into their ASCII equivalents and also to translate a number to a string using emojis
package translator

import (
	"math"
	"strconv"
	"strings"
)

// numeric emojis
const EMOJI_DIGIT_0 = "\u0030\ufe0f"
const EMOJI_DIGIT_1 = "\u0031\ufe0f"
const EMOJI_DIGIT_2 = "\u0032\ufe0f"
const EMOJI_DIGIT_3 = "\u0033\ufe0f"
const EMOJI_DIGIT_4 = "\u0034\ufe0f"
const EMOJI_DIGIT_5 = "\u0035\ufe0f"
const EMOJI_DIGIT_6 = "\u0036\ufe0f"
const EMOJI_DIGIT_7 = "\u0037\ufe0f"
const EMOJI_DIGIT_8 = "\u0038\ufe0f"
const EMOJI_DIGIT_9 = "\u0039\ufe0f"
const EMOJI_DIGIT_10 = "\U0001F51F"
const EMOJI_DIGIT_100 = "\U0001F4AF"
const EMOJI_BALL_8 = "\U0001f3b1"

// operator emojis
const EMOJI_PLUS = "\u2795"
const EMOJI_MINUS = "\u2796"
const EMOJI_MULTIPLY = "\u2716"
const EMOJI_DIVIDE = "\u2797"
const EMOJI_NO_RESULT = "\U0001f937"

/**
* translates an string that may content emojis with their ascii equivalents
 */
func EmojiToAscii(expression string) string {
	// multiplication
	expression = strings.ReplaceAll(expression, EMOJI_MULTIPLY, "*")
	// division
	expression = strings.ReplaceAll(expression, EMOJI_DIVIDE, "/")
	// addition
	expression = strings.ReplaceAll(expression, EMOJI_PLUS, "+")
	// substraction
	expression = strings.ReplaceAll(expression, EMOJI_MINUS, "-")
	//digits
	expression = strings.ReplaceAll(expression, EMOJI_DIGIT_1, "1")
	expression = strings.ReplaceAll(expression, EMOJI_DIGIT_2, "2")
	expression = strings.ReplaceAll(expression, EMOJI_DIGIT_3, "3")
	expression = strings.ReplaceAll(expression, EMOJI_DIGIT_4, "4")
	expression = strings.ReplaceAll(expression, EMOJI_DIGIT_5, "5")
	expression = strings.ReplaceAll(expression, EMOJI_DIGIT_6, "6")
	expression = strings.ReplaceAll(expression, EMOJI_DIGIT_7, "7")
	expression = strings.ReplaceAll(expression, EMOJI_DIGIT_8, "8")
	expression = strings.ReplaceAll(expression, EMOJI_DIGIT_9, "9")
	expression = strings.ReplaceAll(expression, EMOJI_DIGIT_10, "10")
	expression = strings.ReplaceAll(expression, EMOJI_DIGIT_100, "100")
	expression = strings.ReplaceAll(expression, EMOJI_BALL_8, "8")

	//sometimes an extra \ufe0f non visible unicode character remains as part of the string so we remove all remaining instances of that
	expression = strings.ReplaceAll(expression, "\ufe0f", "")
	return expression
}

/**
* translates a number into an string with emojis and ascii numbers
 */
func NumberToEmoji(number float64) string {
	if math.IsNaN(number) {
		return EMOJI_NO_RESULT
	}
	result := strconv.Itoa(int(number))
	result = strings.ReplaceAll(result, "100", EMOJI_DIGIT_100)
	result = strings.ReplaceAll(result, "10", EMOJI_DIGIT_10)
	result = strings.ReplaceAll(result, "8", EMOJI_BALL_8)
	result = strings.ReplaceAll(result, "1", EMOJI_DIGIT_1)
	result = strings.ReplaceAll(result, "2", EMOJI_DIGIT_2)
	result = strings.ReplaceAll(result, "3", EMOJI_DIGIT_3)
	result = strings.ReplaceAll(result, "4", EMOJI_DIGIT_4)
	result = strings.ReplaceAll(result, "5", EMOJI_DIGIT_5)
	result = strings.ReplaceAll(result, "6", EMOJI_DIGIT_6)
	result = strings.ReplaceAll(result, "7", EMOJI_DIGIT_7)
	result = strings.ReplaceAll(result, "9", EMOJI_DIGIT_9)
	return result
}
