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

// planetsCmd represents the planets command
var planetsCmd = &cobra.Command{
	Use:   "planets",
	Short: "Get all planets in Starwars with information in command line",
	Long:  `Get all planets in Starwars with information in command line`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, limit, err := api.CheckAllFlags(cmd)
		if err != nil {
			fmt.Println("error with flag:", err)
			return
		}
		utils.PrintVerboseMode(verbose)
		getAllPlanets(limit, verbose)
	},
}

func init() {
	rootCmd.AddCommand(planetsCmd)
}

func getAllPlanets(limit int, verbose bool) {
	resBytes := api.GetAllData("/planets")
	var planets api.Planets
	err := json.Unmarshal(resBytes, &planets)
	if err != nil {
		errString := fmt.Sprint("error when parse planets: ", err)
		panic(errString)
	}
	allPlanets := planets.Planets
	limitPlanets := allPlanets[:]

	if limit > 0 && limit <= len(allPlanets) {
		limitPlanets = allPlanets[:limit]
	}
	for _, p := range limitPlanets {
		fmt.Println("Name:", p.Name)
		if verbose {
			fmt.Println("Climate:", p.Climate)
			fmt.Println("Diameter:", p.Diameter)
		}
		fmt.Println("--------------")
	}
}
