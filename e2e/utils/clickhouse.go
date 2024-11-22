package utils

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/clickhouse"
)

type chContainer struct {
	testcontainers.Container
}

func SetupClickhouse(dbName string) (c *chContainer, endpoint, passwd string, err error) {
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
		return nil, "", "", err
	}

	endpoint, err = clickhouseContainer.ConnectionHost(ctx)
	if err != nil {
		return nil, "", "", err
	}

	return &chContainer{Container: clickhouseContainer}, endpoint, dbPassword, nil
}
