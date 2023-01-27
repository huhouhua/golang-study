package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	v3 "golang-study/web/v3"
	"strconv"
	"time"
)

type MiddlewareBuilder struct {
	Namespace string
	Subsystem string
	Name      string
	Help      string
}

func (m MiddlewareBuilder) Build() v3.Middleware {
	vector := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:      m.Name,
		Subsystem: m.Subsystem,
		Namespace: m.Namespace,
		Help:      m.Help,
		Objectives: map[float64]float64{
			0.5:   0.01,
			0.75:  0.01,
			0.90:  0.01,
			0.99:  0.001,
			0.999: 0.0001,
		},
	}, []string{"pattern", "method", "status"})

	prometheus.MustRegister(vector)

	return func(next v3.HandlerFunc) v3.HandlerFunc {
		return func(ctx *v3.Context) {
			startTime := time.Now()
			defer func() {
				duration := time.Now().Sub(startTime).Milliseconds()

				pattern := ctx.MatchedRoute
				if pattern == "" {
					pattern = "unknown"
				}
				vector.WithLabelValues(pattern, ctx.Request.Method, strconv.Itoa(ctx.ResponseStatusCode)).
					Observe(float64(duration))
			}()
			next(ctx)

		}
	}
}
