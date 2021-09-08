package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/ydb-platform/ydb-go-sdk/v3/connect"

	"github.com/ydb-platform/ydb-go-examples/pkg/cli"
)

type Command struct {
}

func (cmd *Command) Run(ctx context.Context, params cli.Parameters) error {
	connectCtx, cancel := context.WithTimeout(ctx, params.ConnectTimeout)
	defer cancel()
	db, err := connect.New(
		connectCtx,
		params.ConnectParams,
		connect.WithAnonymousCredentials(),
	)
	if err != nil {
		return fmt.Errorf("connect error: %w", err)
	}
	defer db.Close()

	// work with db instance

	return nil
}

func (cmd *Command) ExportFlags(context.Context, *flag.FlagSet) {}
