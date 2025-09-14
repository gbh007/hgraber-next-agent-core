package main

import (
	"context"
	"fmt"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/gbh007/hgraber-next-agent-core/application/agent"
	"github.com/gbh007/hgraber-next-agent-core/config"
	"github.com/gbh007/hgraber-next-agent-core/dataprovider/loader"
	"github.com/gbh007/hgraber-next-agent-core/dataprovider/webcache"
	"github.com/gbh007/hgraber-next-agent-core/domain/hgraber"
	"github.com/gbh007/hgraber-next-agent-core/request"
	"github.com/gbh007/hgraber-next/adapters/metric"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	agent.Serve(ctx, func(
		ctx context.Context,
		logger *slog.Logger,
		cfg config.Config[config.Parsers],
		async agent.Async,
		metricProvider *metric.MetricProvider,
	) ([]hgraber.Parser, error) {
		if cfg.Parsers == nil {
			logger.DebugContext(ctx, "nil parser config, skipping")

			return []hgraber.Parser{}, nil
		}

		var cache request.Cache

		if cfg.Parsers.Cache.Enabled {
			wc, err := webcache.New(
				cfg.Parsers.Cache.Path,
				logger,
				cfg.Parsers.Cache.TTL,
				cfg.Parsers.Cache.CleanInterval,
				metricProvider,
			)
			if err != nil {
				return nil, fmt.Errorf("create web cache: %w", err)
			}

			async.RegisterRunner(ctx, wc)
			cache = wc
		}

		return loader.NewDefaultParsers(
			logger,
			cfg.Parsers.HG4Token,
			cfg.Application.ClientTimeout,
			cfg.Parsers.Enabled,
			cache,
		), nil
	}, config.DefaultParsers)
}
