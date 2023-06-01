package main

import (
	"os"
	"log"
	"encoding/csv"
)

func main() {
	parsePlayersTest()
}

func parsePlayersTest() error {
	playersTypes := make(map[string]string)

	playersTypes["playerID"] = "VARCHAR"
	playersTypes["birthYear"] = "INTEGER"
	playersTypes["nameFirst"] = "VARCHAR"
	playersTypes["nameLast"] = "VARCHAR"

	file, err := os.Open("../baseballdatabank/core/People.csv")
	if err != nil {
		log.Println(err)
	}

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	columnsRow := records[0]

	table := ""
	tableValues := ""
	for _, columnName := range columnsRow {
		columnType, ok := playersTypes[columnName]
		if ok {
			if table != "" {
				table += ", " 
			}
			table += columnName + " " + columnType 
			tableValues += columnName
		}
	}

	log.Println(table)

	playerRows := records[1:]

	for _, playerRow := range playerRows {
		values := ""
		for i, playerCell := range playerRow {
			columnName := columnsRow[i]
		  _, ok := playersTypes[columnName]
			if (ok) {
				if values == "" {
					values += playerCell
				} else {
					values += ", " + playerCell
				}
			}
		}
		log.Println(values)
	}

	return nil
}
