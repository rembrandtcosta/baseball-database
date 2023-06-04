package main

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/rembrandtcosta/baseball-database/backend/app"
)

func main() {
	app.Run(context.Background())
}
