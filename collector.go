package otelfx

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx"
)

var CollectorModule = fx.Module("otel(collector)",
	fx.Provide(
		NewCollectorTraceProvider,
		NewCollectorMeterProvider,
		NewCollectorLogProvider,
	),
)

var collectorExportInterval = 3 * time.Second

type (
	CollectorEndpoint   string
	CollectorIsInsecure bool
)

func NewCollectorTraceProvider(ctx context.Context, lc fx.Lifecycle, endpoint CollectorEndpoint, insecure CollectorIsInsecure) (*trace.TracerProvider, error) {
	opts := []otlptracegrpc.Option{otlptracegrpc.WithEndpoint(string(endpoint))}
	if insecure {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}
	traceExporter, err := otlptracegrpc.New(ctx, opts...)
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			// Default is 5s. Set to 1s for demonstrative purposes.
			trace.WithBatchTimeout(time.Second)),
	)

	lc.Append(fx.StopHook(traceProvider.Shutdown))

	return traceProvider, nil
}

func NewCollectorMeterProvider(ctx context.Context, lc fx.Lifecycle, endpoint CollectorEndpoint, insecure CollectorIsInsecure) (*metric.MeterProvider, error) {
	opts := []otlpmetricgrpc.Option{otlpmetricgrpc.WithEndpoint(string(endpoint))}
	if insecure {
		opts = append(opts, otlpmetricgrpc.WithInsecure())
	}
	metricExporter, err := otlpmetricgrpc.New(ctx, opts...)
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter,
			// Default is 1m. Set to 3s for demonstrative purposes.
			metric.WithInterval(collectorExportInterval))),
	)
	lc.Append(fx.StopHook(meterProvider.Shutdown))
	return meterProvider, nil
}

func NewCollectorLogProvider(ctx context.Context, lc fx.Lifecycle, endpoint CollectorEndpoint, insecure CollectorIsInsecure) (*log.LoggerProvider, error) {
	opts := []otlploggrpc.Option{otlploggrpc.WithEndpoint(string(endpoint))}
	if insecure {
		opts = append(opts, otlploggrpc.WithInsecure())
	}
	// Set up logger provider.
	logExporter, err := otlploggrpc.New(ctx, opts...)
	if err != nil {
		return nil, err
	}

	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)

	lc.Append(fx.StopHook(loggerProvider.Shutdown))
	return loggerProvider, nil
}
