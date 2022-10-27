package main

import (
    "bufio"
    "os"
)

func importtxt() []string {

    var word []string

    readFile, _ := os.Open("words.txt")

    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
        word = append(word,fileScanner.Text())
    }

    readFile.Close()

    return word
}