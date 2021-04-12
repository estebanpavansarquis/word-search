package utils

import (
	"reflect"
)

type WordSearchInput struct {
	AlphabetSoup [][]rune
	Words        []string
}

func GetWordSearchInput() WordSearchInput {
	return WordSearchInput{
		[][]rune{
			{'P', 'S', 'S', 'A', 'B', 'I', 'A'},
			{'P', 'B', 'S', 'C', 'C', 'U', 'A'},
			{'P', 'A', 'L', 'A', 'B', 'R', 'A'},
			{'Q', 'S', 'S', 'P', 'B', 'Ñ', 'K'},
			{'P', 'B', 'S', 'C', 'C', 'Y', 'I'},
			{'Q', 'S', 'S', 'S', 'B', 'Ñ', 'K'},
			{'P', 'D', 'A', 'A', 'I', 'Y', 'I'},
			{'O', 'D', 'N', 'T', 'T', 'Ñ', 'K'},
			{'V', 'U', 'E', 'L', 'A', 'Y', 'I'},
		},
		[]string{"PALABRA", "ABRA", "X","BALA","Q","SALA","AAA","NOESTA","DANDOVUELTAS"},
	}
}

var alphabetSoup [][]rune

func IsPresent(soup [][]rune, word string) bool {
	alphabetSoup = soup
	wordSlice := []rune(word)
	for y, row := range alphabetSoup {
		for x, letter := range row {
			if letter == wordSlice[0] {
				if wordSearch(x, y, wordSlice, [][]int{{}}) {
					return true
				}
			}
		}
	}
	return false
}

func wordSearch(col int, row int, wordSlice []rune, analyzedCoords [][]int) bool {
	if len(wordSlice) == 0 {
		return true
	}
	if col < 0 || col > len(alphabetSoup[0])-1 || row < 0 || row > len(alphabetSoup)-1 {
		return false
	}
	for _, cords := range analyzedCoords {
		if reflect.DeepEqual(cords, []int{col, row}) {
			return false
		}
	}
	if alphabetSoup[row][col] != wordSlice[0] {
		return false
	}

	analyzedCoords = append(analyzedCoords, []int{col, row})
	upSearch := wordSearch(col, row-1, wordSlice[1:], analyzedCoords)
	downSearch := wordSearch(col, row+1, wordSlice[1:], analyzedCoords)
	leftSearch := wordSearch(col-1, row, wordSlice[1:], analyzedCoords)
	rightSearch := wordSearch(col+1, row, wordSlice[1:], analyzedCoords)

	return upSearch || downSearch || leftSearch || rightSearch
}
