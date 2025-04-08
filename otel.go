package otelfx

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"otel",
	fx.Invoke(
		Setup,
	),
)

// func SetupPropagator(ctx context.Context, lc fx.Lifecycle, log *slog.Logger, cfg *config.Config) {.
func Setup(tracerProvider *trace.TracerProvider, meterProvider *metric.MeterProvider, loggerProvider *log.LoggerProvider) error {
	prop := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
	otel.SetTextMapPropagator(prop)
	otel.SetTracerProvider(tracerProvider)
	otel.SetMeterProvider(meterProvider)
	global.SetLoggerProvider(loggerProvider)
	return startRuntimePublisher()
}
