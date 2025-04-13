package dataBase

import (
	"context"
	"fmt"
	"log"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/configs"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var DB *pgxpool.Pool

func ConnectToDB() *pgxpool.Pool {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName)
	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable pool_max_conns=10", configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPassword, configs.DBName)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse database config: %v", err)
	}

	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}
	log.Println("Connected to the database")

	err = applyMigrations(dsn)
	if err != nil {
		log.Fatalf("error for apply migration: %v", err)
	}
	return DB
}
