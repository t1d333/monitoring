package metrics

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/t1d333/vk_edu_db_project/internal/pkg/errors"
)

var (
	totalReq = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of requests",
		},
		[]string{"method"},
	)

	totalFailedRes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_responses_failed",
			Help: "Number of responses with 5xx code",
		},
		[]string{"status_code", "method"},
	)
)

func init() {
	prometheus.Register(totalReq)
	prometheus.Register(totalFailedRes)
}

func RegisterMetrics() {
}

func MetricsMiddleware(ctx *routing.Context) error {
	totalReq.With(prometheus.Labels{"method": string(ctx.Method())}).Inc()
	err := ctx.Next()
	if err != nil {
		code := strconv.Itoa(errors.ErrorToStatusCode[err])
		totalFailedRes.WithLabelValues(code, string(ctx.Method())).Inc()
	}
	return err
}

func ServeMetrics(addr string) error {
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(addr, mux)
	return err
}
