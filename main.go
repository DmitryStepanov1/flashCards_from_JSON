package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"unicode"
)

func containsCyrillicOrLatin(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Cyrillic, r) || unicode.Is(unicode.Latin, r) {
			return true
		}
	}
	return false
}

func randomWord(m map[string]string) string {
	//k := rand.Intn(len(m))

	for i, v := range m {
		s := fmt.Sprintf("Переведи %s:", i)
		fmt.Println(s)
		return v
	}

	return ""

}

func main() {

	// Create a map to store input values
	inputMap := make(map[string]string)

	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

jumpTo1:

	fmt.Println("Enter json file name to parse:")

	// Scan for the next token (which is a line)
	scanner.Scan()

	// Check if the file name is valid
	regex := regexp.MustCompile(`^[a-zA-Z0-9_-]+\.json$`)

	inputString := scanner.Text()

	regex.MatchString(inputString)

	if regex.MatchString(scanner.Text()) == false {
		fmt.Println("File is wrong. Try again.")
		goto jumpTo1
	}

	jsonFile, err := os.Open(inputString)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened json-file")
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &inputMap)

	fmt.Println(inputMap)

	fmt.Println("Data from JSON file stored in inputMap:")
	for key, value := range inputMap {
		fmt.Printf("%s: %s\n", key, value)
	}

	for {
		v := randomWord(inputMap)

		scanner := bufio.NewScanner(os.Stdin)

	jumpTo2:

		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		// Check if the user wants to finish entering input
		if input == "done" {
			break
		} else if input == v {
			fmt.Println("Correct! Try next word")
			continue
		} else {
			fmt.Println("Wrong, try again")
			goto jumpTo2
		}

	}

	// Print the map
	fmt.Println("Map contents:")
	for key, value := range inputMap {
		fmt.Printf("%s: %s\n", key, value)
	}
}
