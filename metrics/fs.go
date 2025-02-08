package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var fsActionTime = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Namespace: SystemName,
	Subsystem: SubSystemName,
	Name:      "fs_action_seconds",
	Help:      "Время действий с файловой системой",
}, []string{"action"})

func RegisterFSActionTime(action string, d time.Duration) {
	fsActionTime.WithLabelValues(action).Observe(d.Seconds())
}
