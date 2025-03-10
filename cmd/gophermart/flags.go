package main

import (
	"flag"
	"os"
)

var (
	flagRunAddr            string
	flagLogLevel           string
	flagAccuralSystemAddr  string
	flagDBConnectionString string
)

// -d "host=localhost user=metrics password=userpassword dbname=metrics sslmode=disable"
func parseFlags() error {
	flag.StringVar(&flagRunAddr, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&flagAccuralSystemAddr, "r", "localhost:8090", "address of accural system")
	flag.StringVar(&flagLogLevel, "l", "info", "log level")
	flag.StringVar(&flagDBConnectionString, "d", "", "connetction string for postgress db")
	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDRESS"); envRunAddr != "" {
		flagRunAddr = envRunAddr
	}

	if envAccuralSystemAddr := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); envAccuralSystemAddr != "" {
		flagAccuralSystemAddr = envAccuralSystemAddr
	}

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		flagLogLevel = envLogLevel
	}

	if envDBConnectionString := os.Getenv("DATABASE_URI"); envDBConnectionString != "" {
		flagDBConnectionString = envDBConnectionString
	}

	return nil
}
