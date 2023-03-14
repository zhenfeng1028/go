package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 集群指标采集器
type ClusterManager struct {
	Zone         string           // 集群的名称
	OOMCountDesc *prometheus.Desc // 指标描述符
	RAMUsageDesc *prometheus.Desc // 指标描述符
}

// 实现采集工作，每次执行的时候，返回一个按照主机名作为键采集到的数据，两个返回值分别代表了OOM错误计数，和RAM使用指标信息。
func (c *ClusterManager) ReallyExpensiveAssessmentOfTheSystemState() (
	oomCountByHost map[string]int, ramUsageByHost map[string]float64,
) {
	oomCountByHost = map[string]int{
		"foo.example.org": int(rand.Int31n(1000)),
		"bar.example.org": int(rand.Int31n(1000)),
	}
	ramUsageByHost = map[string]float64{
		"foo.example.org": rand.Float64() * 100,
		"bar.example.org": rand.Float64() * 100,
	}
	return
}

// Describe simply sends the two Descs in the struct to the channel.
func (c *ClusterManager) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.OOMCountDesc
	ch <- c.RAMUsageDesc
}

// 将执行抓取函数并返回数据，返回的数据传递到channel中，并且传递的同时绑定原先的指标描述符。以及指标的类型（一个Counter和一个Gauge）
func (c *ClusterManager) Collect(ch chan<- prometheus.Metric) {
	oomCountByHost, ramUsageByHost := c.ReallyExpensiveAssessmentOfTheSystemState()
	for host, oomCount := range oomCountByHost {
		ch <- prometheus.MustNewConstMetric(
			c.OOMCountDesc,
			prometheus.CounterValue,
			float64(oomCount),
			host,
		)
	}
	for host, ramUsage := range ramUsageByHost {
		ch <- prometheus.MustNewConstMetric(
			c.RAMUsageDesc,
			prometheus.GaugeValue,
			ramUsage,
			host,
		)
	}
}

func NewClusterManager(zone string) *ClusterManager {
	return &ClusterManager{
		Zone: zone,
		OOMCountDesc: prometheus.NewDesc(
			"clustermanager_oom_crashes_total",
			"Number of OOM crashes.",
			[]string{"host"},
			prometheus.Labels{"zone": zone},
		),
		RAMUsageDesc: prometheus.NewDesc(
			"clustermanager_ram_usage_bytes",
			"RAM usage as reported to the cluster manager.",
			[]string{"host"},
			prometheus.Labels{"zone": zone},
		),
	}
}

func main() {
	workerDB := NewClusterManager("db")
	// workerCA := NewClusterManager("ca")

	// Since we are dealing with custom Collector implementations, it might
	// be a good idea to try it out with a pedantic registry.
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(workerDB)
	// reg.MustRegister(workerCA)

	gatherers := prometheus.Gatherers{
		prometheus.DefaultGatherer,
		reg,
	}

	h := promhttp.HandlerFor(gatherers,
		promhttp.HandlerOpts{
			ErrorLog:      log.Default(),
			ErrorHandling: promhttp.ContinueOnError,
		})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	log.Println("Start server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error occur when start server %v", err)
		os.Exit(1)
	}
}
