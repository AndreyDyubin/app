package main

import (
	"github.com/labstack/echo"
	"os"
	"syscall"
	"time"
	"go.uber.org/zap"
	"gopkg.in/urfave/cli.v2"
	"github.com/AndreyDyubin/app/logger"
	"os/signal"
	"context"
	"github.com/AndreyDyubin/app/storage"
	"github.com/AndreyDyubin/app/routes"
	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
	"github.com/AndreyDyubin/app/core"
)

func main() {
	var VERSION = "dev"
	var (
		log *zap.SugaredLogger
		db  *reform.DB
	)

	log = logger.NewLogger(zap.DebugLevel).Sugar()

	app := &cli.App{}
	app.Name = "dadataproxy"
	app.Version = VERSION
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "s3bucket", Value: "bucket", EnvVars: []string{"APP_S3BUCKET"}},
		&cli.StringFlag{Name: "s3key", Value: "token", EnvVars: []string{"APP_S3KEY"}},
		&cli.StringFlag{Name: "s3secret", Value: "token", EnvVars: []string{"APP_S3SECRET"}},
	}
	app.Before = func(c *cli.Context) (err error) {
		// инициализация служб и используемых сервисов
		core.UploadService = core.NewUploadService(db, log.Named("upload_service"), c.String("s3bucket"))

		err = storage.ConnectDB()
		if err != nil {
			return err
		}
		err = storage.ConnectS3(c.String("s3key"), c.String("s3secret"))
		return err
	}
	app.Commands = []*cli.Command{
		{
			Name: "api",
			Subcommands: []*cli.Command{
				{
					Name: "run",
					Flags: []cli.Flag{
						&cli.StringFlag{Name: "address", Value: ":8944"},
					},
					Action: func(c *cli.Context) error {
						e := echo.New()
						SetupV1(e) // настройка роутинга

						log.Info("API listen", c.String("address"))

						var appSignal = make(chan struct{}, 2)
						go func() {
							e.Start(c.String("address"))
							appSignal <- struct{}{}
						}()

						osSignal := make(chan os.Signal, 2)
						close := make(chan struct{})
						signal.Notify(
							osSignal,
							os.Interrupt,
							syscall.SIGTERM,
						)

						go func() {

							defer func() {
								close <- struct{}{}
							}()

							select {
							case <-osSignal:
								log.Error("signal completion of the process: OS")
							case <-appSignal:
								log.Error("signal completion of the process: internal (http server, etc..)")
							}

							// TODO: destroy services

							ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
							defer cancel()
							e.Shutdown(ctx)
						}()

						<-close
						os.Exit(0)
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

func SetupV1(e *echo.Echo) {
	e.POST("/upload", routes.Upload)
	e.GET("/download", routes.Download)
	e.GET("/list", routes.List)
}
