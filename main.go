package main

import (
	"context"
	"net"
	"net/http"

	"github.com/lisyaoran51/uberfx-test/echo"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewZapLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	// config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	// encoderCfg := zapcore.EncoderConfig{
	// 	MessageKey:     "msg",
	// 	LevelKey:       "level",
	// 	NameKey:        "logger",
	// 	EncodeLevel:    zapcore.LowercaseLevelEncoder,
	// 	EncodeTime:     zapcore.ISO8601TimeEncoder,
	// 	EncodeDuration: zapcore.StringDurationEncoder,
	// }
	// core := zapcore.NewCore(zapcore.NewJSONEncoder(config), os.Stdout, zap.DebugLevel)
	// return zap.New(core).WithOptions()
	return logger
}

// NewHTTPServer builds an HTTP server that will begin serving requests
// when the Fx application starts.
func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":8089", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func main() {

	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			NewHTTPServer,
			fx.Annotate(
				echo.NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			// echo.NewEchoHandler,
			NewZapLogger,
			fx.Annotate(
				echo.NewEchoHandler,
				fx.As(new(echo.Route)),
				fx.ResultTags(`group:"routes"`),
			),
			fx.Annotate(
				echo.NewHelloHandler,
				fx.As(new(echo.Route)),
				fx.ResultTags(`group:"routes"`),
			),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
