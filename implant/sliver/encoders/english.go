package encoders

/*
	Sliver Implant Framework
	Copyright (C) 2019  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"crypto/rand"
	"math/big"
	"strings"
)

var dictionary map[int][]string

var rawEnglishDictionary []string

func getEnglishDictionary() []string {
	return rawEnglishDictionary
}

// English Encoder - An ASCIIEncoder for binary to english text
type EnglishEncoder struct{}

// Encode - Binary => English
func (e EnglishEncoder) Encode(data []byte) ([]byte, error) {
	if dictionary == nil {
		buildDictionary()
	}
	words := []string{}
	for _, b := range data {
		possibleWords := dictionary[int(b)]
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(possibleWords))))
		if err != nil {
			return nil, err
		}
		words = append(words, possibleWords[idx.Int64()])
	}
	return []byte(strings.Join(words, " ")), nil
}

// Decode - English => Binary
func (e EnglishEncoder) Decode(words []byte) ([]byte, error) {
	wordList := strings.Split(string(words), " ")
	data := []byte{}
	for _, word := range wordList {
		word = strings.TrimSpace(word)
		if len(word) == 0 {
			continue
		}
		byteValue := sumWord(word)
		data = append(data, byte(byteValue))
	}
	return data, nil
}

func buildDictionary() {
	dictionary = map[int][]string{}
	for _, word := range getEnglishDictionary() {
		word = strings.TrimSpace(word)
		sum := sumWord(word)
		dictionary[sum] = append(dictionary[sum], word)
	}
}

func sumWord(word string) int {
	sum := 0
	for _, char := range word {
		sum += int(char)
	}
	return sum % 256
}
