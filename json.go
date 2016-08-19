package main

import (
	"html"
	"strconv"
)

type Game struct {
	Name        string
	Published   int
	Thumbnail   string
	Image       string
	MinPlayers  int
	MaxPlayers  int
	PlayingTime int
	MinPlaytime int
	MaxPlaytime int
	MinAge      int
	Description string
	Categories  []string
	Mechanics   []string
	Rank        int
	Rating      float64
}

func toJSON(query Query) []Game {
	var gameList []Game

	for _, item := range query.GameList[:] {
		var game Game

		game.Name = getPrimaryName(item)
		game.Published = toInt(item.Published.Value)
		game.Thumbnail = item.Thumbnail
		game.Image = item.Image
		game.MinPlayers = toInt(item.MinPlayers.Value)
		game.MaxPlayers = toInt(item.MaxPlayers.Value)
		game.MinPlaytime = toInt(item.MinPlaytime.Value)
		game.MaxPlaytime = toInt(item.MaxPlaytime.Value)
		game.PlayingTime = toInt(item.PlayingTime.Value)
		game.MinAge = toInt(item.MinAge.Value)
		game.Description = html.UnescapeString(item.Description)
		game.Categories = getLink(item, "boardgamecategory")
		game.Mechanics = getLink(item, "boardgamemechanic")
		game.Rank = toInt(getPrimaryRank(item))
		game.Rating = toFloat(item.Rating.Value)

		gameList = append(gameList, game)
	}

	return gameList
}

func getPrimaryName(item Item) string {
	for _, entry := range item.Names[:] {
		if entry.Type == "primary" {
			return entry.Value
		}
	}
	return ""
}

func getPrimaryRank(item Item) string {
	for _, entry := range item.Ranks[:] {
		if entry.Name == "boardgame" {
			return entry.Value
		}
	}
	return ""
}

func toInt(str string) int {
	num, err := strconv.Atoi(str)
	handleError(err, "convert string")
	return num
}

func toFloat(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	handleError(err, "convert string")
	return num
}

func getLink(item Item, label string) []string {
	var filtered []string
	for _, link := range item.Links[:] {
		if link.Type == label {
			filtered = append(filtered, link.Value)
		}
	}
	return filtered
}
