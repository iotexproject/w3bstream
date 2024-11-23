package services

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/clickhouse"
)

type chContainer struct {
	testcontainers.Container
}

func SetupClickhouse(dbName string) (*chContainer, string, error) {
	ctx := context.Background()
	dbUser := "default"
	dbPassword := "password"

	clickhouseContainer, err := clickhouse.Run(ctx,
		"clickhouse/clickhouse-server:24.8-alpine",
		clickhouse.WithDatabase(dbName),
		clickhouse.WithUsername(dbUser),
		clickhouse.WithPassword(dbPassword),
	)
	if err != nil {
		return nil, "", err
	}

	dsn, err := clickhouseContainer.ConnectionString(ctx, "secure=false")
	if err != nil {
		return nil, "", err
	}

	return &chContainer{Container: clickhouseContainer}, dsn, nil
}
