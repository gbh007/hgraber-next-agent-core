package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var parserActionTime = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Namespace: SystemName,
	Subsystem: SubSystemName,
	Name:      "parser_action_seconds",
	Help:      "Время действий парсинга",
}, []string{"action", "parser_name"})

func RegisterParserActionTime(action, parserName string, d time.Duration) {
	parserActionTime.WithLabelValues(action, parserName).Observe(d.Seconds())
}
