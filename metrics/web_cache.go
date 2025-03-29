package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var webCacheCounter = promauto.NewCounterVec(prometheus.CounterOpts{
	Namespace: SystemName,
	Subsystem: SubSystemName,
	Name:      "web_cache_total",
	Help:      "Количество действий с кешом для реквестера",
}, []string{"action"})

func IncWebCacheCounter(action string) {
	webCacheCounter.WithLabelValues(action).Inc()
}
