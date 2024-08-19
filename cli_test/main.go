package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type People struct {
	Count      int    `json:"count"`
	Next       string `json:"next"`
	Previous   int    `json:"previous"`
	Characters []struct {
		Name      string `json:"name"`
		Height    string `json:"height"`
		Mass      string `json:"mass"`
		EyeColor  string `json:"eye_color"`
		BirthYear string `json:"birth_year"`
		Gender    string `json:"gender"`
		HairColor string `json:"hair_color"`
		SkinColor string `json:"skin_color"`
		Url       string `json:"url"`
	} `json:"results"`
}

func main() {
	limit := flag.Int("limit", -1, "Limit your results")
	verbose := flag.Bool("verbose", false, "Log all data api output")
	flag.Parse()

	url := "https://swapi.dev/api/people"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Something went wrong with API")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		errString := fmt.Sprint("error when read data api: ", err)
		panic(errString)
	}

	var people People
	err = json.Unmarshal(body, &people)
	if err != nil {
		errString := fmt.Sprint("error when parse json: ", err)
		panic(errString)
	}

	characters := people.Characters
	limitCharacters := characters[:]

	if *limit > 0 && *limit <= len(characters) {
		limitCharacters = characters[:*limit]
	}
	for _, c := range limitCharacters {
		fmt.Println("Name:", c.Name)
		if *verbose {
			fmt.Println("Height:", c.Height, "|", "Eye color:", c.EyeColor)
		}
		fmt.Println("-----------------------------------------")
	}
}
