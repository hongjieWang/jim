package main

import (
	"context"
	"flag"
	"github.com/jim/logger"
	"github.com/jim/services/router"
	"github.com/spf13/cobra"
)

const version = "v1.0"

func main() {
	flag.Parse()

	root := &cobra.Command{
		Use:     "jim",
		Version: version,
		Short:   "July IM Cloud",
	}
	ctx := context.Background()

	root.AddCommand(router.NewServerStartCmd(ctx, version))

	if err := root.Execute(); err != nil {
		logger.WithError(err).Fatal("Could not run command")
	}
}
