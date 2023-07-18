/*
Copyright © 2023 Phillip Shreves phillipshreves@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// combineCmd represents the combine command
var combineCmd = &cobra.Command{
	Use:   "combine",
	Short: "Converts multiple JSON arrays into a single list",
	Long: `Converts multiple JSON arrays into a single list. 
	The input files must be in the 'input' directory. The input files must be named '*.json'.
	
	The output file will be in the 'output' directory. The output file will be named 'combined.txt'. The output file will be overwritten if it already exists.`,
	Run: combiner,
}

func init() {
	rootCmd.AddCommand(combineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// combineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// combineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func combiner(cmd *cobra.Command, args []string) {
	fmt.Println("combine called")
	files, err := os.ReadDir("./input")
	if err != nil {
		log.Fatal(err)
	}

	var wordArr []string

	for _, f := range files {
		var wordArrRead []string

		content, err := os.ReadFile("./input/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(content, &wordArrRead)
		if err != nil {
			log.Fatal(err)
		}

		wordArrUpdated := wordArrRead

		// we want to make sure that all words are at least 6 characters long for sufficient entropy
		for len(wordArrUpdated[0]) < 6 {
			for i, word := range wordArrUpdated {
				if len(word) < 6 {
					extendedWordArr := extendWord(word)
					wordArrUpdated[i] = extendedWordArr[0]
					wordArrUpdated = append(wordArrUpdated, extendedWordArr...)
				}
			}
		}

		wordArr = append(wordArr, wordArrUpdated...)
		wordArr = randomize(wordArr)
	}

	// randomize the list so that we can use a straight insert into the database
	wordArr = randomize(wordArr)
	wordArr = randomize(wordArr)
	os.WriteFile("./output/combined.txt", []byte(strings.Join(wordArr, "\n")), 0644)
}

func extendWord(word string) []string {
	var wordArr []string
	for char := 'A'; char <= 'Z'; char++ {
		wordArr = append(wordArr, word+string(char))
	}
	return wordArr
}

func randomize(wordArr []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(wordArr), func(i, j int) {
		wordArr[i], wordArr[j] = wordArr[j], wordArr[i]
	})
	return wordArr
}
