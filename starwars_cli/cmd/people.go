/*
Copyright © 2024 thaidmfinnick
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"starwars_cli/api"

	"github.com/spf13/cobra"
)

// peopleCmd represents the people command
var peopleCmd = &cobra.Command{
	Use:   "people",
	Short: "Get all people in Starwars with information in command line",
	Long:  `Get all people in Starwars with information in command line`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, limit, err := api.CheckAllFlags(cmd)
		if err != nil {
			fmt.Println("error with flag:", err)
			return
		}
		if verbose {
			fmt.Println("Running in verbose mode")
			fmt.Println("")
		}
		getAllPeople(limit, verbose)
	},
}

func init() {
	rootCmd.AddCommand(peopleCmd)
}

func getAllPeople(limit int, verbose bool) {
	resBytes := api.GetAllData("/people")
	var people api.People
	err := json.Unmarshal(resBytes, &people)
	if err != nil {
		errString := fmt.Sprint("error when parse people: ", err)
		panic(errString)
	}
	characters := people.Characters
	limitCharacters := characters[:]

	if limit > 0 && limit <= len(characters) {
		limitCharacters = characters[:limit]
	}
	for _, c := range limitCharacters {
		fmt.Println("Name:", c.Name)
		if verbose {
			fmt.Println("Height:", c.Height)
			fmt.Println("Eye color:", c.EyeColor)
		}
		fmt.Println("--------------")
	}
}
