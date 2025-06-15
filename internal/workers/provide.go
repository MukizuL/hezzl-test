package workers

import "go.uber.org/fx"

func Provide() fx.Option {
	return fx.Provide(newClickHouseWriter, newNATSConsumer)
}
