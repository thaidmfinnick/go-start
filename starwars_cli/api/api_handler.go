package api

import (
	"fmt"
	"io"
	"net/http"
)

func GetAllData(typeData string) []byte {
	url := "https://swapi.dev/api"
	fullUrl := fmt.Sprint(url, typeData)
	res, err := http.Get(fullUrl)
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

	return body
}
