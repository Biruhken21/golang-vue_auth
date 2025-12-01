package graphql

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import "last-go/internal/database"

type Resolver struct {
	DB *database.DB
}
