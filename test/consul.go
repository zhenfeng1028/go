package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	ServiceName = "NATS_SERVER"
	ServiceId   = "nats_server"
	servicePort = 4222
	checkPort   = 4223
)

var localAddress = getHostIP()

func main() {
	startMetricService(localAddress+":4223", ServiceName)

	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	dcNames, err := client.Catalog().Datacenters()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dcNames:%v\n\n", dcNames)

	service := &api.AgentServiceRegistration{
		ID:      ServiceId,
		Name:    ServiceName,
		Port:    servicePort,
		Address: localAddress,
		Check: &api.AgentServiceCheck{
			HTTP:     "http://" + localAddress + fmt.Sprint(":", checkPort, "/") + ServiceName + "/status",
			Interval: "5s",
			Timeout:  "1s",
		},
	}

	if err := client.Agent().ServiceRegister(service); err != nil {
		panic(err)
	}

	services, _, err := client.Catalog().Services(&api.QueryOptions{})
	if err != nil {
		panic(err)
	}
	bs, _ := json.Marshal(services)
	fmt.Printf("services:%s\n\n", string(bs))

	// 此时还没通过健康检查，因此nats_server服务的状态是critical
	serviceEntries, _, err := client.Health().Service(ServiceName, "", false, &api.QueryOptions{})
	if err != nil {
		panic(err)
	}
	bs2, _ := json.Marshal(serviceEntries)
	fmt.Printf("serviceEntries:%s\n\n", string(bs2))

	healthChecks, _, err := client.Health().State("any", &api.QueryOptions{WaitTime: 3 * time.Second})
	if err != nil {
		panic(err)
	}
	bs3, _ := json.Marshal(healthChecks)
	fmt.Printf("healthChecks:%s\n\n", string(bs3))

	time.Sleep(10 * time.Second)

	// 此时已经通过健康检查，nats_server服务的状态变为passing

	err = client.Agent().ServiceDeregister(ServiceId)
	if err != nil {
		panic(err)
	}
}

func getHostIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		// 判断是否正确获取到IP
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && !ipnet.IP.IsLinkLocalMulticast() && !ipnet.IP.IsLinkLocalUnicast() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	os.Stderr.WriteString("No Networking Interface Err!")
	return ""
}

func startMetricService(addr, serviceName string) {
	router := gin.Default()
	router.GET("/"+serviceName+"/status", statusHandler)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	pprofGroup := router.Group("/debug/pprof")
	pprofGroup.GET("goroutine", gin.WrapH(pprof.Handler("goroutine")))
	pprofGroup.GET("heap", gin.WrapH(pprof.Handler("heap")))
	pprofGroup.GET("block", gin.WrapH(pprof.Handler("block")))
	pprofGroup.GET("mutex", gin.WrapH(pprof.Handler("mutex")))
	pprofGroup.GET("threadcreate", gin.WrapH(pprof.Handler("threadcreate")))

	httpSvr := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := httpSvr.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
}

// 健康状态检查
func statusHandler(c *gin.Context) {
	c.String(http.StatusOK, "status ok!")
}
