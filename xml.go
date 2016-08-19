package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Value struct {
	Value string `xml:"value,attr"`
	Type  string `xml:"type,attr,omitempty"`
}

type Result struct {
	Value    string `xml:"value,attr"`
	NumVotes string `xml:"numvotes,attr"`
	Level    string `xml:"level,attr,omitempty"`
}

type Link struct {
	Type  string `xml:"type,attr"`
	ID    string `xml:"id,attr"`
	Value string `xml:"value,attr"`
}

type Rank struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Item struct {
	Names       []Value `xml:"name"`
	Published   Value   `xml:"yearpublished"`
	Thumbnail   string  `xml:"thumbnail"`
	Image       string  `xml:"image"`
	MinPlayers  Value   `xml:"minplayers"`
	MaxPlayers  Value   `xml:"maxplayers"`
	PlayingTime Value   `xml:"playingtime"`
	MinPlaytime Value   `xml:"minplaytime"`
	MaxPlaytime Value   `xml:"maxplaytime"`
	MinAge      Value   `xml:"minage"`
	Description string  `xml:"description"`
	Links       []Link  `xml:"link"`
	Rating      Value   `xml:"statistics>ratings>average"`
	Ranks       []Rank  `xml:"statistics>ratings>ranks>rank"`
}

type Query struct {
	GameList []Item `xml:"item"`
}

func parseXML(data []byte) Query {
	var q Query

	err := xml.Unmarshal([]byte(data), &q)
	if err != nil {
		fmt.Printf("bgg: error unmarshaling: %v", err)
		os.Exit(1)
	}

	return q
}
