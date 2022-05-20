package router

import (
	"context"
	"github.com/jim/logger"
	"github.com/jim/services/router/apis"
	"github.com/jim/services/router/conf"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ServerStartOptions struct {
	config string
	data   string
}

// NewServerStartCmd creates a new http server command
func NewServerStartCmd(ctx context.Context, version string) *cobra.Command {
	opts := &ServerStartOptions{}
	cmd := &cobra.Command{
		Use:   "router",
		Short: "Start a router",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunServerStart(ctx, opts, version)
		},
	}
	cmd.PersistentFlags().StringVarP(&opts.config, "config", "c", "./router/conf.yaml", "Config file")
	cmd.PersistentFlags().StringVarP(&opts.data, "data", "d", "./router/data", "data path")
	return cmd
}

// RunServerStart run http server
func RunServerStart(ctx context.Context, opts *ServerStartOptions, version string) error {
	config, err := conf.Init(opts.config)
	if err != nil {
		return err
	}
	_ = logger.Init(logger.Settings{
		Level:    "info",
		Filename: "./data/router.log",
	})

	app := iris.New()

	app.Use(func(ctx iris.Context) {
		ctx.Header("Server", "Iris MongoDB/"+version)
		ctx.Next()
	})

	app.Get("/health", func(ctx iris.Context) {
		_, _ = ctx.WriteString("ok")
	})

	storeAPI := app.Party("/api/messageTemplate")
	{
		movieHandler := apis.NewMessageTemplateHandler()
		storeAPI.Post("", movieHandler.Add)
	}

	userAPI := app.Party("/api/user")
	{
		userHandler := apis.NewUserSynHandler()
		userAPI.Get("/syn", userHandler.Syn)
	}

	businessAPI := app.Party("/api/business")
	{
		businessHandler := apis.NewBusinessHandler()
		businessAPI.Post("/create", businessHandler.Add)
		businessAPI.Get("/get", businessHandler.Get)
	}

	logrus.Infof("load regions - %v", "ds")

	// Start server
	return app.Listen(config.Listen, iris.WithOptimizations)
}
