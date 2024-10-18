package main

import (
	"context"
	"go-template/module"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Main Code Here
	core := module.NewCoreModule()
	core.Helper.LoggerHelper.LogAndContinue("Henlo %s", os.Getenv("POSTGRES_URL"))
}

func init() {
	godotenv.Load()
	core := module.NewCoreModule()

	core.Database.PostgresDatabase.Open()
	rows, err := core.Database.PostgresDatabase.GetConn().Query(context.Background(), "SELECT * FROM test")
	if err != nil {
		core.Helper.LoggerHelper.LogErrAndExit(1, err, "Failed Query")
	}

	for rows.Next() {
		type placeholder struct {
			id   int
			text string
		}

		var p placeholder
		err := rows.Scan(&p.id, &p.text)
		if err != nil {
			core.Helper.LoggerHelper.LogErrAndExit(1, err, "Failed Scanning")
		}

		core.Helper.LoggerHelper.LogAndContinue("ID: %d\nData: %v\n\n", p.id, p.text)
	}
}
