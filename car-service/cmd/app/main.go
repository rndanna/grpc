package main

import (
	"car-service/internal/storage/postgresql"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
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

	db := postgresql.New(pool)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// grpcServer := grpc.NewServer()
	// gatewayGRPCService := gatewayGRPCService.New()
	// gatewayService.RegisterGatewayServiceServer(grpcServer, gatewayGRPCService)
	// reflection.Register(grpcServer)

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}
