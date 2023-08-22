package crud

import (
	"log"
	"orm/ent/user"
	"orm/internal/connection"
)

func EntCreate() {
	db, _ := connection.EntOpen()
	u, err := db.User.
		Create().
		SetAge(30).
		SetName("a8m").
		SetIsActive(true).
		Save(nil)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
}

func QueryUser() {
	db, _ := connection.EntOpen()
	u, err := db.User.
		Query().
		Where(user.Name("a8m")).
		Only(nil)
	if err != nil {
		log.Fatalf("failed querying user: %v", err)
	}
	log.Println("user returned: ", u)
}
