package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 初始化 web_reqeust_total，counter类型指标，表示接收http请求总次数
var WebRequestTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "web_reqeust_total",
		Help: "Number of hello requests in total",
	},
	// 设置两个标签 请求方法 路径
	[]string{"method", "endpoint"},
)

// web_request_duration_seconds，Histogram类型指标，bucket代表duration的分布区间
var WebRequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "web_request_duration_seconds",
		Help:    "web request duration distribution",
		Buckets: []float64{0.1, 0.3, 0.5, 0.7, 0.9, 1},
	},
	[]string{"method", "endpoint"},
)

func init() {
	// 注册监控指标
	prometheus.MustRegister(WebRequestTotal)
	prometheus.MustRegister(WebRequestDuration)
}

// 包装 handler function,不侵入业务逻辑
func Monitor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h(w, r)
		duration := time.Since(start)
		WebRequestTotal.With(prometheus.Labels{"method": r.Method, "endpoint": r.URL.Path}).Inc()
		WebRequestDuration.With(prometheus.Labels{"method": r.Method, "endpoint": r.URL.Path}).Observe(duration.Seconds())
	}
}

func main() {
	// expose prometheus metrics接口
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/query", Monitor(Query))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// hello
func Hello(w http.ResponseWriter, r *http.Request) {
	// 请求计数
	WebRequestTotal.With(prometheus.Labels{"method": r.Method, "endpoint": r.URL.Path}).Inc()
	_, _ = io.WriteString(w, "hello world!")
}

// query
func Query(w http.ResponseWriter, r *http.Request) {
	// 模拟业务查询耗时0~1s
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	_, _ = io.WriteString(w, "some results")
}
