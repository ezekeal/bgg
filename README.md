# bgg
fetches data from board game geek and provides a json file as output

## How it works

First create a csv file with the id numbers of the Board games you want data for.

You can find these by searching for the game on [Board Game Geek](boardgamegeek.com)

boardgamegeek.com/boardgame/**40692**/small-world

now feed the file to bgg

`$ bgg mygames.csv`

It will create `bgg-data.json` with the game data
