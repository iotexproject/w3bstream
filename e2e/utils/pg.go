package utils

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

type pgContainer struct {
	testcontainers.Container
}

func SetupPostgres(dbName string) (*pgContainer, string, error) {
	ctx := context.Background()

	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.BasicWaitStrategies(),
	)

	if err != nil {
		return nil, "", err
	}

	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, "", err
	}

	fmt.Println(connStr)

	// test connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, "", err
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, "", err
	}

	return &pgContainer{Container: postgresContainer}, connStr, nil
}
