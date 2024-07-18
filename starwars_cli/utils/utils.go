package utils

import "fmt"

func PrintVerboseMode(verbose bool) {
	if verbose {
		fmt.Println("Running in verbose mode")
		fmt.Println("")
	}
}
