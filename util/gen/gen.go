package gen

import (
	"math/rand"

	"github.com/ashashev/go-generics-example/util/call"
	"github.com/ashashev/go-generics-example/util/slice"
)

var alphabet = []rune{
	'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'z', 'x', 'c', 'v', 'b', 'n', 'm',
	'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P', 'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'Z', 'X', 'C', 'V', 'B', 'N', 'M',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'_', '-', '+', '=', '/', '\\', '.', ',', '*', '&', '^', '%', '$', '#', '@', '!', '?', ' ',
}

func RandFromAlphabet[T any](alphabet []T) T {
	return alphabet[rand.Intn(len(alphabet))]
}

func RandAlpha(len int) string {
	return string(slice.Fill(len, call.Semi1(RandFromAlphabet[rune], alphabet[:26])))
}

func RandAlphaNum(len int) string {
	return string(slice.Fill(len, call.Semi1(RandFromAlphabet[rune], alphabet[:36])))
}

func RandStr(len int) string {
	return string(slice.Fill(len, call.Semi1(RandFromAlphabet[rune], alphabet)))
}
