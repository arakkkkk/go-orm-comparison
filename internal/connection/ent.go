package connection

import (
	"log"
	"orm/ent"
	"orm/internal/config"
)

func EntOpen() (*ent.Client, error) {
	dsn := config.DSN()
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error open mysql ent client: %v\n", err)
		return nil, err
	}
	return client, nil
}

