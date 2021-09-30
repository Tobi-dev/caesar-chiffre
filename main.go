package main

import (
	"flag"
	"fmt"
	"math"
	"strings"
)

var asciiLowercase = []byte("abcdefghijklmnopqrstuvwxyz")
var asciiLowercaseLen = len(asciiLowercase)
var commonSentenceSigns = []byte("!§$%&/()=?*'+#,.-;:_ ")

func main() {
	stringParam := flag.String("string", "", "please provide a string with -s")
	decodeOrEncodeParameter := flag.Bool("decode", false, "Please provide to encode or decode")
	decodeRoundsParameter := flag.Int("rounds", 0, "Please provide the rounds for decoding")
	flag.Parse()

	stringParamLowercase := strings.ToLower(*stringParam)
	stringParamBytes := []byte(stringParamLowercase)

	if *decodeOrEncodeParameter {
		decode(stringParamBytes, *decodeRoundsParameter)
		return
	}
	encode(stringParamBytes, *decodeRoundsParameter)
}

func encode(toEncode []byte, rounds int) {
	encoded := []byte("")

	for _, toEncodeCharacter := range toEncode {
		for _, commonSign := range commonSentenceSigns {
			if toEncodeCharacter == commonSign {
				encoded = append(encoded, commonSign)
			}
		}

		for asciiLowercaseIndex, lowerCaseAlphabet := range asciiLowercase {
			currentIndex := asciiLowercaseIndex
			if toEncodeCharacter == lowerCaseAlphabet {
				if (currentIndex + rounds) > (asciiLowercaseLen - 1) {

					toShift := math.Abs((float64(asciiLowercaseLen) - float64(currentIndex)) - float64(rounds))
					encoded = append(encoded, asciiLowercase[int(toShift)])
					continue
				}
				encoded = append(encoded, asciiLowercase[currentIndex + rounds])
			}
		}
	}
	fmt.Println(string(encoded))
}

func decode(toDecode []byte, rounds int) {
	decoded := []byte("")

	for _, toDecodeCharacter := range toDecode {
		for _, commonSign := range commonSentenceSigns {

			if toDecodeCharacter == commonSign {
				decoded = append(decoded, commonSign)
			}
		}

		for asciiLowercaseIndex, lowerCaseAlphabet := range asciiLowercase {
			currentIndex := asciiLowercaseIndex

			if toDecodeCharacter == lowerCaseAlphabet {
				if (currentIndex - rounds) < 0 {
					toShift := float64(currentIndex) -  math.Abs(float64(rounds))
					decoded = append(decoded, asciiLowercase[asciiLowercaseLen - int(math.Abs(toShift))])
					continue
				}
				decoded = append(decoded, asciiLowercase[asciiLowercaseIndex - rounds])
			}
		}
	}
	fmt.Println(string(decoded))
}