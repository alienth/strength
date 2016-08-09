package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"math"
	"math/big"
	"os"
	"unicode"
)

var _ = fmt.Println
var _ = flag.String
var _ = os.Args

// flags:
// -f, --format=       -f words   -f hexstring
// -e, --entropy=      desired bits of entropy
// -w, --words=	       number of words to print
// -b, --whitespace=   use whitespace for words
// -s, --strength      predefined strength class. overrides format, entropy, words.

// should print out roughly how long it would take to bruteforce, for fun

func main() {
	list := filterListByLength(3, 7)
	generateWords(2000, 4, list)
	generateWords(4100, 5, list)
	generateWords(9500, 5, list)
	list = filterListByLength(3, 8)
	generateWords(34000, 5, list)

	generateHexStrings(5, 16)
}

func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateHexStrings(number int, bytes int) {
	fmt.Printf("%d bits of entropy:\n", bytes*8)

	for i := 0; i < number; i++ {
		fmt.Print("    ")
		r, _ := randomBytes(bytes)
		fmt.Println(hex.EncodeToString(r))
	}
}

const punctuation = "!?."

func generateWords(size int, numberOfWords int, list []string) {
	permutations := math.Pow(float64(size), float64(numberOfWords)) * float64(len(punctuation))
	entropy := math.Ilogb(permutations)
	fmt.Printf("%d bits of entropy:\n", entropy)
	for x := 0; x < numberOfWords; x++ {
		fmt.Print("    ")
		for i := 0; i < numberOfWords; i++ {
			foo, _ := rand.Int(rand.Reader, big.NewInt(int64(size)))
			word := capitalize(list[foo.Int64()])
			fmt.Printf("%s", word)
		}
		foo, _ := rand.Int(rand.Reader, big.NewInt(int64(len(punctuation))))
		punc := punctuation[foo.Int64()]
		fmt.Println(string(punc))
	}
	fmt.Println()
}

func capitalize(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	s = string(a)
	return s
}

func filterListByLength(min int, max int) []string {
	var words []string

	for i := 0; i < len(wordList); i++ {
		word := wordList[i]
		if min != 0 && len(word) < min {
			continue
		}
		if max != 0 && len(word) > max {
			continue
		}
		words = append(words, word)
	}

	return words
}
