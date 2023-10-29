package main

import (
	handle "engine-service/internal/server/nats"
	"engine-service/internal/storage/postgresql"

	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Databse struct {
		Postgresql struct {
			Host     string `mapstructure:"host"`
			Port     string `mapstructure:"port"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
			Name     string `mapstructure:"name"`
		} `mapstructure:"postgresql"`
	} `mapstructure:"database"`
}

func main() {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Databse.Postgresql.Username,
		cfg.Databse.Postgresql.Password,
		cfg.Databse.Postgresql.Host,
		cfg.Databse.Postgresql.Port,
		cfg.Databse.Postgresql.Name,
	)

	pool, err := sql.Open("postgres",
		dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	nc, _ := nats.Connect(nats.DefaultURL)

	jsx, err := nc.JetStream()
	if err != nil {
		log.Fatalf("err nc.JetStream, %v", err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatalf("err jetstream.New, %v", err)
	}

	db := postgresql.New(pool)
	handle := handle.New(db, js, jsx)

	handle.Worker()

	e := echo.New()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}
