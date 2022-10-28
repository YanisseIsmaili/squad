package main 

import (
	"bufio"
	"os"
    
)

func main() {
	words := importtxt()

	SearchWords(suj2, words)
}


func importtxt() []string {

	var word []string

	readFile, _ := os.Open("words.txt")

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		word = append(word, fileScanner.Text())
	}

	readFile.Close()

	return word
}


var suj2 = [10][10]rune{
	{'c', 'o', 't', 'd', 't', 'r', 's', 'n', 'e', 'c'},
	{'r', 'e', 'e', 'o', 't', 'e', 'o', 'h', 'u', 'c'},
	{'ê', 'u', 'h', 'h', 't', 'r', 'l', 'a', 'o', 'a'},
	{'p', 'f', 'w', 'a', 'e', 'r', 'e', 'a', 'a', 'f'},
	{'e', 's', 't', 'p', 'r', 't', 'a', 'e', 'r', 'é'},
	{'s', 'a', 'p', 'y', 'u', 'i', 'o', 'e', 's', 'd'},
	{'p', 'a', 'e', 'o', 'i', 't', 'c', 'd', 't', 'e'},
	{'n', 'n', 'c', 'n', 'c', 'w', 'b', 'o', 'b', 'n'},
	{'f', 'o', 'u', 'r', 'c', 'h', 'e', 't', 't', 'e'},
	{'b', 'a', 'i', 'e', 's', 't', 'h', 'n', 'w', 's'},
}