package main

import (
	"fmt"

	"github.com/estebanpavansarquis/word-search/utils"
)

func main() {
	wordSearchInput := utils.GetWordSearchInput()
	fmt.Println("\nMy wordSearch input:", wordSearchInput)
	for _, word := range wordSearchInput.Words {
		fmt.Println(fmt.Sprintf("Result of running wordSearch serching for %s is: %t", word, utils.IsPresent(wordSearchInput.AlphabetSoup, word)))
	}
}
