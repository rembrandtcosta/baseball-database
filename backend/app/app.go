package app

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/rembrandtcosta/baseball-database/backend/app/database"
	h "github.com/rembrandtcosta/baseball-database/backend/app/handlers"
	p "github.com/rembrandtcosta/baseball-database/backend/app/postgresql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)



func Run(ctx context.Context) error {
	log.Print("Prepare db...")

	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := prepare(db); err != nil {
		log.Fatal(err)
	}

	parsePlayers(ctx, db)

	log.Print("Listening 8000")
	r := mux.NewRouter()
	r.HandleFunc("/", h.BlogHandler)
	r.HandleFunc("/player", h.PlayerHandler)
	log.Fatal(http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r)))

	return nil
}

func parsePlayers(ctx context.Context, db *sql.DB) error {
	defer db.Close()

	q := p.New(db)

	
	playersTypes := make(map[string]string)

	playersTypes["playerID"] = "VARCHAR"
	playersTypes["birthYear"] = "INTEGER"
	playersTypes["nameFirst"] = "VARCHAR"
	playersTypes["nameLast"] = "VARCHAR"

	file, err := os.Open("./baseballdatabank/core/People.csv")
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
				tableValues += ", "
			}
			table += columnName + " " + columnType 
			tableValues += columnName
		}
	}

	log.Println(table)

	_, err = q.CreatePlayer(ctx, p.CreatePlayerParams{
		Playerid: "teste01",
		Birthyear:  sql.NullInt32{Int32: 2001, Valid: true},
		Namefirst: sql.NullString{String: "Test", Valid: true},
		Namelast: sql.NullString{String: "Teste", Valid: true},
	})

	playerRows := records[1:]

	for _, playerRow := range playerRows {
	  cntField := 0
		player := p.CreatePlayerParams{}
		for i, playerCell := range playerRow {
			columnName := columnsRow[i]
		  _, ok := playersTypes[columnName]
			if (ok) {
				//field := strings.Title(columnName)
				
				s := reflect.ValueOf(&player).Elem()
				log.Println(cntField)
				typeField := s.Field(cntField).Type()

				if typeField.String() == "string" {
					s.Field(cntField).SetString(playerCell)
				} else if typeField.String() == "sql.NullInt32" {
					val, _ := strconv.Atoi(playerCell)
					s.Field(cntField).Set(
						reflect.ValueOf(sql.NullInt32{Int32: int32(val),  Valid: true}))
				} else if typeField.String() == "sql.NullString" {
					s.Field(cntField).Set(reflect.ValueOf(sql.NullString{String: playerCell, Valid: true}))
				}

				cntField++
			}
		}
		_, err = q.CreatePlayer(ctx, player)	
		/*if _, err := db.Exec("INSERT INTO players" +  " VALUES (" + values + ")"); err != nil {
			return err
		}*/
	}

	return nil
}

func prepare(db *sql.DB) error {
	for i := 0; i < 60; i++ {
		if err := db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	if _, err := db.Exec("DROP TABLE IF EXISTS blog"); err != nil {
		return err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS blog (id SERIAL, title VARCHAR)"); err != nil {
		return err
	}

	for i := 0; i < 5; i++ {
		if _, err := db.Exec("INSERT INTO blog (title) VALUES ($1);", fmt.Sprintf("Blog post #%d", i)); err != nil {
			return err
		}
	}
	return nil
}


