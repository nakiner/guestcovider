package main

import (
	"context"
	"fmt"
	"github.com/nakiner/guestcovider/internal/database"
	"github.com/nakiner/guestcovider/internal/userRepository"
	"net/http"
	"os"

	"github.com/nakiner/guestcovider/pkg/health"
	"github.com/nakiner/guestcovider/pkg/user"

	"github.com/go-kit/kit/log/level"
	"github.com/nakiner/guestcovider/configs"
	"github.com/nakiner/guestcovider/internal/server"
	"github.com/nakiner/guestcovider/tools/logging"
	"github.com/nakiner/guestcovider/tools/metrics"
	"github.com/nakiner/guestcovider/tools/sentry"
	"github.com/nakiner/guestcovider/tools/tracing"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load config
	cfg := configs.NewConfig()
	if err := cfg.Read(); err != nil {
		fmt.Fprintf(os.Stderr, "read config: %s", err)
		os.Exit(1)
	}
	// Print config
	if err := cfg.Print(); err != nil {
		fmt.Fprintf(os.Stderr, "read config: %s", err)
		os.Exit(1)
	}

	logger, err := logging.NewLogger(cfg.Logger.Level, cfg.Logger.TimeFormat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to init logger: %s", err)
		os.Exit(1)
	}
	ctx = logging.WithContext(ctx, logger)

	if cfg.Tracer.Enabled {
		tracer, closer, err := tracing.NewJaegerTracer(
			ctx,
			fmt.Sprintf("%s:%d", cfg.Tracer.Host, cfg.Tracer.Port),
			cfg.Tracer.Name,
		)
		if err != nil {
			level.Error(logger).Log("err", err, "msg", "failed to init tracer")
		}
		defer closer.Close()
		ctx = tracing.WithContext(ctx, tracer)
	}
	if cfg.Sentry.Enabled {
		if err := sentry.NewSentry(cfg); err != nil {
			level.Error(logger).Log("err", err, "msg", "failed to init sentry")
		}
	}

	if cfg.Metrics.Enabled {
		ctx = metrics.WithContext(ctx)
	}

	dbConn, err := database.Connect(ctx, cfg.Postgres)
	if err != nil {
		level.Error(logger).Log("msg", "db connect error", "err", err)
	}

	defer dbConn.Close()

	userRepo := userRepository.NewUserDBRepository(dbConn)
	if cfg.Tracer.Enabled {
		userRepo = userRepository.NewTracingRepository(ctx, userRepo)
	}

	healthService := initHealthService(ctx, cfg)
	userService := initUserService(ctx, cfg, userRepo)

	s, err := server.NewServer(
		server.SetConfig(cfg),
		server.SetLogger(logger),
		server.SetHandler(
			map[string]http.Handler{
				"health": health.MakeHTTPHandler(ctx, healthService),
				"user":   user.MakeHTTPHandler(ctx, userService),
			}),
		server.SetGRPC(
			health.JoinGRPC(ctx, healthService),
			user.JoinGRPC(ctx, userService),
		),
	)
	if err != nil {
		level.Error(logger).Log("init", "server", "err", err)
		os.Exit(1)
	}
	defer s.Close()

	if err := s.AddHTTP(); err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}

	if err = s.AddGRPC(); err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}

	if err = s.AddMetrics(); err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}

	s.AddSignalHandler()
	s.Run()
}

func initHealthService(ctx context.Context, cfg *configs.Config) health.Service {
	healthService := health.NewHealthService()
	if cfg.Metrics.Enabled {
		healthService = health.NewMetricsService(ctx, healthService)
	}
	healthService = health.NewLoggingService(ctx, healthService)
	if cfg.Tracer.Enabled {
		healthService = health.NewTracingService(ctx, healthService)
	}
	if cfg.Sentry.Enabled {
		healthService = health.NewSentryService(healthService)
	}
	return healthService
}

func initUserService(ctx context.Context, cfg *configs.Config, repo userRepository.Repository) user.Service {
	userService := user.NewUserService(repo)
	if cfg.Metrics.Enabled {
		userService = user.NewMetricsService(ctx, userService)
	}
	userService = user.NewLoggingService(ctx, userService)
	if cfg.Tracer.Enabled {
		userService = user.NewTracingService(ctx, userService)
	}
	if cfg.Sentry.Enabled {
		userService = user.NewSentryService(userService)
	}
	return userService
}
