package config

import "os"

var (
	Host string
	Port string
	DBHost string
	DBPort string
)

func Init(){
	os.Setenv("PORT", "8081")
	os.Setenv("LOCAL", "127.0.0.1:")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "5432")
	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
}

