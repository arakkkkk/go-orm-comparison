package main

import (
	// "orm/internal/crud"
	"orm/internal/migration"
)

func main() {
  migration.Migrate()
  crud.Create()
}
