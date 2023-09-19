package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/xiaomi-tc/log15"
)

const (
	ServiceName = "NATS_SERVER"
	ServiceId   = "nats_server"
)

var localAddress = getHostIP()

type Service struct {
	ServiceId    string
	ServiceName  string
	consulClient *api.Client
}

type ServiceInfo struct {
	ServiceID   string
	ServiceName string
	IP          string
	Port        int
	Status      string
	PackageName string
	Load        int
	Times       int   //被调用次数
	Timestamp   int64 //load updated ts
}

func main() {
	srv := &Service{}
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	srv.consulClient = client

	dcNames, err := srv.consulClient.Catalog().Datacenters()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dcNames:%v\n\n", dcNames)

	service := &api.AgentServiceRegistration{
		ID:      ServiceId,
		Name:    ServiceName,
		Port:    4222,
		Address: localAddress,
		Check: &api.AgentServiceCheck{
			HTTP:     "http://" + localAddress + ":4223/" + ServiceName + "/status",
			Interval: "10s",
			Timeout:  "1s",
		},
	}

	if err := srv.consulClient.Agent().ServiceRegister(service); err != nil {
		log.Error("RegistService", "error", err)
	}

	services, _, err := srv.consulClient.Catalog().Services(&api.QueryOptions{})
	if err != nil {
		panic(err)
	}
	bs, _ := json.Marshal(services)
	fmt.Printf("services:%s\n\n", string(bs))

	serviceEntries, _, err := srv.consulClient.Health().Service(ServiceName, "", false, &api.QueryOptions{})
	if err != nil {
		panic(err)
	}
	bs2, _ := json.Marshal(serviceEntries)
	fmt.Printf("serviceEntries:%s\n\n", string(bs2))

	items := make(map[string]*ServiceInfo)
	for _, entry := range serviceEntries {
		if entry.Service.Service != ServiceName {
			continue
		}
		for _, health := range entry.Checks {
			if health.ServiceName != ServiceName {
				continue
			}
			if health.Status != "passing" {
				continue
			}
			sInfo := new(ServiceInfo)
			sInfo.ServiceName = health.ServiceName
			sInfo.ServiceID = health.ServiceID
			if len(health.ServiceTags) > 0 {
				sInfo.PackageName = health.ServiceTags[0]
			}
			strIP, nPort := getBaseInfoWithServiceID(sInfo.ServiceName, sInfo.ServiceID)
			if strIP == "" || nPort == 0 {
				continue
			}
			sInfo.IP = strIP
			sInfo.Port = nPort
			sInfo.Status = health.Status
			items[sInfo.ServiceID] = sInfo
		}
	}
	bs3, _ := json.Marshal(items)
	fmt.Printf("items:%s\n\n", string(bs3))

	healthChecks, _, err := srv.consulClient.Health().State("any", &api.QueryOptions{WaitTime: 3 * time.Second})
	if err != nil {
		panic(err)
	}
	bs4, _ := json.Marshal(healthChecks)
	fmt.Printf("healthChecks:%s\n\n", string(bs4))

	err = srv.consulClient.Agent().ServiceDeregister(ServiceId)
	if err != nil {
		panic(err)
	}
}

func getBaseInfoWithServiceID(strServiceName, strServiceId string) (string, int) {
	listStr := strings.Split(strServiceId, "-")
	if len(listStr) == 3 && listStr[0] == strServiceName && listStr[1] != "" && listStr[2] != "" {
		nPort, _ := strconv.Atoi(listStr[2])
		if nPort > 0 {
			return listStr[1], nPort
		}
	}

	return "", 0
}

func startService(addr, serviceName string) {
	log.Info("start listen...", "addr", addr)

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

func getHostIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		// 判断是否正确获取到IP
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && !ipnet.IP.IsLinkLocalMulticast() && !ipnet.IP.IsLinkLocalUnicast() {
			log.Info("getHostIP", "ip", ipnet)
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	os.Stderr.WriteString("No Networking Interface Err!")
	return ""
}
