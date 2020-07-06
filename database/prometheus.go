package database

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Prometheus struct {
	Port          string `json:"port"`
	responseTime  *prometheus.GaugeVec
	responseError *prometheus.CounterVec
}

const (
	PrometheusName = "Prometheus"
	SubsystemName  = "statusok"
)

//GetDatabaseName return database name
func (p Prometheus) GetDatabaseName() string {
	return PrometheusName
}

//Initialize prometheus metrics
func (p *Prometheus) Initialize() error {
	log.Println("Prometheus : Initial some metrics")

	pt, err := strconv.Atoi(p.Port)
	if err != nil {
		return err
	}
	if pt < 1024 || pt > 49151 {
		return fmt.Errorf("Listen port cannot be %s", p.Port)
	}

	p.responseTime = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: SubsystemName,
		Name:      "response_time_msec",
		Help:      "API响应时间",
	}, []string{"url", "method", "status_code"})
	p.responseError = promauto.NewCounterVec(prometheus.CounterOpts{
		Subsystem: SubsystemName,
		Name:      "response_error",
		Help:      "API响应失败计数",
	}, []string{"url", "method", "status_code", "body", "reason"})

	go metricsServe(p.Port)
	return nil
}

//AddRequestInfo observe a request information
func (p Prometheus) AddRequestInfo(requestInfo RequestInfo) error {
	label := prometheus.Labels{
		"url":         requestInfo.Url,
		"method":      requestInfo.RequestType,
		"status_code": fmt.Sprintf("%d", requestInfo.ResponseCode),
	}
	p.responseTime.With(label).Set(float64(requestInfo.ResponseTime))

	return nil
}

//AddErrorInfo add Error information
func (p Prometheus) AddErrorInfo(errorInfo ErrorInfo) error {
	label := prometheus.Labels{
		"url":         errorInfo.Url,
		"method":      errorInfo.RequestType,
		"status_code": fmt.Sprintf("%d", errorInfo.ResponseCode),
		"body":        errorInfo.ResponseBody,
		"reason":      errorInfo.Reason.Error(),
	}
	p.responseError.With(label).Inc()

	return nil
}

func metricsServe(port string) {
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Starting metrics server at http://localhost:%s/metrics", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
