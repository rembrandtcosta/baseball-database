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

	models := make([]interface{}, 0)

	model := getModel(getModelReflection(p.CreatePlayerParams{}))
	models = appendStruct(models, []interface{}{model})
	model = getModel(getModelReflection(p.CreateFranchiseParams{}))
	models = appendStruct(models, []interface{}{model})

	locations := []string{"./baseballdatabank/core/People.csv", "./baseballdatabank/core/TeamsFranchises.csv"}

	for i := 0; i < len(locations); i++ {
		file, err := os.Open(locations[i])
		if err != nil {
			log.Println(err)
		}
		go parseTable(ctx, db, models[i], file)
	}



	log.Print("Listening 8000")
	r := mux.NewRouter()
	r.HandleFunc("/", h.BlogHandler)
	r.HandleFunc("/player", h.PlayerHandler)
	log.Fatal(http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r)))

	return nil
}

func appendStruct(slice []interface{}, elem []interface{}) []interface{} {
	j := len(slice) + 1
	newSlice := make([]interface{}, j)
	for i, e := range slice {
		newSlice[i] = e
	}
	for _, e := range elem {
		newSlice[j-1] = e
	}
	return newSlice 
}

func createModel(q *p.Queries, ctx context.Context, model interface{}) {
	switch model.(type) {
	case p.CreatePlayerParams:
		q.CreatePlayer(ctx, model.(p.CreatePlayerParams))
	case p.CreateFranchiseParams:
		log.Println(model)
		q.CreateFranchise(ctx, model.(p.CreateFranchiseParams))
	}
}

func getModel(model reflect.Value) interface{} {
	switch model.Interface().(type) {
	case p.CreatePlayerParams:
		m := model.Interface().(p.CreatePlayerParams)
		return m
	case p.CreateFranchiseParams:
		m := model.Interface().(p.CreateFranchiseParams)
		return m
	}
	return model
}

func getModelReflection(model interface{}) reflect.Value {
  switch model.(type) {
	case p.CreatePlayerParams:
		m := model.(p.CreatePlayerParams)
		return reflect.ValueOf(&m).Elem()
	case p.CreateFranchiseParams:
		m := model.(p.CreateFranchiseParams)
		return reflect.ValueOf(&m).Elem()
	}
  return reflect.ValueOf(&model).Elem() 
}

func parseTable(ctx context.Context, db *sql.DB, model interface{}, file *os.File) error {
	defer db.Close()

	q := p.New(db)


	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	rows := records[1:]

	s := getModelReflection(model)
	log.Println(s)

	for _, row := range rows {
		log.Println(row)
		for i, cell := range row {
			typeField := s.Field(i).Type()
			setFieldWithType(s, i, typeField.String(), cell)
		}
		model = getModel(s)
		createModel(q, ctx, model)
	}

	return nil
}

func setFieldWithType(structure reflect.Value, index int, typeField string, val string) {
	if typeField == "string" {
		structure.Field(index).SetString(val)
	} else if typeField == "sql.NullInt32" {
		s, _ := strconv.Atoi(val)
		structure.Field(index).Set(reflect.ValueOf(sql.NullInt32{Int32: int32(s), Valid: true}))
	} else if typeField == "sql.NullString" {
		structure.Field(index).Set(reflect.ValueOf(sql.NullString{String: val, Valid: true}))
	} else if typeField == "sql.NullTime" {
		s, _ := time.Parse("2006-01-02", val)
		structure.Field(index).Set(reflect.ValueOf(sql.NullTime{Time: s, Valid: true}))
	}
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
