// A program to map a URL to a seed phrase

package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	// "regexp"
	"encoding/gob"
	"strconv"
)

var urlMap = map[string]string{}
var reverseLookupMap = map[string]string{}

func main() {

	initMaps("urls.gob", urlMap)
	initMaps("reverseUrls.gob", reverseLookupMap)
	// fmt.Println(mapToPhrase("www.google.com"))
	mapToPhrase("www.google.com")
	fmt.Println(urlMap)
	fmt.Println(reverseLookupMap)
}

// Map a URL to a phrase
func mapToPhrase(link string) string {

	if urlMap[link] != "" {
		return urlMap[link]
	}

	randomPhrase := generatePhrase()
	urlMap[link] = randomPhrase
	reverseLookupMap[randomPhrase] = link

	saveToStorage("urls.gob", urlMap)
	saveToStorage("reverseUrls.gob", reverseLookupMap)

	return randomPhrase
}

// Lookup the URL by providing its relevant phrase
func lookupUrl(phrase string) string {

	return reverseLookupMap[phrase]
}

// Validate URLs
func validateUrl(inputText string) string {

	return ""
}

// Load maps from storage
func initMaps(fileName string, data map[string]string) {

	// Open file
	decodeFile, err := os.Open(fileName)
	if err != nil {
		// panic(err)
		fmt.Println("Running for the first time...")
	}
	defer decodeFile.Close()

	decoder := gob.NewDecoder(decodeFile)

	decoder.Decode(&data)
	// decoder.Decode(&urlMap)
}

// Save changes to storage
func saveToStorage(fileName string, data map[string]string) {

	// Create FIle
	encodeFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	encoder := gob.NewEncoder(encodeFile)

	// Write to file
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}
	encodeFile.Close()
}

// Concatenate random words together and generate a random phrase
func generatePhrase() string {

	phrase := readRandomWordFromFile() + "-" + readRandomWordFromFile() + "-" + strconv.Itoa(rand.IntN(999))
	return phrase
}

// Read a random word from a word list
func readRandomWordFromFile() string {

	randomNumber := rand.IntN(2048)
	wordCount := 0

	file, err := os.Open("words.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if wordCount == randomNumber {
			return scanner.Text()
		}
		wordCount++
	}

	return ""
}
