package Config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// CloudCastDBConnect connects to the database and returns the connection object.
func CloudCastDBConnect() (*pgx.Conn, error) {
	databaseUrl := "postgres://postgres:alpintod@localhost:5432/weathercast"
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	// Don't close the connection here! Return it so the caller can use it.
	return conn, nil
}
