package dbconfig

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/karim-w/techunicorn-task/ent"
	"go.uber.org/fx"
)

type DbSettingsConfig struct {
	Server       string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

func LoadFromEnv(key string) string {
	return os.Getenv(key)
}

func LoadDBVarsFromEnv() *DbSettingsConfig {
	return &DbSettingsConfig{
		Server:       LoadFromEnv("DBServer"),
		Port:         LoadFromEnv("DBPort"),
		User:         LoadFromEnv("DBUser"),
		Password:     LoadFromEnv("DBPassword"),
		DatabaseName: LoadFromEnv("DBName"),
	}
}
func StartEntConnection() *ent.Client {
	// Build connection string
	config := LoadDBVarsFromEnv()
	connString := fmt.Sprintf("postgresql://%s:%s@%s/%s", config.User, config.Password, config.Server, config.DatabaseName)
	client := Open(connString)
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
		return nil
	}
	return client
}

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

var DBModule = fx.Option(fx.Provide(StartEntConnection))
