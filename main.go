package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	file := os.Args[1]
	var gameIds []string

	data, err := os.Open(file)
	handleError(err, "open csv")
	r := csv.NewReader(bufio.NewReader(data))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		handleError(err, "read csv")

		gameIds = append(gameIds, record[0])
	}

	baseURL := "http://www.boardgamegeek.com/xmlapi2/thing?id="
	url := baseURL + strings.Join(gameIds, ",")

	response, err := http.Get(url)
	handleError(err, "get data")

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	handleError(err, "read data")

	gameList := parseXML(body)

	jsonGameList := toJSON(gameList)

	js, _ := json.MarshalIndent(jsonGameList, "", "  ")

	ioutil.WriteFile("bgg-data.json", js, 0644)
}

func handleError(err error, action string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "bgg - %s: %v\n", action, err)
		os.Exit(1)
	}
}
