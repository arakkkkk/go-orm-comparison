package migration

import (
	"context"
	"log"
	"orm/internal/connection"
)

// table definition is in orm/ent/schema/user
/*
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").
            Default("unknown"),
        field.Int("age").
            Positive(),
        field.Bool("is_active").
            Default(false),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
*/

// Fields of the User.
func EntOpen() {
	client, err := connection.EntOpen()
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
