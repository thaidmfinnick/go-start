/*
Copyright © 2024 thaidmfinnick
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"starwars_cli/api"
	"starwars_cli/utils"

	"github.com/spf13/cobra"
)

// spaceshipsCmd represents the spaceships command
var spaceshipsCmd = &cobra.Command{
	Use:   "ships",
	Short: "Get all spaceships in Starwars with information in command line",
	Long:  `Get all spaceships in Starwars with information in command line and see what happened with terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, limit, err := api.CheckAllFlags(cmd)
		if err != nil {
			fmt.Println("error with flag:", err)
			return
		}
		utils.PrintVerboseMode(verbose)
		getAllShips(limit, verbose)
	},
}

func init() {
	rootCmd.AddCommand(spaceshipsCmd)
}

func getAllShips(limit int, verbose bool) {
	resBytes := api.GetAllData("/starships")
	var ships api.SpaceShips
	err := json.Unmarshal(resBytes, &ships)
	if err != nil {
		errString := fmt.Sprint("error when parse spaceship: ", err)
		panic(errString)
	}
	starShips := ships.Ships
	limiStarShips := starShips[:]

	if limit > 0 && limit <= len(starShips) {
		limiStarShips = starShips[:limit]
	}
	for _, s := range limiStarShips {
		fmt.Println("Name:", s.Name)
		if verbose {
			fmt.Println("Model:", s.Model)
			fmt.Println("Manufacturer:", s.Manufacturer)
		}
		fmt.Println("--------------")
	}
}
